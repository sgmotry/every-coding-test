package data

import (
	"encoding/json"
	"every-coding-test/internal/model"
	"os"
)

func GetRecipes() ([]model.Recipe, error) {
	bytes, err := os.ReadFile("recipes.json")
	if err != nil {
		return nil, err
	}

	var recipeData []model.Recipe
	if err := json.Unmarshal(bytes, &recipeData); err != nil {
		return nil, err
	}

	return recipeData, err
}
