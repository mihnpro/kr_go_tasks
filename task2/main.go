package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var res []BrainrotMeme

	for _, m := range memes {
		if m.Views > minViews {
			res = append(res, m)
		}
	}
	fmt.Println(res)
	sort.Slice(res, func(i, j int) bool {
		return res[i].TrendLevel > res[j].TrendLevel
	})

	return res
}


func CalculateTrendLevelImpact(memes []BrainrotMeme) map[int]float64 {
	res := make(map[int]float64)
	for _, m := range memes {
		res[m.TrendLevel] += m.Views
	}
	return res
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var res []string
	for _, m := range memes {
		if m.TrendLevel >= 7 || (m.Views > 50 && m.Category == "Sigma") {
			res = append(res, m.Name)
		}
	}
	return res
}


func main(){
	memes := []BrainrotMeme{
		{Name: "Meme 1", TrendLevel: 8, Category: "Sigma", Views: 100},
		{Name: "Meme 2", TrendLevel: 5, Category: "Subo Bratik", Views: 50},
		{Name: "Meme 3", TrendLevel: 9, Category: "Sigma", Views: 200},
		{Name: "Meme 4", TrendLevel: 6, Category: "TUNTUNTUNSAHUR", Views: 75},
		{Name: "Meme 5", TrendLevel: 7, Category: "Sigma", Views: 150},
		{Name: "Meme 6", TrendLevel: 7, Category: "Sigma", Views: 175},
		{Name: "Meme 7", TrendLevel: 8, Category: "Sigma", Views: 125},
	}

	topTrending := FindTopTrending(memes, 100)
	trendLevelImpact := CalculateTrendLevelImpact(memes)
	filteredNames := FilterByComplexCondition(memes)

	fmt.Println("Top Trending Memes:")
	for _, m := range topTrending {
		fmt.Printf("Name: %s, Trend Level: %d, Category: %s, Views: %.2f\n", m.Name, m.TrendLevel, m.Category, m.Views)
	}

	fmt.Println("Trend Level Impact:")
	for trendLevel, impact := range trendLevelImpact {
		fmt.Printf("Trend Level %d: %.2f\n", trendLevel, impact)
	}

	fmt.Println("Filtered Names:")
	for _, name := range filteredNames {
		fmt.Println(name)
	}
}

// Задача #2 Анализ Brainrot-контента

// Вы работаете в отделе анализа интернет-культуры. Ваша задача — проанализировать Brainrot — особый жанр мемов, характеризующийся быстро меняющимися трендами и специфическим контентом.
// Каждый Brainrot-мем описывается структурой с следующими параметрами:

// Name — название мема (строка)

// TrendLevel — уровень трендовости (целое число от 1 до 10, где 10 — максимально трендовый)

// Category — категория контента (строка: "Subo Bratik", "TUNTUNTUNSAHUR", "Sigma", "Skibidi", "Mewing", "Other")

// Views — количество просмотров в миллионах (дробное число)

// Ваша задача:

// 1. Определите структуру BrainrotMeme с указанными полями.

// 2. Создайте срез из минимум 7 различных Brainrot-мемов с разными характеристиками.

// 3. Реализуйте три аналитические функции:
// FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme — возвращает массив мемов, у которых количество просмотров больше minViews, отсортированный по убыванию TrendLevel.

// CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 — возвращает map, где ключ — категория, а значение — суммарное количество просмотров всех мемов этой категории.

// FilterByComplexCondition(memes []BrainrotMeme) []string — возвращает массив названий мемов, которые удовлетворяют сложному условию: TrendLevel ≥ 7 ИЛИ (количество просмотров > 50 И категория "Sigma").
// В функции main() продемонстрируйте работу всех трех функций на вашем наборе данных и выведите результаты в читаемом формате.
