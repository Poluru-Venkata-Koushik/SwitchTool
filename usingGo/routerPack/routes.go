package routerpack

import (
	"github.com/gin-gonic/gin"
	"SwitchTool/handlers"

)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/" , handlers.RootHandler)
	router.GET("/Welcome", handlers.WelcomeHandler)
	router.GET("/GetAllDevices", handlers.GetAllDevices)
	return router
}
