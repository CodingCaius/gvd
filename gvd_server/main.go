package main

import (
	"gvd_server/core"
	_ "gvd_server/docs"
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
)

// @title 文档项目api文档
// @version 1.0
// @description API文档
// @host 101.43.78.114:8000
// @BasePath /
func main() {
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)

	option := flags.Parse()
	if option.Run() {
		return
	}

	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
