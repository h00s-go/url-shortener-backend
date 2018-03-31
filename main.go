package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/config"
	"github.com/h00s/url-shortener-backend/db"
	"github.com/h00s/url-shortener-backend/link"
)

func main() {
	config, err := config.Load("configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
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

	r.Run()
}
