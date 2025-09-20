package handlers

import "github.com/gin-gonic/gin"

func renderError(c *gin.Context, template string, message string, statusCode int) {
	c.HTML(statusCode, template, gin.H{
		"Title": "Error",
		"Error": message,
	})
}
