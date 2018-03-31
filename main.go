package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/config"
	"github.com/h00s/url-shortener-backend/db"
	"github.com/h00s/url-shortener-backend/link"
	"github.com/h00s/url-shortener-backend/logger"
)

func main() {
	l, err := logger.New("url-shortener-backend.log")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	config, err := config.Load("configuration.json")
	if err != nil {
		l.Fatal(err.Error())
	}

	db, err := db.Connect(config)
	if err != nil {
		l.Fatal(err.Error())
	}

	err = db.Migrate()
	if err != nil {
		l.Fatal(err.Error())
	}

	lc := link.NewController(db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/links/:name", lc.GetLink)
		v1.GET("/links/:name/redirect", lc.RedirectToLink)
		v1.GET("/links/:name/stats", lc.GetLinkActivityStats)
		v1.POST("/links", lc.InsertLink)
	}

	l.Info("Started gin")
	r.Run()
}
