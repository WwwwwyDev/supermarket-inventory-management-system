package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func purchaseListRouter(r *gin.RouterGroup) {
	r.GET("/purchaseList", controllers.GetPurchaseList)
	r.DELETE("/purchaseList", controllers.DelPurchaseList)
	r.PUT("/purchaseList", controllers.UpdatePurchaseList)
	r.POST("/purchaseList", controllers.AddPurchaseList)
}