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
	getOssConfigReceive := new(commonalityModel.GetOssConfigReceiveStruct)
	if err := ctx.ShouldBind(getOssConfigReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := commonality.GetOssConfig(getOssConfigReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
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

//GetFullPathOfImage 获取图片完整路径
func (c *Controllers) GetFullPathOfImage(ctx *gin.Context) {
	getFullPathOfImageMethodReceive := new(commonalityModel.GetFullPathOfImageMethodStruct)
	if err := ctx.ShouldBind(getFullPathOfImageMethodReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := commonality.GetFullPathOfImage(getFullPathOfImageMethodReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
