package controller

import (
	"github.com/gin-gonic/gin"
)

// GetLink get link with specific name
func (*Controller) GetLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "ok",
	})
}

// PostLink adds new link
func (*Controller) PostLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "ok",
	})
}
