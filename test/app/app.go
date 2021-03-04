package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awd_probe_client"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"net/http"
)

func main() {
	awd_probe_client.ProbeHost = "127.0.0.1:13500"
	r := gin.New()
	r.Use(awd_probe_client.Proxy)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	awesome_error.CheckFatal(r.Run(":8888"))
}
