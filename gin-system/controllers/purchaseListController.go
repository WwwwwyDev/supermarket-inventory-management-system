package controllers

import (
	"encoding/json"
	"gin-system/models"
	"gin-system/pkg/app"
	"gin-system/pkg/e"
	"gin-system/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"time"
)

func GetPurchaseList(c *gin.Context)  {
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
	purchaseListParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchName": searchName,
	}
	err, info, total := services.GetAllPurchaseList(purchaseListParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"查询成功")
}


func DelPurchaseList(c *gin.Context){
	id := -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.DelPurchaseList(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"删除成功")
}

func UpdatePurchaseList(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["id"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	id := -1
	id = com.StrTo(m["id"]).MustInt()
	purchaseListStaff := 0
	purchaseListStaff = com.StrTo(m["purchaseListStaff"]).MustInt()
	purchaseListNumber := 0
	purchaseListNumber = com.StrTo(m["purchaseListNumber"]).MustInt()
	purchaseListPrice := 0
	purchaseListPrice = com.StrTo(m["purchaseListPrice"]).MustInt()
	purchaseListRemarks := m["purchaseListRemarks"]
	err := services.UpdatePurchaseList(models.PurchaseList{id, purchaseListStaff, "",purchaseListNumber,purchaseListPrice,time.Time{},purchaseListRemarks,time.Now(),false})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"更新成功")
}

func AddPurchaseList(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	purchaseListStaff := -1
	purchaseListStaff = com.StrTo(m["purchaseListStaff"]).MustInt()
	purchaseListNumber := 0
	purchaseListNumber = com.StrTo(m["purchaseListNumber"]).MustInt()
	purchaseListPrice := 0
	purchaseListPrice = com.StrTo(m["purchaseListPrice"]).MustInt()
	purchaseListRemarks := m["purchaseListRemarks"]
	err := services.AddPurchaseList(models.PurchaseList{-1,purchaseListStaff,"",purchaseListNumber,purchaseListPrice,time.Now(),purchaseListRemarks, time.Now(), false})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"添加成功")
}
