package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/db"
)

// Controller for Link methods
type Controller struct {
	db *db.Database
}

// PostLinkData represents data which is sent when posting new url
type PostLinkData struct {
	URL string `json:"url" binding:"required"`
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
func (lc *Controller) PostLink(c *gin.Context) {
	var postLinkData PostLinkData
	if err := c.BindJSON(&postLinkData); err == nil {
		l, err := InsertLink(lc, postLinkData.URL, c.ClientIP())
		if err == nil {
			c.JSON(200, l)
		} else {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
		}
	} else {
		c.JSON(404, gin.H{
			"message": "request is invalid",
		})
	}
}
