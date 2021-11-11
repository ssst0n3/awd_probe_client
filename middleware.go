package awd_probe_client

import (
	"bytes"
	"context"
	beegoContext "github.com/astaxie/beego/context"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func Proxy(request *http.Request, clientIp string) {
	if request.Header.Get("Probe-Repeater") == "yes" {
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	request.Body = ioutil.NopCloser(bytes.NewReader(body))
	req := request.Clone(context.TODO())
	req.RequestURI = ""
	req.URL.Host = ProbeHost
	req.URL.Scheme = "http"
	if request.TLS != nil {
		req.URL.Scheme = "https"
	}
	req.Header.Set("Probe-Forwarded-For", clientIp)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	dump, _ := httputil.DumpResponse(resp, true)
	log.Logger.Debug(dump)
}

func Gin(c *gin.Context) {
	Proxy(c.Request, c.ClientIP())
}

func getClientIp(req *http.Request) string {
	clientIP := req.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(req.Header.Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}
	if addr := req.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
		return addr
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func Beego(c *beegoContext.Context) {
	Proxy(c.Request, getClientIp(c.Request))
}
