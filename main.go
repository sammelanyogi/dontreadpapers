package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	server.Run(":8080")
}
