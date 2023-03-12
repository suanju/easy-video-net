package home

import (
	"Go-Live/controllers"
	receive "Go-Live/interaction/receive/home"
	"Go-Live/logic/home"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	controllers.BaseControllers
}

//GetHomeInfo 获取主页信息
func (c Controllers) GetHomeInfo(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.GetHomeInfoReceiveStruct)); err == nil {
		results, err := home.GetHomeInfo(rec)
		c.Response(ctx, results, err)
	}
}
