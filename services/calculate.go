package services

import (
	"errors"
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
	num := 1.0
	den := 1.0
	for i := 0; i < k; i++ {
		num *= float64(n - i)
		den *= float64(i + 1)
	}
	return num / den
}

func gameWinProbability(numbers int16) float64 {
	totalComb := combination(60, 6)
	gameComb := combination(int(numbers), 6)
	return gameComb / totalComb
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

func CalculateBestCombinationWithAtLeast(
	budget float64,
	games []entities.Game,
) (entities.OptimizationResult, error) {

	result := entities.OptimizationResult{}
	remaining := budget
	probability := 0.0
	gameCount := make(map[int16]int32)

	for _, g := range games {
		if g.AtLeast > 0 {
			requiredCost := float64(g.AtLeast) * g.Price
			if requiredCost > remaining {
				return result, errors.New("budget insufficient to satisfy AtLeast constraints")
			}

			gameCount[g.Numbers] += int32(g.AtLeast)
			remaining -= requiredCost
			probability += float64(g.AtLeast) * gameWinProbability(g.Numbers)
		}
	}

	for remaining > 0 {
		bestIdx := -1
		bestScore := -1.0

		for i, g := range games {
			if g.Price <= remaining {
				score := gameWinProbability(g.Numbers) / g.Price
				if score > bestScore {
					bestScore = score
					bestIdx = i
				}
			}
		}

		if bestIdx == -1 {
			break
		}

		gameCount[games[bestIdx].Numbers]++
		remaining -= games[bestIdx].Price
		probability += gameWinProbability(games[bestIdx].Numbers)
	}

	// 3️⃣ Monta o resultado final
	totalAmount := budget - remaining

	for _, g := range games {
		if q, ok := gameCount[g.Numbers]; ok && q > 0 {
			result.Items = append(result.Items, entities.ResultItem{
				Quantity: q,
				Amount:   float64(q) * g.Price,
				Game:     g.Numbers,
			})
		}
	}

	result.FinalProbability = math.Min(probability, 1)
	result.TotalAmount = totalAmount

	return result, nil
}
