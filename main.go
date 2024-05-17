package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":         "ok",
			"app_version":     "3.0",
			"app_description": "deploy app version 2.0 build 2",
		})
	})
	r.GET("/env", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":     200,
			"data_env": os.Environ(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
