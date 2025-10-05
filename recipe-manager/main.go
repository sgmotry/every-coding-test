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
	utils.SortRecipe(data, "id", "desc")
}
