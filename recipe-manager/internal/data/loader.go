package data

import (
	"encoding/json"
	"every-coding-test/internal/model"
	"fmt"
	"os"
)

// ローカルファイルからrecipes.jsonを読み込む
func GetRecipes() ([]model.Recipe, error) {
	bytes, err := os.ReadFile("recipes.json")
	if err != nil {
		return nil, fmt.Errorf("error while reading recipes.json.\n%e", err)
	}

	var recipeData []model.Recipe
	if err := json.Unmarshal(bytes, &recipeData); err != nil {
		return nil, fmt.Errorf("error while unmarshaling json.\n%e", err)
	}

	return recipeData, err
}
