package awd_probe_client

import (
	"bytes"
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func Proxy(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	req := c.Request.Clone(context.TODO())
	req.RequestURI = ""
	req.URL.Host = ProbeHost
	req.URL.Scheme = "http"
	if c.Request.TLS != nil {
		req.URL.Scheme = "https"
	}
	req.Header.Set("Probe-Forwarded-For", c.ClientIP())
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	dump, _ := httputil.DumpResponse(resp, true)
	spew.Dump(dump)
}
