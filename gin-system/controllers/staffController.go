package controllers

import (
	"encoding/json"
	"gin-system/dao"
	"gin-system/models"
	"gin-system/pkg/app"
	"gin-system/pkg/e"
	"gin-system/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"time"
)

func GetStaff(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchName := ""
	if arg := c.Query("searchName"); arg != "" {
		searchName = arg
	}
	staffParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchName": searchName,
	}
	err, info, total := services.GetAllStaff(staffParam)
	//fmt.Println(info);
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"查询成功")
}

func DelStaff(c *gin.Context){
	id := -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1{
		app.INFO(c,30001,"参数错误")
		return
	}
	err2, staff := services.GetStaffById(id)
	if err2 !=nil{
		app.Error(c,e.ERROR,err2,err2.Error())
		return
	}
	if staff.StaffLevel == 2{
		app.INFO(c,30000,"管理员用户不能删除")
		return
	}
	err := services.DelStaff(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"删除成功")
}

func UpdateStaff(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["id"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	id := com.StrTo(m["id"]).MustInt()
	staffName := m["staffName"]
	staffPassword := m["staffPassword"]
	staffLevel := com.StrTo(m["staffLevel"]).MustInt()
	staffTelephone := m["staffTelephone"]
	staffSalary := com.StrTo(m["staffSalary"]).MustInt()
	staffRemarks := m["staffRemarks"]
	err := dao.UpdateStaff(models.Staff{id,staffName,staffPassword,staffLevel,staffTelephone,staffSalary,staffRemarks,false,time.Now()})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"更新成功")
}

func AddStaff(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	staffName := m["staffName"]
	staffPassword := m["staffPassword"]
	staffLevel := com.StrTo(m["staffLevel"]).MustInt()
	staffTelephone := m["staffTelephone"]
	staffSalary := com.StrTo(m["staffSalary"]).MustInt()
	staffRemarks := m["staffRemarks"]
	err := dao.AddStaff(models.Staff{-1,staffName,staffPassword,staffLevel,staffTelephone,staffSalary,staffRemarks,false,time.Now()})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"添加成功")
}