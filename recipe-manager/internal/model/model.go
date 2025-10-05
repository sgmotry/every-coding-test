package model

func NewRecipe(id uint32, name string, category string, howmany int, time Amount, descript []string, ingredient []Ingredient, nutrition Nutrition) Recipe {
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
	Id          uint32       `json:"id"`          // 識別用ID（プライマリキー）
	Name        string       `json:"name"`        // レシピ名
	Category    string       `json:"category"`    // 分類
	Howmany     int          `json:"howmany"`     // 何人分か
	Time        Amount       `json:"time"`        // 調理にかかる時間
	Descript    []string     `json:"descript"`    // 作り方の説明
	Ingredients []Ingredient `json:"ingredients"` // 材料
	Nutrition   Nutrition    `json:"nutrition"`   // 栄養価
}

type Amount struct {
	Number float32 `json:"number"` // 数値
	Unit   string  `json:"unit"`   // 単位
}

type Descript struct {
	Text string `json:"text"`
	// Image string `json:"image"`
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

type Nutrition struct {
	Calorie        Amount           `json:"calorie"`
	Protein        Amount           `json:"protein"`
	Lipids         Amount           `json:"lipids"`
	Carbohydrates  Amount           `json:"carbohydrates"`
	OtherNutrition []OtherNutrition `json:"otherNutrition"`
}
type OtherNutrition struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}
