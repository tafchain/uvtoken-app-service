package initialize

import (
	"github.com/gin-gonic/gin"
	"service/server/global"
	"service/server/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 打开就能玩https了
	// Router.Use(middleware.LoadTls())
	// 跨域
	// Router.Use(middleware.Cors())

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("wallet")
	PublicGroup.Use().Use()
	{
		router.InitTxRouter(PublicGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
