package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	InitializeRouters(router)

	router.Run(":8080")
}
