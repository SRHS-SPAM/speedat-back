package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"speedat-back/services"
	"time"
)

func NewController(port string) {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://www.speedat.site/"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		MaxAge:       24 * time.Hour,
	}))

	auth := r.Group("auth")
	{
		auth.POST("/verify", func(c *gin.Context) {
			services.VerifySend(c)
		})
	}

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
