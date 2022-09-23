package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	models "github.com/thephiri/gogos-recipes-backend/Models"
	database "github.com/thephiri/gogos-recipes-backend/database"
	helper "github.com/thephiri/gogos-recipes-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Print("Error : %v", err)
		return "Error"
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or password is incorrect")
		check = false
	}

	return check, msg
}

func SignUp(c *fiber.Ctx) error {
	var userCollection = database.MI.DB.Collection("users")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Error binding json",
				"status":  "error",
				"error":   err,
			},
		)
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Validation error",
				"status":  "error",
				"error":   validationErr,
			},
		)
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()

	if err != nil {
		log.Printf("Error getting count: %v", err)
		c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Error getting count",
				"status":  "error",
				"error":   err,
			},
		)
		return err

	}

	password := HashPassword(user.Password)
	user.Password = password

	if count > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "User already exists",
				"status":  "error",
				"error":   err,
			},
		)
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_id)
	user.Token = token
	user.Refresh_token = refreshToken

	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error parsing body",
				"status":  "error",
				"error":   err,
			},
		)
	}

	defer cancel()

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"message": "User created",
			"status":  "success",
			"data":    resultInsertionNumber,
		},
	)
}

func Login(c *fiber.Ctx) error {
	var userCollection = database.MI.DB.Collection("users")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	var foundUser models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error parsing body",
				"status":  "error",
				"error":   err,
			},
		)
	}

	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Email not found",
				"status":  "error",
				"error":   err,
			},
		)
	}

	passwordIsValid, msg := VerifyPassword(*&user.Password, *&foundUser.Password)
	defer cancel()
	if passwordIsValid != true {
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Incorrect Password",
				"status":  "error",
				"error":   msg,
			},
		)
	}

	token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_id)

	helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Logged in",
			"status":  "success",
			"data":    foundUser,
		},
	)

}
