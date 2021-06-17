package v2

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"service/server/global"
	"service/server/model/request"
	"service/server/model/response"
	"service/server/service"
)

func GetEthHashsV2(c *gin.Context) {
	var reqInfo request.Hashearch
	_ = c.ShouldBindQuery(&reqInfo)
	if err, hashs := service.GetEthHashV2(reqInfo); err != nil {
		global.GVA_LOG.Error("GetEthHash fail!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(hashs, "获取成功", c)
	}
}

func GetEthTxV2(c *gin.Context) {
	var reqInfo request.TxSearch
	_ = c.ShouldBindQuery(&reqInfo)
	if err, ethTx := service.GetEthTransactionV2(reqInfo); err != nil {
		global.GVA_LOG.Error("GetEthTransaction fail!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(ethTx, "获取成功", c)
	}
}
