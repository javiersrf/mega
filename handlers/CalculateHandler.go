package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javiersrf/mega/entities"
	"github.com/javiersrf/mega/schemas"
	"github.com/javiersrf/mega/services"
)

// @BasePath /api/v1

// CalculateHandler handles the calculation of game results based on the provided budget and game data.
//
//	@Summary		Calculate game results
//	@Description	This endpoint receives a JSON request containing a budget and a list of games,
//
//	processes the data, and returns the calculated results including the quantity,
//	amount, game details, and probability for each game.
//
//	@Tags			games
//	@Accept			json
//	@Produce		json
//	@Param			request	body		schemas.CalculateRequest		true	"Calculate request containing budget and games"
//	@Success		200		{object}	schemas.ResultListResponse		"Successful calculation of game results"
//	@Failure		400		{object}	schemas.ErrorResponse	"Bad request - invalid JSON or missing required fields"
//	@Router			/megasena/calculate [post]
func CalculateHandler(c *gin.Context) {

	var req schemas.CalculateRequest
	responseError := schemas.ErrorResponse{}

	if err := c.ShouldBindJSON(&req); err != nil {
		responseError.Error = err.Error()
		c.JSON(http.StatusBadRequest, responseError)
		return
	}
	convertedGames := make([]entities.Game, 0)
	for _, value := range req.Games {
		convertedGames = append(convertedGames, entities.Game{
			Numbers: value.Numbers,
			Price:   value.Price,
		})

	}
	output := services.CalculateBestCombination(req.Budget, convertedGames)

	responseGames := make([]schemas.ResultItemResponse, 0)
	for _, value := range output.Items {
		responseGames = append(responseGames, schemas.ResultItemResponse{
			Quantity: value.Quantity,
			Amount:   value.Amount,
			Game:     value.Game,
		})
	}

	response := schemas.ResultListResponse{
		Games:       responseGames,
		Probability: output.FinalProbability,
	}

	c.JSON(http.StatusOK, response)
}
