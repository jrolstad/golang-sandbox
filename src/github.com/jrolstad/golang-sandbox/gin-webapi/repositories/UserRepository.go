package repositories

import (
	"github.com/jrolstad/golang-sandbox/gin-webapi/models"
)

type UserRepository struct{}

func (r *UserRepository) Get() []models.User{
	data:=[]models.User{}
	return data
}