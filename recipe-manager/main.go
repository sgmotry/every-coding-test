package main

import (
	"every-coding-test/internal/data"
	"fmt"
)

func main() {
	data, err := data.GetRecipes()

	if err != nil {
		fmt.Println("Error while getting recipes.\n", err)
	}
	fmt.Println(data[0].Category)
}
