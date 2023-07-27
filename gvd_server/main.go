package main

import (
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/routers"
)

func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()

	global.Log.Infof("test1")
	global.Log.Warnf("test2")
	global.Log.Errorf("test3")


	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}