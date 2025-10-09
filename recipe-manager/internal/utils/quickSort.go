package utils

import (
	"every-coding-test/internal/model"
)

func quickSort(recipes []model.Recipe, orderBy OrderByOptions, order OrderOptions) []model.Recipe {
	if len(recipes) <= 1 {
		return recipes
	}

	var left []model.Recipe
	var right []model.Recipe
	var pivotRecipe model.Recipe
	switch orderBy {
	// 最初の２つを比較し大きいほうをピボットに
	case Id:
		var pivot uint32
		var pivotIndex int
		if recipes[0].Id > recipes[1].Id {
			pivot = recipes[0].Id
			pivotIndex = 0
		} else {
			pivot = recipes[1].Id
			pivotIndex = 1
		}
		pivotRecipe = recipes[pivotIndex]
		for index, value := range recipes {
			if index == pivotIndex {
				continue
			}

			isSmallerThanPivot := value.Id < pivot
			// オプションがDescの場合は判定を反転
			if order == Desc {
				isSmallerThanPivot = !isSmallerThanPivot
			}

			if isSmallerThanPivot {
				left = append(left, value)
			} else {
				right = append(right, value)
			}
		}
	case Name:
		var pivot string
		var pivotIndex int
		if recipes[0].Name > recipes[1].Name {
			pivot = recipes[0].Name
			pivotIndex = 0
		} else {
			pivot = recipes[1].Name
			pivotIndex = 1
		}
		pivotRecipe = recipes[pivotIndex]
		for index, value := range recipes {
			if index == pivotIndex {
				continue
			}

			isSmallerThanPivot := value.Name < pivot
			if order == Desc {
				isSmallerThanPivot = !isSmallerThanPivot
			}

			if isSmallerThanPivot {
				left = append(left, value)
			} else {
				right = append(right, value)
			}
		}
	case Calories:
		var pivot float32
		var pivotIndex int
		if recipes[0].Nutrition.Calorie > recipes[1].Nutrition.Calorie {
			pivot = recipes[0].Nutrition.Calorie
			pivotIndex = 0
		} else {
			pivot = recipes[1].Nutrition.Calorie
			pivotIndex = 1
		}
		pivotRecipe = recipes[pivotIndex]
		for index, value := range recipes {
			if index == pivotIndex {
				continue
			}

			isSmallerThanPivot := value.Nutrition.Calorie < pivot
			if order == Desc {
				isSmallerThanPivot = !isSmallerThanPivot
			}

			if isSmallerThanPivot {
				left = append(left, value)
			} else {
				right = append(right, value)
			}
		}
	case CookingTime:
		var pivot int
		var pivotIndex int
		if recipes[0].Time > recipes[1].Time {
			pivot = recipes[0].Time
			pivotIndex = 0
		} else {
			pivot = recipes[1].Time
			pivotIndex = 1
		}
		pivotRecipe = recipes[pivotIndex]
		for index, value := range recipes {
			if index == pivotIndex {
				continue
			}

			isSmallerThanPivot := value.Time < pivot
			if order == Desc {
				isSmallerThanPivot = !isSmallerThanPivot
			}

			if isSmallerThanPivot {
				left = append(left, value)
			} else {
				right = append(right, value)
			}
		}
	}

	left = quickSort(left, orderBy, order)
	right = quickSort(right, orderBy, order)

	var sortedRecipe []model.Recipe
	sortedRecipe = append(left, pivotRecipe)
	sortedRecipe = append(sortedRecipe, right...)

	return sortedRecipe
}
