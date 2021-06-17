package controllers

import (
	"encoding/json"
	"gin-system/pkg/app"
	"gin-system/pkg/e"
	"gin-system/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	username := m["username"]
	password := m["password"]

	err, staffData, _ := services.GetStaffByName(username)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	if len(staffData) < 1{
		app.INFO(c, 30000,"没有此用户")
		return
	}
	if staffData[0].StaffPassword != password{
		app.INFO(c, 30001,"密码错误")
		return
	}
	app.OK(c, map[string]interface{}{"user": staffData[0]},"OK")
}
