package users

import (
	"easy-video-net/controllers/users"
	"easy-video-net/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *LoginRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(middlewares.VerificationToken())
	{
		usersControllers := new(users.UserControllers)
		router.POST("/getUserInfo", usersControllers.GetUserInfo)
		router.POST("/setUserInfo", usersControllers.SetUserInfo)
		router.POST("/determineNameExists", usersControllers.DetermineNameExists)
		router.POST("/updateAvatar", usersControllers.UpdateAvatar)
		router.POST("/getLiveData", usersControllers.GetLiveData)
		router.POST("/saveLiveData", usersControllers.SaveLiveData)
		router.POST("/sendEmailVerificationCodeByChangePassword", usersControllers.SendEmailVerificationCodeByChangePassword)
		router.POST("/changePassword", usersControllers.ChangePassword)
		router.POST("/attention", usersControllers.Attention)
		router.POST("/createFavorites", usersControllers.CreateFavorites)
		router.POST("/getFavoritesList", usersControllers.GetFavoritesList)
		router.POST("/deleteFavorites", usersControllers.DeleteFavorites)
		router.POST("/favoriteVideo", usersControllers.FavoriteVideo)
		router.POST("/getFavoritesListByFavoriteVideo", usersControllers.GetFavoritesListByFavoriteVideo)
		router.POST("/getFavoriteVideoList", usersControllers.GetFavoriteVideoList)
		router.POST("/getRecordList", usersControllers.GetRecordList)
		router.POST("/clearRecord", usersControllers.ClearRecord)
		router.POST("/deleteRecordByID", usersControllers.DeleteRecordByID)
		router.POST("/getNoticeList", usersControllers.GetNoticeList)
		router.POST("/getChatList", usersControllers.GetChatList)
		router.POST("/getChatHistoryMsg", usersControllers.GetChatHistoryMsg)
		router.POST("/personalLetter", usersControllers.PersonalLetter)
		router.POST("/deleteChatItem", usersControllers.DeleteChatItem)
	}
}
