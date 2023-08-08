package api

import (
	"gvd_server/api/image_api"
	"gvd_server/api/user_api"
)

//便于组织 api 模块的结构体
/*
为了将 api 模块组织成一个整体，
并且在应用程序中只需要一个全局变量来管理所有的 api 功能
*/
type Api struct {
	UserApi  user_api.UserApi
	ImageApi image_api.ImageApi
}

var App = new(Api)
