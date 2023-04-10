package callback

import (
	"easy-video-net/controllers"
	receive "easy-video-net/interaction/receive/callback"
	"easy-video-net/logic/callback"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	controllers.BaseControllers
}

//AliyunTranscodingMedia 阿里云媒体转码成功回调
func (c *Controllers) AliyunTranscodingMedia(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.AliyunMediaCallback[receive.AliyunTranscodingMediaStruct])); err == nil {
		results, err := callback.AliyunTranscodingMedia(rec)
		c.Response(ctx, results, err)
	}
}
