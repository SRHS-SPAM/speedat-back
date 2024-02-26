package controllers

import (
	"github.com/gin-gonic/gin"
	"speedat-back/services"
)

func NewController(port string) {
	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/verify", func(c *gin.Context) {
		services.VerifyEmail()
	})

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
