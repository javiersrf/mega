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
		prod *= (1 - p)
	}
	return 1 - prod
}

func CalculateBestCombination(budget float64, games []entities.Game) entities.OptimizationResult {
	cents := int(math.Round(float64(budget * 100)))

	// dp[i] = melhor benefício possível com i centavos
	dp := make([]float64, cents+1)
	// keep[i] = índice do jogo escolhido quando atingiu esse valor
	keep := make([][]bool, len(games))
	for i := range keep {
		keep[i] = make([]bool, cents+1)
	}

	for idx, g := range games {
		price := int(math.Round(float64(g.Price * 100)))
		benefit := megaSenaProbability(int(g.Numbers))

		// percorre de trás pra frente (0/1 knapsack)
		for c := cents; c >= price; c-- {
			if dp[c-price]+benefit > dp[c] {
				dp[c] = dp[c-price] + benefit
				keep[idx][c] = true
			}
		}
	}

	// reconstrução do subconjunto
	result := []entities.ResultItem{}
	var totalAmount float64
	c := cents

	for i := len(games) - 1; i >= 0; i-- {
		price := int(math.Round(float64(games[i].Price * 100)))
		if price <= c && keep[i][c] {
			result = append(result, entities.ResultItem{
				Game:     games[i].Numbers,
				Quantity: 1,
				Amount:   games[i].Price,
			})
			totalAmount += games[i].Price
			c -= price
		}
	}

	finalProb := computeFinalProbability(result)

	return entities.OptimizationResult{
		Items:            result,
		TotalAmount:      totalAmount,
		TotalBenefit:     dp[cents],
		FinalProbability: finalProb,
	}
}
