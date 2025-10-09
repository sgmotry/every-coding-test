package utils

import (
	"every-coding-test/internal/model"
	"fmt"
)

type OrderByOptions string

const (
	Id          OrderByOptions = "id"
	Name        OrderByOptions = "name"
	Calories    OrderByOptions = "calories"
	CookingTime OrderByOptions = "cookingTime"
)

func validateOrderBy(orderBy OrderByOptions) error {
	switch orderBy {
	case Id, Name, Calories, CookingTime:
		return nil
	default:
		return fmt.Errorf("'%s' is invalid orderby option", orderBy)
	}
}

type OrderOptions string

const (
	Asc  OrderOptions = "asc"
	Desc OrderOptions = "desc"
)

func validateOrder(order OrderOptions) error {
	switch order {
	case Asc, Desc:
		return nil
	default:
		return fmt.Errorf("'%s' is invalid order option", order)
	}
}

func SortRecipe(recipes []model.Recipe, orderBy OrderByOptions, order OrderOptions) []model.Recipe {
	orderByErr := validateOrderBy(orderBy)
	if orderByErr != nil {
		fmt.Println("Error while checking orderBy.\n", orderByErr)
	}
	orderErr := validateOrder(order)
	if orderErr != nil {
		fmt.Println("Error while checking order.\n", orderErr)
	}

	sortedRecipe := quickSort(recipes, orderBy, order)
	return sortedRecipe
}
