import { PageInfo } from "../idnex"

//getLiveRoom 获取直播接口请求需要
export interface GetLiveRoomRes {
    address: string,
    key: string
}

//轮播图
export interface Rotograph {
    title: string
    cover: string
    color: string
    type: string
    to_id: number
}
export type RotographList = Array<Rotograph>

//视频信息
export interface VideoInfo {
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

export type VideoInfoList = Array<VideoInfo>



export interface GetHomeInfoReq {
    page_info: PageInfo
}

export interface GetHomeInfoRes {
    rotograph: RotographList
    videoList: VideoInfoList
}