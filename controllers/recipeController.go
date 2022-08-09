package controllers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	models "github.com/thephiri/gogos-recipes-backend/Models"
	"github.com/thephiri/gogos-recipes-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRecipe(c *fiber.Ctx) error {
	recipeCollection := database.MI.DB.Collection("recipes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	recipe := new(models.Recipe)

	if err := c.BodyParser(recipe); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error parsing body",
				"success": false,
				"error":   err,
			},
		)
	}

	recipe.CreatedAt = time.Now()
	recipe.UpdatedAt = time.Now()

	result, err := recipeCollection.InsertOne(ctx, recipe)
	if err != nil {
		log.Printf("Error inserting recipe: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error inserting recipe",
				"success": false,
				"error":   err,
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"message": "Recipe created",
			"success": true,
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
		log.Printf("Error finding recipes: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error finding recipes",
				"success": false,
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
					"success": false,
					"error":   err,
				},
			)
		}

		recipes = append(recipes, recipe)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Recipes found",
			"success": true,
			"data":    recipes,
		},
	)

}

func GetRecipeById(c *fiber.Ctx) error {
	recipeCollection := "recipes"

	recipe, err := database.GetById(recipeCollection, c.Params("id"))
	if err != nil {
		log.Printf("Error finding recipe: %v", err)
		return c.Status(500).JSON(
			fiber.Map{
				"message": "Error finding recipe",
				"success": false,
				"error":   err,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Recipe found",
			"success": true,
			"data":    recipe,
		},
	)

}
