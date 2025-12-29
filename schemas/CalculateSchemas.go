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
	Quantity int32   `json:"quantity"`
	Amount   float64 `json:"amount"`
	Game     int16   `json:"game"`
}

type ResultListResponse struct {
	Games       []ResultItemResponse `json:"games"`
	Probability float64              `json:"probability"`
	TotalAmount float64              `json:"totalAmount"`
}
