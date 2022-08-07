package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/thephiri/gogos-recipes-backend/database"
	"github.com/thephiri/gogos-recipes-backend/routes"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Use(logger.New())
	database.Connect()

	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
