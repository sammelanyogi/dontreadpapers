package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
