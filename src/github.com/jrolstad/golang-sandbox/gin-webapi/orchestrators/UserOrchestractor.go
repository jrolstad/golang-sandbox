package orchestrators

import (
	"github.com/jrolstad/golang-sandbox/gin-webapi/models"
	"github.com/jrolstad/golang-sandbox/gin-webapi/repositories"
)

type UserOrchestrator struct{}

func (o UserOrchestrator) Get() []models.User {
	repository := repositories.UserRepository{}
	data := repository.Get()

	return data
}
