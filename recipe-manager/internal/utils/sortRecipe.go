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

// レシピをオプションに従ってソートする
//
// オプション一覧
//
// orderBy（ソートの基準）: "id", "name", "calories", "cookingTime"
//
// order（並び順）: "asc", "desc"
func SortRecipe(recipes []model.Recipe, orderBy OrderByOptions, order OrderOptions) ([]model.Recipe, error) {
	orderByErr := validateOrderBy(orderBy)
	if orderByErr != nil {
		return nil, fmt.Errorf("error while checking orderBy.\n%s", orderByErr)
	}
	orderErr := validateOrder(order)
	if orderErr != nil {
		return nil, fmt.Errorf("error while checking order.\n%s", orderErr)
	}

	sortedRecipe := quickSort(recipes, orderBy, order)
	return sortedRecipe, nil
}
