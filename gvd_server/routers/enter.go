package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
  	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// Routers 函数用于创建 Gin 的路由引擎（*gin.Engine）并配置路由
func Routers() *gin.Engine {
	//创建了一个默认的 Gin 路由引擎
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//创建api路由组
	apiGroup := router.Group("api")
	routerGroup := RouterGroup{apiGroup}
	//定义具体的路由
	routerGroup.UserRouter()
	routerGroup.ImageRouter()
	return router
}
