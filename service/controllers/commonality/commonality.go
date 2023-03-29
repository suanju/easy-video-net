package commonality

import (
	"easy-video-net/controllers"
	receive "easy-video-net/interaction/receive/commonality"
	"easy-video-net/logic/commonality"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	controllers.BaseControllers
}

//OssSTS 获取oss sts
func (c *Controllers) OssSTS(ctx *gin.Context) {
	results, err := commonality.OssSTS()
	c.Response(ctx, results, err)
}

//Upload  文件上传
func (c *Controllers) Upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	results, err := commonality.Upload(file, ctx)
	c.Response(ctx, results, err)
}

//UploadSlice  文件分片上传
func (c *Controllers) UploadSlice(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	results, err := commonality.UploadSlice(file, ctx)
	c.Response(ctx, results, err)
}

//UploadCheck 验证文件是否已经上传
func (c *Controllers) UploadCheck(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.UploadCheckStruct)); err == nil {
		results, err := commonality.UploadCheck(rec)
		c.Response(ctx, results, err)
	}
}

//UploadMerge 合并分片上传保存
func (c *Controllers) UploadMerge(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.UploadMergeStruct)); err == nil {
		results, err := commonality.UploadMerge(rec)
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

//UploadingDir 获取上传文件保存地址
func (c *Controllers) UploadingDir(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.UploadingDirStruct)); err == nil {
		results, err := commonality.UploadingDir(rec)
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

//Search 搜索接口
func (c *Controllers) Search(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	if rec, err := controllers.ShouldBind(ctx, new(receive.SearchStruct)); err == nil {
		results, err := commonality.Search(rec, uid)
		c.Response(ctx, results, err)
	}
}
