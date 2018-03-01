package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/config"
	"github.com/h00s/url-shortener-backend/db"
	"github.com/h00s/url-shortener-backend/link"
)

func main() {
	config, err := config.LoadConfiguration("configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	lc := link.NewController(db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/link/:name", lc.GetLink)
		v1.POST("/link", lc.PostLink)
	}

	r.Run()
}
