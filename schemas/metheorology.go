package schemas

import (
	"time"

	"gorm.io/gorm"
)

type WindForecast struct {
	gorm.Model
	Data  time.Time `json:"data"`
	Vel   float64   `json:"vel"`
	Dir   int       `json:"dir"`
	Alert bool      `json:"alert"`
}
