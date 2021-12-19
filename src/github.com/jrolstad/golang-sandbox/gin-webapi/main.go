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

		data := orchestrator.Get()
		c.JSON(http.StatusOK, data)
	})

	apiV1.GET("users", func(c *gin.Context) {
		orchestrator := orchestrators.UserOrchestrator{}

		data := orchestrator.Get()
		c.JSON(http.StatusOK, data)
	})

	router.Run()
}
