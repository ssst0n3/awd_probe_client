package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awd_probe_client"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(awd_probe_client.Proxy)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	awesome_error.CheckFatal(r.Run(":8080"))
}
