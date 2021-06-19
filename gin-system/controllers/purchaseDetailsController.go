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

func GetPurchaseDetails(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchList := -1
	if arg := c.Query("searchList"); arg != "" {
		searchList = com.StrTo(arg).MustInt()
	}
	searchGoodsName := ""
	if arg := c.Query("searchGoodsName"); arg != "" {
		searchGoodsName = arg
	}
	supplierParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchList": searchList,
		"searchGoodsName" : searchGoodsName,
	}
	err, info, total := services.GetAllPurchaseDetails(supplierParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"查询成功")
}

func DelPurchaseDetails(c *gin.Context){
	id := -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.DelPurchaseDetails(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"删除成功")
}

func UpdatePurchaseDetails(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["id"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	id := -1;
	id = com.StrTo(m["id"]).MustInt()
	purchaseDetailsGoods := -1
	purchaseDetailsGoods = com.StrTo(m["purchaseDetailsGoods"]).MustInt()
	purchaseDetailsNumber := -1
	purchaseDetailsNumber = com.StrTo(m["purchaseDetailsNumber"]).MustInt()
	purchaseDetailsList := -1
	purchaseDetailsList = com.StrTo(m["purchaseDetailsList"]).MustInt()
	purchaseDetailsRemarks := m["purchaseDetailsRemarks"]
	err2, price := services.GetGoodPriceById(purchaseDetailsGoods)
	price = price * purchaseDetailsNumber
	if err2 !=nil{
		app.Error(c,e.ERROR,err2,err2.Error())
		return
	}
	err := services.UpdatePurchaseDetails(models.PuchaseDetails{id,purchaseDetailsList,purchaseDetailsGoods,"",purchaseDetailsNumber,price,purchaseDetailsRemarks,time.Now(),false})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"更新成功")
}

func AddPurchaseDetails(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	purchaseDetailsGoods := -1
	purchaseDetailsGoods = com.StrTo(m["purchaseDetailsGoods"]).MustInt()
	purchaseDetailsNumber := -1
	purchaseDetailsNumber = com.StrTo(m["purchaseDetailsNumber"]).MustInt()
	purchaseDetailsList := -1
	purchaseDetailsList = com.StrTo(m["purchaseDetailsList"]).MustInt()
	purchaseDetailsRemarks := m["purchaseDetailsRemarks"]
	err2, price := services.GetGoodPriceById(purchaseDetailsGoods)
	price = price * purchaseDetailsNumber
	if err2 !=nil{
		app.Error(c,e.ERROR,err2,err2.Error())
		return
	}
	err := services.AddPurchaseDetails(models.PuchaseDetails{-1,purchaseDetailsList,purchaseDetailsGoods,"",purchaseDetailsNumber,price,purchaseDetailsRemarks,time.Now(),false})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"添加成功")
}