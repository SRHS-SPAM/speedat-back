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
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		MaxAge:       24 * time.Hour,
	}))

	auth := r.Group("auth")
	{
		auth.POST("/verify", func(c *gin.Context) {
			err := services.VerifySend(c)
			log.Println(err)
		})
	}

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
