package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/db"
	"github.com/h00s/url-shortener-backend/link"
)

func main() {
	db, err := db.NewDatabase()
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
