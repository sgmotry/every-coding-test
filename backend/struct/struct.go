package recipe

func NewRecipe(id int, name string, category string, howmany int, time int, descript []string, ingredient []Ingredient, nutrition []Nutrition) Recipe {
	return Recipe{
		id:          id,
		name:        name,
		category:    category,
		howmany:     howmany,
		time:        time,
		descript:    descript,
		ingredients: ingredient,
		nutrition:   nutrition,
	}
}

type Recipe struct {
	id          int          // 識別用ID（プライマリキー）
	name        string       // レシピ名
	category    string       // 分類
	howmany     int          // 何人分か
	time        int          // 調理にかかる時間
	descript    []string     // 作り方の説明
	ingredients []Ingredient // 材料
	nutrition   []Nutrition  // 栄養価
}

type Ingredient struct {
	name   string
	amount string
}
type Nutrition struct {
	name   string
	amount string
}
