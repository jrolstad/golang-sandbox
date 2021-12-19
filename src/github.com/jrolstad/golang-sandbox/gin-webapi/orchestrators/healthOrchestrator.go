package orchestrators

import (
	"github.com/jrolstad/golang-sandbox/gin-webapi/models"
)

func GetHealth() models.Health {
	healthInfo := models.Health{}
	healthInfo.Status = "A-Ok"
	
	return healthInfo
}
