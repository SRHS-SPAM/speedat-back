package services

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"speedat-back/entities"
)

func VerifySend(c *gin.Context) {
	var userDTO entities.UserDTO
	RandomNumber := rand.Intn(900000) + 100000

	// JSON 요청 바인딩
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	VerifyEmail(userDTO.Email, RandomNumber)
	c.JSON(http.StatusOK, gin.H{
		"message": RandomNumber,
	})
}
