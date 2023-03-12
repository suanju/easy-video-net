package commonality

import (
	"Go-Live/controllers"
	receive "Go-Live/interaction/receive/commonality"
	"Go-Live/logic/commonality"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	controllers.BaseControllers
}

//GetOssConfig 获取oss配置
func (c *Controllers) GetOssConfig(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetOssConfigReceiveStruct)); err == nil {
		results, err := commonality.GetOssConfig(rec)
		c.Response(ctx, results, err)
	}
}

//UploadingMethod 获取上传文件配置
func (c *Controllers) UploadingMethod(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.UploadingMethodStruct)); err == nil {
		results, err := commonality.UploadingMethod(rec)
		c.Response(ctx, results, err)
	}
}

//GetFullPathOfImage 获取图片完整路径
func (c *Controllers) GetFullPathOfImage(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetFullPathOfImageMethodStruct)); err == nil {
		results, err := commonality.GetFullPathOfImage(rec)
		c.Response(ctx, results, err)
	}
}
