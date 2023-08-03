//添加用户的路由

package routers

import "gvd_server/api"

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi
	router.POST("users", app.UserCreateView)
	router.POST("login", app.UserLoginView)
}
