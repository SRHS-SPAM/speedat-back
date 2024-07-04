package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"speedat-back/repositories"
	"time"
)

func NewController(port string) {
	r := gin.New()
	rdb := repositories.MySQLInit()

	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		MaxAge:       24 * time.Hour,
	}))

	Auth(r, rdb)
	Post(r, rdb)

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
