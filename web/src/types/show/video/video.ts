import { PageInfo } from "@/types/idnex"

export interface GetVideoContributionByIDReq {
    video_id: number
}

export interface GetVideoContributionByIDRes {
    videoInfo: VInfo
    recommendList: RecommendList
}

interface creatorInfo {
    id: number
    username: string
    avatar: string
    signature: string
    is_attention: boolean
}

export interface VideoInfo {
    videoInfo: VInfo
    recommendList: RecommendList
    barrageList: GetVideoBarrageListRes
}

export interface VInfo {
    id: number
    uid: number
    title: string
    video: string
    cover: string
    video_duration: number
    label: Array<string>
    introduce: string
    heat: number
    barrageNumber: number
    comments: commentsList
    comments_number: number
    is_like: boolean,
    is_collect: boolean,
    creatorInfo: creatorInfo
    created_at: string
}

export interface RecommendViodeInfo {
    id: number
    uid: number
    title: string
    video: string
    cover: string
    video_duration: number
    label: Array<string>
    introduce: string
    heat: number
    barrageNumber: number
    username: string
    created_at: string
}

export type RecommendList = Array<RecommendViodeInfo>


export interface GetVideoBarrageListReq {
    id: string
}

export interface GetVideoBarrageListRes {
    time: number
    text: string
    sendTime: string
}


export interface SendVideoBarrageReq {
    author: string
    color: number
    id: string
    text: string
    time: number
    type: number
    token: string
}


export interface VideoPostCommentReq {
    video_id: number
    content: string
    content_id: number
}

export interface GetVideoCommentReq {
    pageInfo: PageInfo
    video_id: number
}
export interface Comments {
    id: number
    comment_id: number
    created_at: string
    context: string
    comment_user_id: number
    comment_user_name: string
    uid: number
    username: string
    photo: string
    lowerComments: any
}


type commentsList = Array<Comments>

export interface GetVideoCommentRes {
    id: number
    comments: commentsList
    comments_number: number
}

export interface LikeVideoReq {
    video_id: number
}