package routers

import (
	"gin-system/controllers"
	"github.com/gin-gonic/gin"
)

func loginRouter(r *gin.RouterGroup) {
	r.GET("/login", controllers.GetStaff)
}

