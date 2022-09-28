package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	models "github.com/thephiri/gogos-recipes-backend/Models"
	"github.com/thephiri/gogos-recipes-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var recipeCollection = "recipes"

//TODO: Sepearate out create and get all logic to database.go file

func CreateRecipe(c *fiber.Ctx) error {
	//check if user is logged in
	recipeCollection := database.MI.DB.Collection("recipes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	recipe := new(models.Recipe)

	if err := c.BodyParser(recipe); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error parsing body",
				"status":  "error",
				"error":   err,
			},
		)
	}

	recipe.CreatedAt = time.Now()
	recipe.UpdatedAt = time.Now()

	result, err := recipeCollection.InsertOne(ctx, recipe)
	if err != nil {
		fmt.Printf("Error inserting recipe: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error inserting recipe",
				"status":  "error",
				"error":   err,
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"message": "Recipe created",
			"status":  "success",
			"data":    result,
		},
	)
}

func GetAllRecipes(c *fiber.Ctx) error {
	recipeCollection := database.MI.DB.Collection("recipes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var recipes []models.Recipe

	filter := bson.M{}
	findOptions := options.Find()

	cur, err := recipeCollection.Find(ctx, filter, findOptions)
	if err != nil {
		fmt.Printf("Error finding recipes: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error finding recipes",
				"status":  "error",
				"error":   err,
			},
		)
	}

	for cur.Next(ctx) {
		var recipe models.Recipe
		err := cur.Decode(&recipe)
		if err != nil {
			log.Printf("Error decoding recipe: %v", err)
			return c.Status(500).JSON(
				fiber.Map{
					"message": "Error decoding recipe",
					"status":  "error",
					"error":   err,
				},
			)
		}

		recipes = append(recipes, recipe)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Recipes found",
			"status":  "error",
			"data":    recipes,
		},
	)

}

func GetRecipeById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		log.Print("Error: id not found")
		return errors.New("id not found")
	}

	fmt.Printf("Id: %v", id)

	recipe, err := database.GetById(recipeCollection, id)
	if err != nil {
		log.Printf("Error finding recipe: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error finding recipe",
				"status":  "error",
				"error":   err,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Recipe found",
			"status":  "error",
			"data":    recipe,
		},
	)

}

func GetRecipeByCulture(c *fiber.Ctx) error {
	fmt.Printf("Culture: %v", c.Params("culture"))

	culture := c.Params("culture")
	fmt.Printf("Culture: %v", culture)
	if culture == "" {
		log.Print("Error: culture not found")
		return errors.New("culture not found")
	}

	cultureFilter := bson.M{"culture": culture}

	recipes, err := database.GetByFilter(recipeCollection, cultureFilter)
	if err != nil {
		log.Printf("Error finding recipes: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error finding recipe",
				"status":  "error",
				"error":   err,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Recipe found",
			"status":  "error",
			"data":    recipes,
		},
	)
}

//update recipe function

//delete recipe function

//get recipes by user id
func GetByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")

	if userId == "" {
		log.Print("Error: userid not found")
		return errors.New("userid not found")
	}

	recipes, err := database.GetByFilter(recipeCollection, bson.M{"userId": userId})
	if err != nil {
		log.Printf("Error finding recipes: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error finding recipes",
				"status":  "error",
				"error":   err,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Recipes found",
			"status":  "error",
			"data":    recipes,
		},
	)
}
