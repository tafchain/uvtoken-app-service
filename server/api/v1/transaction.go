package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"service/server/global"
	"service/server/model/request"
	"service/server/model/response"
	"service/server/service"
)

func GetBtcHashs(c *gin.Context) {
	var reqInfo request.Hashearch
	_ = c.ShouldBindQuery(&reqInfo)
	if err, hashs := service.GetBtcHash(reqInfo); err != nil {
		global.GVA_LOG.Error("GetBtcHash fail!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(hashs, "获取成功", c)
	}
}

func GetBtcTx(c *gin.Context) {
	var reqInfo request.TxSearch
	_ = c.ShouldBindQuery(&reqInfo)
	if err, list := service.GetBtcTransaction(reqInfo); err != nil {
		global.GVA_LOG.Error("GetBtcTransaction fail!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.ListResult{
			List: list,
		}, "获取成功", c)
	}
}

func GetEthHashs(c *gin.Context) {
	var reqInfo request.Hashearch
	_ = c.ShouldBindQuery(&reqInfo)
	if err, hashs := service.GetEthHash(reqInfo); err != nil {
		global.GVA_LOG.Error("GetEthHash fail!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(hashs, "获取成功", c)
	}
}

func GetEthTx(c *gin.Context) {
	var reqInfo request.TxSearch
	_ = c.ShouldBindQuery(&reqInfo)
	if err, ethTx := service.GetEthTransaction(reqInfo); err != nil {
		global.GVA_LOG.Error("GetEthTransaction fail!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(ethTx, "获取成功", c)
	}
}
