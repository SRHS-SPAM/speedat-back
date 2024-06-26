package repositories

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"speedat-back/entities"
)

func MySQLInit() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entities.User{})

	return db
}
