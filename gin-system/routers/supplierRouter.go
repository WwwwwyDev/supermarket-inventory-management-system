package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func supplierRouter(r *gin.RouterGroup) {
	r.GET("/supplier", controllers.GetSupplier)
	r.DELETE("/supplier", controllers.DelSupplier)
	r.PUT("/supplier", controllers.UpdateSupplier)
	r.POST("/supplier", controllers.AddSupplier)
}

