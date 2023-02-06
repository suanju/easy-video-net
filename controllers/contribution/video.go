package contribution

import (
	receive "Go-Live/interaction/receive/contribution/video"
	"Go-Live/logic/contribution"
	"Go-Live/logic/contribution/socket"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/websocket"
	"strconv"
)

type Controllers struct {
}

//CreateVideoContribution 发布视频
func (C Controllers) CreateVideoContribution(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	CreateVideoContributionReceive := new(receive.CreateVideoContributionReceiveStruct)
	if err := ctx.ShouldBind(CreateVideoContributionReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.CreateVideoContribution(CreateVideoContributionReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

// GetVideoContributionByID  根据id获取视频信息
func (C Controllers) GetVideoContributionByID(ctx *gin.Context) {
	GetVideoContributionByIDReceive := new(receive.GetVideoContributionByIDReceiveStruct)
	if err := ctx.ShouldBind(GetVideoContributionByIDReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetVideoContributionByID(GetVideoContributionByIDReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

// SendVideoBarrage  发送视频弹幕
func (C Controllers) SendVideoBarrage(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	SendVideoBarrageReceive := new(receive.SendVideoBarrageReceiveStruct)
	if err := ctx.ShouldBindBodyWith(SendVideoBarrageReceive, binding.JSON); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	_, err := contribution.SendVideoBarrage(SendVideoBarrageReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Default(ctx)
}

// GetVideoBarrage  获取视频弹幕 (播放器）
func (C Controllers) GetVideoBarrage(ctx *gin.Context) {
	GetVideoBarrageReceive := new(receive.GetVideoBarrageReceiveStruct)
	GetVideoBarrageReceive.ID = ctx.Query("id")
	results, err := contribution.GetVideoBarrage(GetVideoBarrageReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.BarrageSuccess(ctx, results)
}

// GetVideoBarrageList  获取视频弹幕展示
func (C Controllers) GetVideoBarrageList(ctx *gin.Context) {
	GetVideoBarrageReceive := new(receive.GetVideoBarrageListReceiveStruct)
	GetVideoBarrageReceive.ID = ctx.Query("id")
	results, err := contribution.GetVideoBarrageList(GetVideoBarrageReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.BarrageSuccess(ctx, results)
}

// VideoSocket  观看视频建立的socket
func (C Controllers) VideoSocket(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	conn, _ := ctx.Get("conn")
	ws := conn.(*websocket.Conn)

	//判断是否创建视频socket房间
	id, _ := strconv.Atoi(ctx.Query("videoID"))
	videoID := uint(id)
	if socket.Severe.VideoRoom[videoID] == nil {
		//无人观看主动创建
		socket.Severe.VideoRoom[videoID] = make(socket.UserMapChannel, 10)
	}
	err := socket.CreateVideoSocket(userID, videoID, ws)
	if err != nil {
		response.ErrorWs(ws, "创建socket失败")
	}
}

//VideoPostComment 视频评论
func (C Controllers) VideoPostComment(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	VideoPostCommentReceive := new(receive.VideosPostCommentReceiveStruct)
	if err := ctx.ShouldBind(VideoPostCommentReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.VideoPostComment(VideoPostCommentReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}

//GetVideoComment 获取视频评论
func (C Controllers) GetVideoComment(ctx *gin.Context) {
	GetVideoCommentReceive := new(receive.GetVideoCommentReceiveStruct)
	if err := ctx.ShouldBind(GetVideoCommentReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := contribution.GetVideoComment(GetVideoCommentReceive)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
