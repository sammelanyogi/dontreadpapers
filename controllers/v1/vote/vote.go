package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpvotePaper(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func DownvotePaper(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
