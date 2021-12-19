package orchestrators

import (
	"github.com/jrolstad/golang-sandbox/gin-webapi/models"
)

type UserOrchestrator struct{}

func (o UserOrchestrator) Get() []models.User{
	data:=[]models.User{}
	return data
}