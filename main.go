package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/thephiri/gogos-recipes-backend/src/database"
	"github.com/thephiri/gogos-recipes-backend/src/routes"
)

func main() {
	fmt.Println("Starting GoGoS Recipes Backend")
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Print("Error loading .env file")

		}

	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Use(logger.New())
	database.Connect()

	routes.Setup(app)

	fmt.Println("Starting server on port 5000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Println("Listening on port " + port)

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
