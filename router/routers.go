package router

import (
	"github.com/WalmirSouza/Challenge/handler"
	"github.com/gin-gonic/gin"
)

// $ curl "https://api.open-meteo.com/v1/forecast?latitude=-27.6289&longitude=-48.4478&past_days=7&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"

//	{
//	  "hourly": {
//	    "time": ["2022-06-19T00:00","2022-06-19T01:00", ...]
//	    "wind_speed_10m": [3.16,3.02,3.3,3.14,3.2,2.95, ...],
//	    "temperature_2m": [13.7,13.3,12.8,12.3,11.8, ...],
//	    "relative_humidity_2m": [82,83,86,85,88,88,84,76, ...],
//	  }
//	}
func InitializeRouters(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/forecast?latitude=-27.6289&longitude=-48.4478&past_days=7&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m", handler.ShowMeteorology)

	}

}
