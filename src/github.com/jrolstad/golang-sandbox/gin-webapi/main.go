package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrolstad/golang-sandbox/gin-webapi/orchestrators"
)

func main() {
	router := gin.Default()

	apiV1 := router.Group("/v1")

	apiV1.GET("health", func(c *gin.Context) {
		orchestrator := orchestrators.HealthOrchestrator{}

		healthInfo := orchestrator.Get()
		c.JSON(http.StatusOK, healthInfo)
	})

	router.Run()
}
