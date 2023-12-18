package router

import (
	"github.com/WalmirSouza/Challenge/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouters(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/forecast?latitude=-27.6289&longitude=-48.4478&past_days=7&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m", handler.ShowMeteorology)

	}

}
