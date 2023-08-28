package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) RoleDocRouter() {
	app := api.App.RoleDocApi
	r := router.Group("role_docs")
	r.GET(":id", middleware.JwtAdmin(), app.RoleDocListView) // 角色-文档列表
	r.POST("", middleware.JwtAdmin(), app.RoleDocCreateView)   // 角色-文档 添加
	//r.DELETE("", middleware.JwtAdmin(), app.RoleDocRemoveView) // 角色-文档 删除
	//r.GET("info", middleware.JwtAdmin(), app.RoleDocInfoView)  // 角色-文档 信息
}