package main

import (
	"every-coding-test/internal/data"
	"every-coding-test/internal/utils"
	"fmt"
)

func main() {
	data, err := data.GetRecipes()

	if err != nil {
		fmt.Println(err)
	}
	ids, pro, err := utils.GetMaxProteinRecipeCombination(data, -1, 0)
	fmt.Println(ids, pro, err)
}
