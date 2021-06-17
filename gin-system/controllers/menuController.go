package controllers

import (
	"gin-system/pkg/app"
	"gin-system/pkg/e"
	"gin-system/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetMenu(c *gin.Context){
	level := 0
	if arg := c.Query("level"); arg != "" {
		level = com.StrTo(arg).MustInt()
	}
	err, info := services.GetMenuByLevel(level)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"menus": info},"OK")
}
