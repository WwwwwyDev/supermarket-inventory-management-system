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

func GetSupplier(c *gin.Context)  {
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
	supplierParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchName": searchName,
	}
	err, info, total := services.GetAllSupplier(supplierParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"查询成功")
}

func DelSupplier(c *gin.Context){
	id := -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.DelSupplier(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"删除成功")
}

func UpdateSupplier(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["id"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	id := -1
	id = com.StrTo(m["id"]).MustInt()
	supplierName := m["supplierName"]
	supplierShortname := m["supplierShortname"]
	supplierAddress := m["supplierAddress"]
	supplierCtelephone := m["supplierCtelephone"]
	supplierEmail := m["supplierEmail"]
	supplierLiaisonman := m["supplierLiaisonman"]
	supplierLtelephone := m["supplierLtelephone"]
	supplierRemarks := m["supplierRemarks"]
	err := services.UpdateSupplier(models.Supplier{id,supplierName,supplierShortname,supplierAddress,supplierCtelephone,supplierEmail,supplierLiaisonman,supplierLtelephone,supplierRemarks, false, time.Now()})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"更新成功")
}

func AddSupplier(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	supplierName := m["supplierName"]
	supplierShortname := m["supplierShortname"]
	supplierAddress := m["supplierAddress"]
	supplierCtelephone := m["supplierCtelephone"]
	supplierEmail := m["supplierEmail"]
	supplierLiaisonman := m["supplierLiaisonman"]
	supplierLtelephone := m["supplierLtelephone"]
	supplierRemarks := m["supplierRemarks"]
	err := services.AddSupplier(models.Supplier{-1,supplierName,supplierShortname,supplierAddress,supplierCtelephone,supplierEmail,supplierLiaisonman,supplierLtelephone,supplierRemarks, false, time.Now()})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"添加成功")
}