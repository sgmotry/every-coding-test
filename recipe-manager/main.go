package main

import (
	"every-coding-test/internal/data"
	"every-coding-test/internal/utils"
	"fmt"
)

func main() {
	data, err := data.GetRecipes()

	if err != nil {
		fmt.Println("Error while getting recipes.\n", err)
	}
	a := utils.SortRecipe(data, "cookingTime", "desc")
	for _, value := range a[0:] {
		fmt.Println(value.Time)
	}
}
