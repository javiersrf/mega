package schemas

type GameRequest struct {
	Numbers int16   `json:"numbers" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

type CalculateRequest struct {
	Budget float64       `json:"budget" binding:"required"`
	Games  []GameRequest `json:"games" binding:"required"`
}

type ResultItemResponse struct {
	Quantity int32
	Amount   float64
	Game     int16
}

type ResultListResponse struct {
	Games       []ResultItemResponse
	Probability float64
}
