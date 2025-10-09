package model

// レシピ管理用
type Recipe struct {
	Id          uint32       `json:"id"`          // 識別用ID（プライマリキー）
	Name        string       `json:"name"`        // レシピ名
	Category    string       `json:"category"`    // 分類
	Howmany     int          `json:"howmany"`     // 何人分か
	Time        int          `json:"time"`        // 調理にかかる時間（分）
	Descript    []string     `json:"descript"`    // 作り方の説明
	Ingredients []Ingredient `json:"ingredients"` // 材料
	Nutrition   Nutrition    `json:"nutrition"`   // 栄養価
}

// 単位の決まっていない物の量管理用
type Amount struct {
	Number float32 `json:"number"` // 数値
	Unit   string  `json:"unit"`   // 単位
}

// レシピの説明文管理用
type Descript struct {
	Text string `json:"text"` // 説明テキスト
	// Image string `json:"image"` // 将来的に画像データなどを入れる時、ここにURLを
}

// 材料管理用
type Ingredient struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

// レシピの栄養素管理用
type Nutrition struct {
	Calorie        float32          `json:"calorie"`        // カロリー（kcal）
	Protein        float32          `json:"protein"`        // たんぱく質（g）
	Lipids         float32          `json:"lipids"`         // 脂質（g）
	Carbohydrates  float32          `json:"carbohydrates"`  // 炭水化物（g）
	OtherNutrition []OtherNutrition `json:"otherNutrition"` // その他の栄養素
}

// その他の栄養素管理用
type OtherNutrition struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}
