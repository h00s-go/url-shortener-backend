package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/url-shortener-backend/controller"
)

func main() {
	c, err := controller.NewController()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/link/:name", c.GetLink)
		v1.POST("/link", c.PostLink)
	}

	r.Run()
}
