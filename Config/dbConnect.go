package Config

import (
	"fmt"
	"os"

	"github.com/KaisAlHusrom/RestApi/Models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	//download env file
	godotenv.Load(".env")

	//Database Info
	HOST := os.Getenv("HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER_NAME := os.Getenv("DB_USER_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	//Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, DB_USER_NAME, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	autoMigrate(db)
}

func autoMigrate(db *gorm.DB) {
	db.Debug().AutoMigrate(
		&Models.User{},
		&Models.Product{},
	)
}
