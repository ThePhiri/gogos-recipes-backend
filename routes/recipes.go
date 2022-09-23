package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thephiri/gogos-recipes-backend/controllers"
)

func RecipesRoutes(route fiber.Router) {
	//unprotected routes
	route.Get("/", controllers.GetAllRecipes)
	route.Get("/:id", controllers.GetRecipeById)
	route.Get("/culture/:culture", controllers.GetRecipeByCulture)

	//protected routes
	route.Post("/", controllers.CreateRecipe)
	

}
