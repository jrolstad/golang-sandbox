package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/golang-sandbox/gin-webapi/orchestrators"
)

var ginLambda *ginadapter.GinLambda

func init() {
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

	//router.Run()
	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
