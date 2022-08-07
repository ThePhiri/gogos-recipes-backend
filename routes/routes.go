package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	RecipesRoutes(api.Group("/recipes"))
}
