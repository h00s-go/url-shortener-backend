package link

import (
	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/db"
)

// Controller for Link methods
type Controller struct {
	db *db.Database
}

// InsertLinkData represents data which is sent when posting new url
type InsertLinkData struct {
	URL string `json:"url" binding:"required"`
}

// NewController creates new link controller
func NewController(db *db.Database) *Controller {
	return &Controller{db: db}
}

// GetLink get link with specific name
func (lc *Controller) GetLink(c *gin.Context) {
	name := c.Param("name")
	l, err := getLinkByName(lc, name)

	if err == nil {
		c.Redirect(302, l.URL)
	} else {
		c.JSON(404, gin.H{
			"message": "link not found",
		})
	}
}

// InsertLink adds new link
func (lc *Controller) InsertLink(c *gin.Context) {
	if !lc.isSpammer(c.ClientIP()) {
		var linkData InsertLinkData
		if err := c.BindJSON(&linkData); err == nil {
			l, err := insertLink(lc, linkData.URL, c.ClientIP())
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
	} else {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "too many links posted, please wait couple of minutes",
		})
	}
}

func (lc *Controller) isSpammer(clientAddress string) bool {
	linkCount := 0
	lc.db.Conn.QueryRow(sqlGetPostCountInLastMinutes, clientAddress, 10).Scan(&linkCount)

	if linkCount >= 10 {
		return true
	}
	return false
}
