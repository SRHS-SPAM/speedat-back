package services

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"net/smtp"
	"speedat-back/entities"
)

func SignUp(c *gin.Context, db *gorm.DB) {
	var userDTO entities.UserDTO

	err := c.ShouldBindJSON(&userDTO)
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

func VerifyEmail() {
	smtpHost := "smtp.zoho.com"
	smtpPort := 465
	email := "verify@speedat.site"
	password := "js9xUc!g"

	// 발신자 및 수신자 이메일 주소 설정
	from := email
	to := "shrie0602@gmail.com"

	// 이메일 메시지 구성
	subject := "speedat 및 스팸 시스템 인증 메일 입니다."
	body := "김채운 개병신"

	// 이메일 헤더 및 내용 구성
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body)

	// SMTP 인증 구성
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// TLS 구성
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// TLS 연결 설정
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", smtpHost, smtpPort), tlsConfig)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}

	// SMTP 클라이언트 생성 및 인증 및 전송
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		return
	}
	defer client.Quit()

	// 인증
	if err := client.Auth(auth); err != nil {
		fmt.Println("Error authenticating:", err)
		return
	}

	// 메시지 전송
	if err := client.Mail(from); err != nil {
		fmt.Println("Error setting sender:", err)
		return
	}
	if err := client.Rcpt(to); err != nil {
		fmt.Println("Error setting recipient:", err)
		return
	}
	w, err := client.Data()
	if err != nil {
		fmt.Println("Error getting data:", err)
		return
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing message:", err)
		return
	}
	err = w.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}
	fmt.Println("Email sent successfully!")
}
