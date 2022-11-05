package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thephiri/gogos-recipes-backend/controllers"
)

func UserRoutes(route fiber.Router) {

	route.Post("/signup", controllers.SignUp)
	route.Post("/login", controllers.Login)

}
