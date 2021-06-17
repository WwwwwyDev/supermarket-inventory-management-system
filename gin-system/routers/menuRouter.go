package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func menuRouter(r *gin.RouterGroup) {
	r.GET("/menu", controllers.GetMenu)
}
