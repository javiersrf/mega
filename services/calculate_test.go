package services

import (
	"math"
	"testing"

	"github.com/javiersrf/mega/entities"
)

func almostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestCalculateBestCombinationWithAtLeast_RespectsAtLeastAndBudget(t *testing.T) {
	budget := float64(200)

	games := []entities.Game{
		{Numbers: 6, Price: 5.0, AtLeast: 2},  // deve garantir 2 jogos
		{Numbers: 7, Price: 35.0, AtLeast: 0}, // opcional
		{Numbers: 8, Price: 140.0},
	}

	result, err := CalculateBestCombinationWithAtLeast(budget, games)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Items) == 0 {
		t.Fatal("expected at least one selected game, got none")
	}

	// garante que o campo AtLeast foi respeitado
	found6 := false
	for _, item := range result.Items {
		if item.Game == 6 {
			found6 = true
			if item.Quantity < 2 {
				t.Errorf("expected at least 2 games of 6 numbers, got %d", item.Quantity)
			}
		}
	}

	if !found6 {
		t.Errorf("expected game type 6 to appear due to AtLeast rule")
	}

	// gasto não pode passar orçamento
	if result.TotalAmount > budget+0.01 {
		t.Errorf("total amount exceeded budget: %.2f > %.2f", result.TotalAmount, budget)
	}

	// probabilidade deve estar no intervalo válido
	if result.FinalProbability <= 0 || result.FinalProbability > 1 {
		t.Errorf("invalid final probability: %v", result.FinalProbability)
	}
}

func TestCalculateBestCombinationWithAtLeast_PrefersHigherTotalProbability(t *testing.T) {
	budget := float64(140)

	games := []entities.Game{
		{Numbers: 6, Price: 5.0},
		{Numbers: 8, Price: 140.0},
	}

	result, _ := CalculateBestCombinationWithAtLeast(budget, games)

	if len(result.Items) == 0 {
		t.Fatal("expected at least one game, got none")
	}

	gotProb := result.FinalProbability

	altProb := gameWinProbability(8)

	if gotProb < altProb {
		t.Errorf("algorithm did not maximize probability: got=%v expected>=%v",
			gotProb, altProb)
	}
}

func TestCalculateBestCombinationWithAtLeast_NoGameFitsBudget(t *testing.T) {
	budget := float64(10)

	games := []entities.Game{
		{Numbers: 7, Price: 35.0},
		{Numbers: 8, Price: 140.0},
	}

	result, _ := CalculateBestCombinationWithAtLeast(budget, games)

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

	p6 := gameWinProbability(6)
	p7 := gameWinProbability(7)

	expected := 1 - ((1 - p6) * (1 - p7))
	got := computeFinalProbability(items)

	if !almostEqual(expected, got, 1e-12) {
		t.Errorf("final probability mismatch: got %v, expected %v", got, expected)
	}
}
