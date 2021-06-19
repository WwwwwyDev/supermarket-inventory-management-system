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

func GetGoods(c *gin.Context)  {
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
	err, info, total := services.GetAllGoods(supplierParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"查询成功")
}

func DelGoods(c *gin.Context){
	id := -1
	if arg := c.Query("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.DelGoods(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"删除成功")
}

func UpdateGoods(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	if m["id"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	id := com.StrTo(m["id"]).MustInt()
	goodsName := m["goodsName"]
	goodsPrice := com.StrTo(m["goodsPrice"]).MustInt()
	goodsSupplier := com.StrTo(m["goodsSupplier"]).MustInt()
	goodsSynopsis := m["goodsSynopsis"]
	goodsRemarks := m["goodsRemarks"]
	err := services.UpdateGoods(models.Goods{id,goodsName,goodsPrice,goodsSupplier,"",goodsSynopsis,goodsRemarks, false, time.Now()})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"更新成功")
}

func AddGoods(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(b, &m)
	goodsName := m["goodsName"]
	goodsPrice := com.StrTo(m["goodsPrice"]).MustInt()
	goodsSupplier := com.StrTo(m["goodsSupplier"]).MustInt()
	goodsSynopsis := m["goodsSynopsis"]
	goodsRemarks := m["goodsRemarks"]
	err := services.AddGoods(models.Goods{-1,goodsName,goodsPrice,goodsSupplier,"",goodsSynopsis,goodsRemarks, false, time.Now()})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"添加成功")
}