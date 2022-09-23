package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	helper "github.com/thephiri/gogos-recipes-backend/helpers"
)

func Authentication(c *fiber.Ctx) {
	clientToken := c.Params("token")
	if clientToken == "" {
		log.Print("Error: token not found")
		return
	}

	claims, err := helper.ValidateToken(clientToken)
	if err != "" {
		c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Error binding json",
				"status":  "error",
				"error":   err,
			},
		)
		return
	}

	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("uid", claims.Uid)

	c.Next()
}
