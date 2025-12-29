package services

import (
	"math"

	"github.com/javiersrf/mega/entities"
)

func combination(n, k int) float64 {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	k = int(math.Min(float64(k), float64(n-k)))
	num, den := 1.0, 1.0
	for i := 1; i <= k; i++ {
		num *= float64(n - (k - i))
		den *= float64(i)
	}
	return num / den
}

func megaSenaProbability(k int) float64 {
	return combination(k, 6) / combination(60, 6)
}

func computeFinalProbability(items []entities.ResultItem) float64 {
	prod := 1.0
	for _, it := range items {
		p := megaSenaProbability(int(it.Game))
		for i := 0; i < int(it.Quantity); i++ {
			prod *= (1 - p)
		}
	}
	return 1 - prod
}

func CalculateBestCombination(budget float64, games []entities.Game) entities.OptimizationResult {
	cents := int(math.Round(float64(budget * 100)))

	dp := make([]float64, cents+1)
	from := make([]int, cents+1)

	for i := range from {
		from[i] = -1
	}

	for idx, g := range games {
		price := int(math.Round(float64(g.Price * 100)))
		benefit := megaSenaProbability(int(g.Numbers))

		for c := price; c <= cents; c++ {
			if dp[c-price]+benefit > dp[c] {
				dp[c] = dp[c-price] + benefit
				from[c] = idx
			}
		}
	}

	resultMap := map[int16]*entities.ResultItem{}
	var totalAmount float64
	c := cents

	for c > 0 && from[c] != -1 {
		i := from[c]
		g := games[i]
		price := int(math.Round(float64(g.Price * 100)))

		if _, ok := resultMap[g.Numbers]; !ok {
			resultMap[g.Numbers] = &entities.ResultItem{
				Game:     g.Numbers,
				Quantity: 0,
				Amount:   0,
			}
		}

		item := resultMap[g.Numbers]
		item.Quantity++
		item.Amount += g.Price
		totalAmount += g.Price

		c -= price
	}

	items := []entities.ResultItem{}
	for _, v := range resultMap {
		items = append(items, *v)
	}

	finalProb := computeFinalProbability(items)

	return entities.OptimizationResult{
		Items:            items,
		TotalAmount:      totalAmount,
		TotalBenefit:     dp[cents],
		FinalProbability: finalProb,
	}
}
