package model

func NewRecipe(id int, name string, category string, howmany int, time int, descript []string, ingredient []Ingredient, nutrition []Nutrition) Recipe {
	return Recipe{
		Id:          id,
		Name:        name,
		Category:    category,
		Howmany:     howmany,
		Time:        time,
		Descript:    descript,
		Ingredients: ingredient,
		Nutrition:   nutrition,
	}
}

type Recipe struct {
	Id          int          `json:"id"`          // 識別用ID（プライマリキー）
	Name        string       `json:"name"`        // レシピ名
	Category    string       `json:"category"`    // 分類
	Howmany     int          `json:"howmany"`     // 何人分か
	Time        int          `json:"time"`        // 調理にかかる時間
	Descript    []string     `json:"descript"`    // 作り方の説明
	Ingredients []Ingredient `json:"ingredients"` // 材料
	Nutrition   []Nutrition  `json:"nutrition"`   // 栄養価
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}
type Nutrition struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}
