package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	docs "github.com/javiersrf/mega/docs"
	"github.com/javiersrf/mega/handlers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	r := gin.Default()
	r.Use(CORSMiddleware())
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/health", handlers.HealthCheckHandler)

	api := r.Group("/api")

	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		v1 := api.Group("/v1")
		{
			v1.POST("/megasena/calculate", handlers.CalculateHandler)
		}
	}

	r.Run(":" + *port)
}
