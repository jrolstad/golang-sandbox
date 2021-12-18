package main

import (
	"fmt"

	"github.com/jrolstad/golang-sandbox/multi-directory/models"
	"github.com/jrolstad/golang-sandbox/multi-directory/routes"
)

func main() {
	fmt.Println("Main package - main file")
	models.AllUsers()
	routes.ApiPostRoute()
}
