package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func goodsRouter(r *gin.RouterGroup) {
	r.GET("/goods", controllers.GetGoods)
	r.DELETE("/goods", controllers.DelGoods)
	//r.PUT("/staff", controllers.UpdateStaff)
	//r.POST("/staff", controllers.AddStaff)
}
