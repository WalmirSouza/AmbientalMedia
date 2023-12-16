package router

import (
	"github.com/WalmirSouza/Challenge/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouters(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/router", handler.ShowMeteorology)

	}

}
