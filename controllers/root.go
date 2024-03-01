package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"speedat-back/services"
	"time"
)

func NewController(port string) {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://www.speedat.site/", "http://112.159.30.237", "http://122.203.181.62", "http://175.114.18.77", "http://180.70.171.163", "http://210.204.194.10", "http://58.224.240.125"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		MaxAge:       24 * time.Hour,
	}))

	auth := r.Group("auth")
	{
		auth.POST("/verify", func(c *gin.Context) {
			err := services.VerifySend(c)
			log.Fatalln(err)
		})
	}

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
