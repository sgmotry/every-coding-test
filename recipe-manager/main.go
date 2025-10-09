package main

import (
	"every-coding-test/internal/data"
	"every-coding-test/internal/utils"
	"fmt"
)

func main() {
	data, err := data.GetRecipes()
	if err != nil {
		fmt.Println("JSONデータ読み込み中にエラー\n", err)
	}
	fmt.Println("\n----読み込んだJSONデータ----\n\n", data)

	sortedData, err := utils.SortRecipe(data, "calories", "asc")
	if err != nil {
		fmt.Println("レシピソート中にエラー\n", err)
	}
	fmt.Println("\n----カロリーを基準に昇順にソートしたデータ----\n\n", sortedData)
	fmt.Print("\n----カロリーを基準に昇順にソートしたデータ（カロリーのみ）----\n\n")
	for _, rec := range sortedData {
		fmt.Println(rec.Nutrition.Calorie)
	}

	sortedData, err = utils.SortRecipe(data, "cookingTime", "desc")
	if err != nil {
		fmt.Println("レシピソート中にエラー\n", err)
	}
	fmt.Println("\n----料理時間を基準に降順にソートしたデータ----\n\n", sortedData)
	fmt.Print("\n----料理時間を基準に降順にソートしたデータ（時間のみ）----\n\n")
	for _, rec := range sortedData {
		fmt.Println(rec.Time)
	}

	_, err = utils.SortRecipe(nil, "hoge", "hoge")
	fmt.Print("\n----ソート関数のエラーログ----\n\n")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n----カロリー上限1500、調理時間上限60分でたんぱく質が最大になるレシピの組み合わせ")
	ids, protein, err := utils.GetMaxProteinRecipeCombination(data, 1500, 60)
	if err != nil {
		fmt.Println("ナップサック問題中にエラー\n", err)
	}
	fmt.Println("IDの配列\n", ids)
	fmt.Println("たんぱく質の量(g)\n", protein)

	_, _, err = utils.GetMaxProteinRecipeCombination(nil, -1, -1)
	if err != nil {
		fmt.Print("\n----ナップサック問題のエラーログ----\n\n", err)
	}
}
