package routers

import (
	"gin-system/middleWare"
	"github.com/gin-gonic/gin"
)

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/apis")
	staffRouter(r)
	loginRouter(r)
	menuRouter(r)
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleWare.Cors())
	g := r.Group("/system")

	sysNoCheckRoleRouter(g)

	return r
}
