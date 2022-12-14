package routes

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	fmt.Println("Setup routes")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "You are at the root endpoint 😉",
		})
	})

	api := app.Group("/api")

	RecipesRoutes(api.Group("/recipes"))
	UserRoutes(api.Group("/users"))

}
