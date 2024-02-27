package services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"speedat-back/entities"
)

var RandomNumber = rand.Intn(900000) + 100000

func SignUp(c *gin.Context, db *gorm.DB) {
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		panic("Failed to migrate database")
	}

	var userDTO entities.UserDTO

	err := c.ShouldBindJSON(&userDTO)

	if RandomNumber != userDTO.VerifyCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "인증 코드가 맞지 않습니다.",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not hash password",
		})
		return
	}

	user := entities.User{
		Email:    userDTO.Email,
		Password: string(hashedPassword),
		Name:     userDTO.Name,
		Grade:    userDTO.Grade,
		Class:    userDTO.Class,
		Number:   userDTO.Number,
	}

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func VerifySend(c *gin.Context, db *gorm.DB) {
	var userDTO entities.UserDTO

	// JSON 요청 바인딩
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	VerifyEmail(userDTO.Email, RandomNumber)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
