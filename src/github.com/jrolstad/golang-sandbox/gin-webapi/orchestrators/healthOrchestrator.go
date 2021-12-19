package orchestrators

import (
	"github.com/jrolstad/golang-sandbox/gin-webapi/models"
)

type HealthOrchestrator struct{}

func (o HealthOrchestrator) Get() models.Health {
	healthInfo := models.Health{}
	healthInfo.Status = "Nothing to see here, move along"

	return healthInfo
}
