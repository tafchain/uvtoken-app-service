package main

import (
	"service/server/core"
	"service/server/global"
	"service/server/initialize"
)

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_MONGO = initialize.Mongo()
	core.RunWindowsServer()
}
