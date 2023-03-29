package controllers

import (
	"easy-video-net/utils/response"
	"easy-video-net/utils/validator"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BaseControllers struct {
}

//Response 控制器响应输出
func (c BaseControllers) Response(ctx *gin.Context, results interface{}, err error) {
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//ShouldBind 结构体方法无法使用泛型
func ShouldBind[T interface{}](ctx *gin.Context, data T) (t T, err error) {
	if err := ctx.ShouldBind(data); err != nil {
		fmt.Println(err)
		validator.CheckParams(ctx, err)
		return t, err
	}
	return data, nil
}
