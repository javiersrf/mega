package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckResponse represents the response structure for the health check.
type HealthCheckResponse struct {
	Message string `json:"message"`
}

// HealthCheckHandler godoc
//
//	@Summary		Health Check
//	@Description	Returns a pong message to indicate the service is up
//	@Success		200	{object}	HealthCheckResponse
//	@Router			/health [get]
func HealthCheckHandler(c *gin.Context) {
	response := HealthCheckResponse{
		Message: "ok",
	}
	c.JSON(http.StatusOK, response)
}
