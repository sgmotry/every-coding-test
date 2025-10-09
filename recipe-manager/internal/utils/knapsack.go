package utils

import (
	"every-coding-test/internal/model"
	"fmt"
	"strings"
)

func validateArg(recipes []model.Recipe, maxCalories int, maxCookingTime int) error {
	var errors []string
	if len(recipes) == 0 {
		errors = append(errors, "no recipe data input")
	}
	if maxCalories <= 0 {
		errors = append(errors, "maxCalories must be greater than or equal to 1")
	}
	if maxCookingTime <= 0 {
		errors = append(errors, "maxCookingTime must be greater than or equal to 1")
	}
	if len(errors) > 0 {
		return fmt.Errorf("error\nvalidation failed:\n%s", strings.Join(errors, "\n"))
	}
	return nil
}

func GetMaxProteinRecipeCombination(recipes []model.Recipe, maxCalories int, maxCookingTime int) ([]uint32, float32, error) {
	argErr := validateArg(recipes, maxCalories, maxCookingTime)
	if argErr != nil {
		return nil, 0, argErr
	}

	recipesLen := len(recipes)

	dp := make([][][]int, recipesLen+1)
	for i := range dp {
		dp[i] = make([][]int, maxCalories+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxCookingTime+1)
		}
	}

	for i := 1; i <= recipesLen; i++ {
		recipe := recipes[i-1]
		calorie := int(recipe.Nutrition.Calorie)
		time := recipe.Time
		protein := int(recipe.Nutrition.Protein * 10)

		for c := 0; c <= maxCalories; c++ {
			for t := 0; t <= maxCookingTime; t++ {
				dp[i][c][t] = dp[i-1][c][t]

				if c >= calorie && t >= time {
					withRecipe := dp[i-1][c-calorie][t-time] + protein
					if withRecipe > dp[i][c][t] {
						dp[i][c][t] = withRecipe
					}
				}
			}
		}
	}
	totalProtein := float32(dp[recipesLen][maxCalories][maxCookingTime]) / 10

	selectedRecipesId := []uint32{}
	c := maxCalories
	t := maxCookingTime

	for i := recipesLen; i > 0; i-- {
		recipe := recipes[i-1]
		calorie := int(recipe.Nutrition.Calorie)
		time := recipe.Time
		protein := int(recipe.Nutrition.Protein * 10)

		// レシピが選ばれているかチェック
		if c >= calorie && t >= time {
			if dp[i][c][t] == dp[i-1][c-calorie][t-time]+protein && dp[i][c][t] > dp[i-1][c][t] {
				selectedRecipesId = append([]uint32{recipe.Id}, selectedRecipesId...)
				c -= calorie
				t -= time
			}
		}
	}

	return selectedRecipesId, totalProtein, nil
}
