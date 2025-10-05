package utils

import (
	"every-coding-test/internal/model"
)

func quickSort(data []model.Recipe, orderBy OrderByOptions, order OrderOptions) []model.Recipe {
	if len(data) <= 1 {
		return data
	}

	var left []model.Recipe
	var right []model.Recipe
	var pivotRecipe model.Recipe
	switch orderBy {
	// 最初の２つを比較し大きいほうをピボットに
	case Id:
		var pivot uint32
		var pivotIndex int
		if data[0].Id > data[1].Id {
			pivot = data[0].Id
			pivotIndex = 0
		} else {
			pivot = data[1].Id
			pivotIndex = 1
		}
		pivotRecipe = data[pivotIndex]
		for index, value := range data {
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
		if data[0].Name > data[1].Name {
			pivot = data[0].Name
			pivotIndex = 0
		} else {
			pivot = data[1].Name
			pivotIndex = 1
		}
		pivotRecipe = data[pivotIndex]
		for index, value := range data {
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
		if data[0].Nutrition.Calorie > data[1].Nutrition.Calorie {
			pivot = data[0].Nutrition.Calorie
			pivotIndex = 0
		} else {
			pivot = data[1].Nutrition.Calorie
			pivotIndex = 1
		}
		pivotRecipe = data[pivotIndex]
		for index, value := range data {
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
		if data[0].Time > data[1].Time {
			pivot = data[0].Time
			pivotIndex = 0
		} else {
			pivot = data[1].Time
			pivotIndex = 1
		}
		pivotRecipe = data[pivotIndex]
		for index, value := range data {
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
