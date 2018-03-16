package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/db"
)

// Controller for Link methods
type Controller struct {
	db *db.Database
}

// NewController creates new link controller
func NewController(db *db.Database) *Controller {
	return &Controller{db: db}
}

// GetLink get link with specific name
func (lc *Controller) GetLink(c *gin.Context) {
	name := c.Param("name")
	l, err := GetLink(lc, name)

	if err == nil {
		c.Redirect(302, l.URL)
	} else {
		c.JSON(404, gin.H{
			"message": "link not found",
		})
	}
}

// PostLink adds new link
func (*Controller) PostLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "ok",
	})
}
