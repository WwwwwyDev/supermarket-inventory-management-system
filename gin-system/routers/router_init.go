package routers

import "github.com/gin-gonic/gin"

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/apis")
	staffRouter(r)
}

func InitRouter() *gin.Engine {
	r := gin.New()

	g := r.Group("/system")

	sysNoCheckRoleRouter(g)

	return r
}
