package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thephiri/gogos-recipes-backend/controllers"
	"github.com/thephiri/gogos-recipes-backend/middleware"
)

func RecipesRoutes(route fiber.Router) {
	//unprotected routes

	route.Get("/", controllers.GetAllRecipes)
	route.Get("/:id", controllers.GetRecipeById)
	route.Get("/culture/:culture", controllers.GetRecipeByCulture)

	route.Use(middleware.Authentication)
	route.Post("/", controllers.CreateRecipe)

}
