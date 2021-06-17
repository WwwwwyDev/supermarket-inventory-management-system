package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func staffRouter(r *gin.RouterGroup) {
	r.GET("/staff", controllers.GetStaff)
}
