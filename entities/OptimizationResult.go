package entities

type OptimizationResult struct {
	Items            []ResultItem
	TotalAmount      float64
	TotalBenefit     float64
	FinalProbability float64
}
