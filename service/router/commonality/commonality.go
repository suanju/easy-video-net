package commonality

import (
	"easy-video-net/controllers/commonality"
	"easy-video-net/middlewares"

	"github.com/gin-gonic/gin"
)

func (r *RouterGroup) InitRouter(Router *gin.RouterGroup) {
	commonalityControllers := new(commonality.Controllers)
	routers := Router.Group("commonality").Use()
	{
		routers.POST("/ossSTS", commonalityControllers.OssSTS)
		routers.POST("/upload", commonalityControllers.Upload)
		routers.POST("/UploadSlice", commonalityControllers.UploadSlice)
		routers.POST("/uploadCheck", commonalityControllers.UploadCheck)
		routers.POST("/uploadMerge", commonalityControllers.UploadMerge)
		routers.POST("/uploadingMethod", commonalityControllers.UploadingMethod)
		routers.POST("/uploadingDir", commonalityControllers.UploadingDir)
		routers.POST("/getFullPathOfImage", commonalityControllers.GetFullPathOfImage)

	}
	//非必须登入
	contributionRouterNotNecessary := Router.Group("commonality").Use(middlewares.VerificationTokenNotNecessary())
	{
		contributionRouterNotNecessary.POST("/search", commonalityControllers.Search)
	}
}
