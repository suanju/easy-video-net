import { DeleteChatItemReq, GetChatHistoryMsgReq, GetChatHistoryMsgRes, GetChatListRes, PersonalLetterReq } from "@/types/personal/chat/chat";
import { GetFavoriteVideoListReq, GetFavoriteVideoListRes } from "@/types/personal/collect/collectList";
import { createFavoritesReq, SaveCreateFavoritesDataReq } from "@/types/personal/collect/createFavorites";
import { DeleteFavoritesReq, FavoriteVideoReq, GetFavoritesListByFavoriteVideoReq, getFavoritesListByFavoriteVideoRes, GetFavoritesListRes } from "@/types/personal/collect/favorites";
import { GetLiveDataRes, SaveLiveDataReq } from "@/types/personal/live/setUp";
import { GetNoticeListReq, GetNoticeListRes } from "@/types/personal/notice/notice";
import { DeleteRecordByIDReq, GetRecordListReq, GetRecordListRes } from "@/types/personal/record/record";
import { changePasswordReq } from "@/types/personal/userInfo/security";
import { AttentionReq, DetermineNameExistsReq, DetermineNameExistsRes, SetUserInfoRes, UpdateAvatarReq, UserInfoRes } from "@/types/personal/userInfo/userInfo";
import httpRequest from "@/utils/requst";
//获取用户信息
export const getUserInfoRequist = () => {
    return httpRequest.post<UserInfoRes>('/user/getUserInfo');
}
//判断用户名是否存在
export const determineNameExistsRequist = (params: DetermineNameExistsReq) => {
    return httpRequest.post<DetermineNameExistsRes>('/user/determineNameExists', params);
}
//设置用户信息
export const setUserInfoRequist = (params: UserInfoRes) => {
    return httpRequest.post<SetUserInfoRes>('/user/setUserInfo', params);
}
//更新用户头像
export const updateAvatarRequist = (params: UpdateAvatarReq) => {
    return httpRequest.post('/user/updateAvatar', params);
}
//获取用户直播信息
export const getLiveDataRequist = () => {
    return httpRequest.post<GetLiveDataRes>('/user/getLiveData');
}
//设置直播信息
export const saveLiveDataRequist = (params: SaveLiveDataReq) => {
    return httpRequest.post('/user/saveLiveData', params);
}

//修改密码
export const changePassword = (params: changePasswordReq) => {
    return httpRequest.post('/user/changePassword', params);
}
//修改密码发送验证码
export const sendEmailVerificationCodeByChangePassword = () => {
    return httpRequest.post('/user/sendEmailVerificationCodeByChangePassword');
}
//关注
export const attention = (params: AttentionReq) => {
    return httpRequest.post('/user/attention', params);
}
//创建收藏夹
export const createFavorites = (params: SaveCreateFavoritesDataReq | createFavoritesReq) => {
    return httpRequest.post('/user/createFavorites', params);
}
//获取收藏夹
export const getFavoritesList = () => {
    return httpRequest.post<GetFavoritesListRes>('/user/getFavoritesList');
}
//删除收藏夹
export const deleteFavorites = (params: DeleteFavoritesReq) => {
    return httpRequest.post('/user/deleteFavorites', params);
}
//收藏视频
export const favoriteVideo = (params: FavoriteVideoReq) => {
    return httpRequest.post('/user/favoriteVideo', params);
}

//获取收藏夹在视频页面页面需要视频id
export const getFavoritesListByFavoriteVideo = (params: GetFavoritesListByFavoriteVideoReq) => {
    return httpRequest.post<getFavoritesListByFavoriteVideoRes>('/user/getFavoritesListByFavoriteVideo', params);
}
//获取收藏夹内的视频
export const getFavoriteVideoList = (params: GetFavoriteVideoListReq) => {
    return httpRequest.post<GetFavoriteVideoListRes>('/user/getFavoriteVideoList', params);
}

//获取历史记录
export const getRecordList = (params: GetRecordListReq) => {
    return httpRequest.post<GetRecordListRes>('/user/getRecordList', params);
}

//清空历史记录
export const clearRecord = () => {
    return httpRequest.post('/user/clearRecord');
}
//删除历史记录
export const deleteRecordByID = (params: DeleteRecordByIDReq) => {
    return httpRequest.post('/user/deleteRecordByID', params);
}
//获取消息通知
export const getNoticeList = (params: GetNoticeListReq) => {
    return httpRequest.post<GetNoticeListRes>('/user/getNoticeList', params);
}
//获取最近聊天列表
export const getChatList = () => {
    return httpRequest.post<GetChatListRes>('/user/getChatList');
}
//点击私信时触发的操作
export const personalLetter = (params: PersonalLetterReq) => {
    return httpRequest.post('/user/personalLetter', params);
}
//删除最近聊天记录
export const deleteChatItem = (params: DeleteChatItemReq) => {
    return httpRequest.post('/user/deleteChatItem', params);
}
//加载历史聊天记录
export const getChatHistoryMsg = (params: GetChatHistoryMsgReq) => {
    return httpRequest.post<GetChatHistoryMsgRes>('/user/getChatHistoryMsg', params);
}