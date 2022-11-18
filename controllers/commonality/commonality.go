package commonality

import (
	"Go-Live/logic/commonality"
	commonalityModel "Go-Live/models/commonality"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
}

//GetOssConfig 获取oss配置
func (c *Controllers) GetOssConfig(ctx *gin.Context) {
	results, err := commonality.GetOssConfig()
	if err != nil {
		response.Error(ctx, err.Error())
	}
	response.Success(ctx, results)
}

//UploadingMethod 获取上传文件配置
func (c *Controllers) UploadingMethod(ctx *gin.Context) {
	UploadingMethodReceive := new(commonalityModel.UploadingMethodStruct)
	if err := ctx.ShouldBind(UploadingMethodReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := commonality.UploadingMethod(UploadingMethodReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
