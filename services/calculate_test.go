package services

import (
	"math"
	"testing"

	"github.com/javiersrf/mega/entities"
)

func almostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestCalculateBestCombination_SelectsBestNonRepeatedGames(t *testing.T) {
	budget := float64(200)

	games := []entities.Game{
		{Numbers: 6, Price: 5.0},
		{Numbers: 7, Price: 35.0},
		{Numbers: 8, Price: 140.0},
		{Numbers: 9, Price: 300.0}, // não cabe
	}

	result := CalculateBestCombination(budget, games)

	if len(result.Items) == 0 {
		t.Fatal("expected at least one selected game, got none")
	}

	// garantir que não há repetição
	for _, item := range result.Items {
		if item.Quantity != 1 {
			t.Errorf("expected Quantity=1 (no repetition), got %d", item.Quantity)
		}
	}

	// garantir que o gasto não passou do orçamento
	if result.TotalAmount > budget+0.01 {
		t.Errorf("total amount exceeded budget: %.2f > %.2f", result.TotalAmount, budget)
	}

	// probabilidade final deve estar no intervalo correto
	if result.FinalProbability <= 0 || result.FinalProbability >= 1 {
		t.Errorf("invalid final probability: %v", result.FinalProbability)
	}
}

func TestCalculateBestCombination_PrefersBetterProbabilityOverCheaperGame(t *testing.T) {
	budget := float64(140)

	games := []entities.Game{
		{Numbers: 6, Price: 5.0},   // barato mas prob menor
		{Numbers: 8, Price: 140.0}, // caro mas prob maior
	}

	result := CalculateBestCombination(budget, games)

	if len(result.Items) != 1 {
		t.Fatalf("expected exactly 1 selected game, got %d", len(result.Items))
	}

	if result.Items[0].Game != 8 {
		t.Errorf("expected game 8 to be selected, got %d", result.Items[0].Game)
	}
}

func TestCalculateBestCombination_NoGameFitsBudget(t *testing.T) {
	budget := float64(10)

	games := []entities.Game{
		{Numbers: 7, Price: 35.0},
		{Numbers: 8, Price: 140.0},
	}

	result := CalculateBestCombination(budget, games)

	if len(result.Items) != 0 {
		t.Errorf("expected no games to be selected, got %d", len(result.Items))
	}

	if result.FinalProbability != 0 {
		t.Errorf("expected probability = 0 when no games selected, got %v", result.FinalProbability)
	}
}

func TestComputeFinalProbabilityMatchesManualCalculation(t *testing.T) {
	items := []entities.ResultItem{
		{Game: 6, Quantity: 1},
		{Game: 7, Quantity: 1},
	}

	p6 := megaSenaProbability(6)
	p7 := megaSenaProbability(7)

	expected := 1 - ((1 - p6) * (1 - p7))
	got := computeFinalProbability(items)

	if !almostEqual(expected, got, 1e-12) {
		t.Errorf("final probability mismatch: got %v, expected %v", got, expected)
	}
}
