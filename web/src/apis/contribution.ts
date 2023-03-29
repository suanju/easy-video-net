import { AreateArticleContributionReq, GetArticleContributionListByUserReq, GetArticleClassificationListRes, getArticleTotalInfoRes, UpdateArticleContributionReq } from "@/types/creation/contribute/contributePage/articleContribution";
import { CreateVideoContributionReq, UpdateVideoContributionReq } from "@/types/creation/contribute/contributePage/vdeoContribution";
import { DeleteArticleByIDReq, GetArticleManagementListReq, GetArticleManagementListRes } from "@/types/creation/manuscript/article";
import { DeleteVideoByIDReq, GetVideoManagementListReq, GetVideoManagementListRes } from "@/types/creation/manuscript/video";
import { GetDiscussArticleListReq, GetDiscussArticleListRes, GetDiscussVideoListReq, GetDiscussVideoListRes } from "@/types/creation/discuss/comment";
import { ArticlePostCommentReq, GetArticleCommentReq, GetArticleCommentRes, GetArticleContributionByIDReq, GetArticleContributionByIDRes } from "@/types/show/article/article";
import { GetVideoContributionByIDReq, GetVideoContributionByIDRes, VideoPostCommentReq, GetVideoBarrageListReq, GetVideoBarrageListRes, SendVideoBarrageReq, GetVideoCommentReq, GetVideoCommentRes, LikeVideoReq } from "@/types/show/video/video";
import { GetArticleContributionListReq, GetArticleContributionListRes } from "@/types/home/column";
import { GetArticleContributionListByUserRes } from "@/types/live/liveRoom";
import httpRequest from "@/utils/requst"
import { GetDiscussBarrageListReq, GetDiscussBarrageListRes } from "@/types/creation/discuss/barrage";


//发布视频
export const createVideoContribution = (params: CreateVideoContributionReq) => {
    return httpRequest.post('/contribution/createVideoContribution', params);
}

//更新视频
export const updateVideoContribution = (params: UpdateVideoContributionReq) => {
    return httpRequest.post('/contribution/updateVideoContribution', params);
}

//发布专栏
export const createArticleContribution = (params: AreateArticleContributionReq) => {
    return httpRequest.post('/contribution/createArticleContribution', params);
}

//更新专栏
export const updateArticleContribution = (params: UpdateArticleContributionReq) => {
    return httpRequest.post('/contribution/updateArticleContribution', params);
}

//查询专栏列表
export const getArticleContributionList = (params: GetArticleContributionListReq) => {
    return httpRequest.post<GetArticleContributionListRes>('/contribution/getArticleContributionList', params);
}
//根据用户获取专栏信息
export const getArticleContributionListByUser = (params: GetArticleContributionListByUserReq) => {
    return httpRequest.post<GetArticleContributionListByUserRes>('/contribution/getArticleContributionListByUser', params);
}

//根据文章ID获取文章信息
export const getArticleContributionByID = (params: GetArticleContributionByIDReq) => {
    return httpRequest.post<GetArticleContributionByIDRes>('/contribution/getArticleContributionByID', params);
}

//文章发布评论
export const articlePostComment = (params: ArticlePostCommentReq) => {
    return httpRequest.post('/contribution/articlePostComment', params);
}

//单独获取文章评论
export const getArticleComment = (params: GetArticleCommentReq) => {
    return httpRequest.post<GetArticleCommentRes>('/contribution/getArticleComment', params);
}

//根据id获取视频信息
export const getVideoContributionByID = (params: GetVideoContributionByIDReq) => {
    return httpRequest.post<GetVideoContributionByIDRes>('/contribution/getVideoContributionByID', params);
}

export const danmakuApi = import.meta.env.VITE_BASE_URL + '/contribution/video/barrage/'

//获取视弹幕列表
export const getVideoBarrageList = (params: GetVideoBarrageListReq) => {
    return httpRequest.get<GetVideoBarrageListRes>('/contribution/getVideoBarrageList', params);
}

//获取视弹幕列表
export const sendVideoBarrage = (params: SendVideoBarrageReq) => {
    return httpRequest.post('/contribution/video/barrage/v3/', params);
}

//文章发布评论
export const videoPostComment = (params: VideoPostCommentReq) => {
    return httpRequest.post('/contribution/videoPostComment', params);
}

//单独获取视频评论
export const getVideoComment = (params: GetVideoCommentReq) => {
    return httpRequest.post<GetVideoCommentRes>('/contribution/getVideoComment', params);
}


//获取文章分类
export const getArticleClassificationList = () => {
    return httpRequest.post<GetArticleClassificationListRes>('/contribution/getArticleClassificationList');
}

//获取文章类目信息
export const getArticleTotalInfo = () => {
    return httpRequest.post<getArticleTotalInfoRes>('/contribution/getArticleTotalInfo');
}

//创作中心获取个人发布视频
export const getVideoManagementList = (params: GetVideoManagementListReq) => {
    return httpRequest.post<GetVideoManagementListRes>('/contribution/getVideoManagementList',params);
}

//根据id删除视频
export const deleteVideoByID = (params: DeleteVideoByIDReq) => {
    return httpRequest.post('/contribution/deleteVideoByID',params);
}

//创作中心获取个人发布专栏
export const getArticleManagementList = (params: GetArticleManagementListReq) => {
    return httpRequest.post<GetArticleManagementListRes>('/contribution/getArticleManagementList',params);
}


//根据id删除专栏
export const deleteArticleByID = (params: DeleteArticleByIDReq) => {
    return httpRequest.post('/contribution/deleteArticleByID',params);
}

//获取评论管理视频评论
export const getDiscussVideoList = (params: GetDiscussVideoListReq) => {
    return httpRequest.post<GetDiscussVideoListRes>('/contribution/getDiscussVideoList',params);
}

//获取评论管理文章评论
export const getDiscussArticleList = (params: GetDiscussArticleListReq) => {
    return httpRequest.post<GetDiscussArticleListRes>('/contribution/getDiscussArticleList',params);
}

//获取弹幕管理
export const getDiscussBarrageList = (params: GetDiscussBarrageListReq) => {
    return httpRequest.post<GetDiscussBarrageListRes>('/contribution/getDiscussBarrageList',params);
}

//获取弹幕管理
export const likeVideo = (params: LikeVideoReq) => {
    return httpRequest.post('/contribution/likeVideo',params);
}
