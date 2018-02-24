package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/handler"
)

func main() {
	h, err := handler.NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/link/:name", h.GetLink)
		v1.POST("/link", h.PostLink)
	}

	r.Run()
}
