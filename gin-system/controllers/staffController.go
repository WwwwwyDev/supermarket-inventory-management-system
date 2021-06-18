package controllers

import (
	"gin-system/pkg/app"
	"gin-system/pkg/e"
	"gin-system/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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