package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func goodsRouter(r *gin.RouterGroup) {
	r.GET("/goods", controllers.GetGoods)
	r.DELETE("/goods", controllers.DelGoods)
	r.PUT("/goods", controllers.UpdateGoods)
	r.POST("/goods", controllers.AddGoods)
}
