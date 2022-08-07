package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thephiri/gogos-recipes-backend/controllers"
)

func RecipesRoutes(route fiber.Router) {
	route.Post("/", controllers.CreateRecipe)
	route.Get("/", controllers.GetAllRecipes)
}
