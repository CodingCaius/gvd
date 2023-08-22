package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) DocRouter() {
	app := api.App.DocApi
	router.POST("docs", middleware.JwtAdmin(), app.DocCreateView) // 创建文档
	//router.GET("docs/info/:id", middleware.JwtAdmin(), app.DocInfoView) // 文档信息
}
