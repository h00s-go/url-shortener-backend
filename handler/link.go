package handler

import (
	"github.com/gin-gonic/gin"
)

// GetLink get link with specific name
func (h *Handler) GetLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "ok",
	})
}

// PostLink adds new link
func (h *Handler) PostLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "ok",
	})
}
