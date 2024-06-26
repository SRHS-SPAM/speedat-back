package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"speedat-back/services"
)

func Auth(r *gin.Engine, rdb *gorm.DB) {
	auth := r.Group("auth")
	{
		auth.POST("/verify", func(c *gin.Context) {
			services.VerifySend(c)
		})
		auth.POST("/signup", func(c *gin.Context) {
			services.SignUp(c, rdb)
		})
	}
}
