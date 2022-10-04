package middleware

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	helper "github.com/thephiri/gogos-recipes-backend/helpers"
)

func Authentication(c *fiber.Ctx) error {
	//maybe switch this to use cookies?
	clientToken := c.Get("token")
	if clientToken == "" {
		log.Print("Error: token not found")
		return errors.New("eish no token")
	}

	log.Printf("client token is %v", clientToken)

	claims, err := helper.ValidateToken(clientToken)
	if err != "" {
		c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Error binding json",
				"status":  "error",
				"error":   err,
			},
		)
		return errors.New(err)
	}

	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("uid", claims.Uid)

	return c.Next()
}
