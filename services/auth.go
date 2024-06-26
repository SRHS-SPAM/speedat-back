package services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"speedat-back/entities"
	"strings"
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

	domain := strings.Split(user.Email, "@")
	if domain[1] != "sonline20.sen.go.kr" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "이메일 도메인이 맞지않음",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "비밀번호 해싱 중 오류가 발생했습니다.",
		})
		return
	}

	upload := &entities.User{
		Email:        user.Email,
		Password:     string(hashedPassword),
		Name:         user.Name,
		Grade:        user.Grade,
		Class:        user.Class,
		Number:       user.Number,
		ProfilePhoto: "기본값",
	}

	if err := rdb.Create(upload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "유저 정보를 데이터베이스에 삽입하는 중 오류가 발생했습니다.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "회원가입 성공",
		"user":    upload,
	})
}

func Login(c *gin.Context, rdb *gorm.DB) {
	
}
