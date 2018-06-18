package main

import (
	"log"

	"github.com/h00s/url-shortener-backend/config"
	"github.com/h00s/url-shortener-backend/db"
	"github.com/h00s/url-shortener-backend/link"
	"github.com/h00s/url-shortener-backend/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	c, err := config.Load("configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.New(c.Log.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	db, err := db.Connect(c.Database)
	if err != nil {
		l.Fatal(err.Error())
	}

	err = db.Migrate()
	if err != nil {
		l.Fatal(err.Error())
	}

	lc := link.NewController(db, l)

	r := echo.New()
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	v1 := r.Group("/v1")
	{
		v1.GET("/links/:name", lc.GetLink)
		v1.GET("/links/:name/redirect", lc.RedirectToLink)
		v1.GET("/links/:name/stats", lc.GetLinkActivityStats)
		v1.POST("/links", lc.InsertLink)
	}

	l.Info("Starting HTTP server...")
	r.Logger.Fatal(r.Start(c.Server.Address))
}
