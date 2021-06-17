package router

import (
	"github.com/gin-gonic/gin"
	v1 "service/server/api/v1"
	v2 "service/server/api/v2"
)

func InitTxRouter(Router *gin.RouterGroup) {
	TxRouter := Router.Group("transaction").Use()
	{
		TxRouter.GET("btchashs", v1.GetBtcHashs)
		TxRouter.GET("btc", v1.GetBtcTx)
		TxRouter.GET("ethhashs", v1.GetEthHashs)
		TxRouter.GET("eth", v1.GetEthTx)
		// v2
		TxRouter.GET("btc/v2", v1.GetBtcTx)
		TxRouter.GET("ethhashs/v2", v2.GetEthHashsV2)
		TxRouter.GET("eth/v2", v2.GetEthTxV2)
	}
}
