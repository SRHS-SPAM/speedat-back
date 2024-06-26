package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"speedat-back/entities"
	"unicode/utf8"
)

func VerifySend(c *gin.Context) error {
	var userDTO entities.UserDTO
	RandomNumber := rand.Intn(900000) + 100000

	// JSON 요청 바인딩
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	VerifyEmail(userDTO.Email, RandomNumber)
	c.JSON(http.StatusOK, gin.H{
		"message": RandomNumber,
	})
	return nil
}

func SignUp(c *gin.Context, rdb *gorm.DB) {
	var user *entities.UserDTO

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	upload := &entities.User{
		Email: user.Email,
		Password: user.Password,
		Name: user.Name,
		Grade: user.Grade,
		Class: user.Class,
		Number: user.Number,
		ProfilePhoto: "기본값",
	}

	

	if upload.Email != {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "이메일 도메인이 맞지않음"
		})
	}

}
