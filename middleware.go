package awd_probe_client

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Proxy(c *gin.Context) {
	req := c.Request.Clone(context.TODO())
	req.Host = "probe:13500"
	http.DefaultClient.Do(req)
}
