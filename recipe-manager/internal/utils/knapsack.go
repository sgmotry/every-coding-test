package utils

import "every-coding-test/internal/model"

func Knapsack(recipes []model.Recipe, maxCalories int, maxCookingTime int) []uint32 {
	recipesLen := len(recipes)

	dp := make([][][]int, recipesLen+1)
	for i := range dp {
		dp[i] = make([][]int, maxCalories+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxCookingTime+1)
		}
	}
	return []uint32{}
}
