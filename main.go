package main

import (
	"log"
	"os"

	"github.com/KaisAlHusrom/RestApi/Config"
	"github.com/KaisAlHusrom/RestApi/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	Config.Connect()

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	app := fiber.New()

	Routes.SetUp(app)

	log.Fatal(app.Listen(":" + portString))
}
