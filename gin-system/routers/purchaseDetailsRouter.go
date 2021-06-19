package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func purchaseDetailsRouter(r *gin.RouterGroup) {
	r.GET("/purchaseDetails", controllers.GetPurchaseDetails)
	r.DELETE("/purchaseDetails", controllers.DelPurchaseDetails)
	r.PUT("/purchaseDetails", controllers.UpdatePurchaseDetails)
	r.POST("/purchaseDetails", controllers.AddPurchaseDetails)
}
