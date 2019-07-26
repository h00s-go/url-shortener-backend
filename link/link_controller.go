package link

import (
	"github.com/h00s/url-shortener-backend/db"
	"github.com/h00s/url-shortener-backend/logger"
	"github.com/labstack/echo"
)

// Controller for Link methods
type Controller struct {
	db     *db.Database
	logger *logger.Logger
}

type errorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

// NewController creates new link controller
func NewController(db *db.Database, logger *logger.Logger) *Controller {
	return &Controller{db: db, logger: logger}
}

// GetLink get link with specific name
func (lc *Controller) GetLink(c echo.Context) error {
	return lc.getLink(c, false)
}

// RedirectToLink get link with specifig name and redirects to it's url
func (lc *Controller) RedirectToLink(c echo.Context) error {
	return lc.getLink(c, true)
}

func (lc *Controller) getLink(c echo.Context, redirect bool) error {
	name := c.Param("name")
	l, err := getLinkByName(lc.db, name)

	switch {
	case l != nil:
		if err := insertActivity(lc.db, l.ID, c.RealIP()); err != nil {
			lc.logger.Error(err.Error())
		}
		if redirect {
			return c.Redirect(302, l.URL)
		}
		return c.JSON(200, l)
	case err != nil:
		lc.logger.Error(err.Error())
		return c.JSON(500, errorResponse{"Error while getting link", "There was an server error when getting link"})
	default:
		return c.JSON(404, errorResponse{"Link not found", "Link with specified name not found"})
	}
}

// GetLinkActivityStats get link with specific name
func (lc *Controller) GetLinkActivityStats(c echo.Context) error {
	name := c.Param("name")
	id := getIDFromName(name)
	s, err := getLinkActivityStats(lc.db, id)

	switch {
	case s != nil:
		return c.JSON(200, s)
	case err != nil:
		lc.logger.Error(err.Error())
		return c.JSON(500, errorResponse{"Error while getting link", "There was an server error when getting link"})
	default:
		return c.JSON(404, errorResponse{"Link not found", "Link with specified name not found"})
	}
}

// InsertLink adds new link
func (lc *Controller) InsertLink(c echo.Context) error {
	if !lc.isSpammer(c.RealIP()) {
		var link Link
		if err := c.Bind(&link); err == nil {
			l, err := insertLink(lc.db, link.URL, "", c.RealIP())
			if err == nil {
				return c.JSON(201, l)
			}
			lc.logger.Error(err.Error())
			return c.JSON(500, errorResponse{"Error while adding link", "There was an server error when adding link"})
		}
		return c.JSON(400, errorResponse{"Request is invalid", "Request should be a JSON object containing url"})
	}
	return c.JSON(429, errorResponse{"Rate limiting", "Too many links posted, please wait couple of minutes"})
}

func (lc *Controller) isSpammer(clientAddress string) bool {
	linkCount := 0
	lc.db.Conn.QueryRow(sqlGetPostCountInLastMinutes, clientAddress, 10).Scan(&linkCount)

	if linkCount >= 10 {
		return true
	}
	return false
}
