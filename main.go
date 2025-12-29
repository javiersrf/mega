package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	docs "github.com/javiersrf/mega/docs"
	"github.com/javiersrf/mega/handlers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/health", handlers.HealthCheckHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.POST("/megasena/calculate", handlers.CalculateHandler)
	}

	r.Run(":" + *port)
}
