package handlers

import (
	"net/http"
	"SwitchTool/db"
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
		db.Inventory,
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

