package handlers

import (
	"net/http"
	_ "SwitchTool/DB"
	"github.com/gin-gonic/gin"
)

func RootHandler (ctx *gin.Context){

	ctx.Redirect(
		http.StatusPermanentRedirect,
		"/Welcome",
	)
	
}

func GetAllDevices (ctx *gin.Context){
	ctx.JSON(
		http.StatusOK,
		nil,
	)
}

func WelcomeHandler (ctx *gin.Context){

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message":"Welcome to SwitchTool - Go",
		},	
	)
}

