package commonality

import (
	"Go-Live/controllers/commonality"
	"github.com/gin-gonic/gin"
)

func (r *RouterGroup) InitRouter(Router *gin.RouterGroup) {
	routers := Router.Group("commonality").Use()
	{
		commonalityControllers := new(commonality.Controllers)
		routers.POST("/ossConfig", commonalityControllers.GetOssConfig)
		routers.POST("/uploadingMethod", commonalityControllers.UploadingMethod)

	}
}
