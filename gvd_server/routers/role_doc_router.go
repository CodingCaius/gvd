package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) RoleDocRouter() {
	app := api.App.RoleDocApi
	r := router.Group("role_docs")
	r.GET(":id", middleware.JwtAdmin(), app.RoleDocListView) // 角色-文档列表
}