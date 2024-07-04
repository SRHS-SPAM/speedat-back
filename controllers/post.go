package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"speedat-back/services"
)

func Post(r *gin.Engine, rdb *gorm.DB) {

	post := r.Group("post")
	post.Use(services.AuthMiddleware())
	{
		post.GET("/", func(c *gin.Context) {
			post.GET("/protected", func(c *gin.Context) {
				email := c.MustGet("email").(string)
				c.JSON(http.StatusOK, gin.H{
					"message": "이곳은 인증된 사용자만 접근할 수 있습니다",
					"email":   email,
				})
			})
			c.JSON(200, "인증 된것이와요.")
		})
	}
}
