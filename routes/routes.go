package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the root endpoint 😉",
		})
	})

	api := app.Group("/api")

	RecipesRoutes(api.Group("/recipes"))
}
