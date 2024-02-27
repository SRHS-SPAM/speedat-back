package controllers

import (
	"github.com/gin-gonic/gin"
	"speedat-back/repository"
	"speedat-back/services"
)

func NewController(port string) {
	r := gin.New()

	r.Use(gin.Logger())
	rdb := repository.MySQLInit()

	auth := r.Group("auth")
	{
		auth.POST("/signup", func(c *gin.Context) {
			services.SignUp(c, rdb)
		})
		auth.GET("/verify", func(c *gin.Context) {
			services.VerifySend(c, rdb)
		})
	}

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
