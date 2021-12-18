package main

import (
	"fmt"

	"github.com/jrolstad/golang-sandbox/models"
	"github.com/jrolstad/golang-sandbox/routes"
)

func main() {
	fmt.Println("Main package - main file")
	models.AllUsers()
	routes.ApiPostRoute()
}
