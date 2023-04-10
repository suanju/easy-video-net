import { PageInfo } from "../idnex";

//获取请求方法
export interface GetUploadingMethodReq {
    method: string
}

export interface GetUploadingMethodRes {
    type: string
}

export interface GetuploadingDirReq {
    interface: string
}
export interface GetuploadingDirRes {
    path: string
    quality: number
}

export interface GetOssConfigdRrq {
    interface: string
}

export interface GteossStsres {
    region: string;
    access_key_id: string;
    access_key_secret: string;
    sts_token: string;
    bucket: string;
    expiration_time: number;
}

export interface GetFullPathOfImageRrq {
    path: string
    type: string
}

export interface SearchReq {
    page_info: PageInfo;
    type: string
}

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

export type SearchRes = any

export interface UserInfo {
    id: number
    username: string
    photo: string
    signature: string
    is_attention: boolean
}

export type UserInfoList = Array<UserInfo>

export interface UploadSliceInfo {
    index: number
    hash: string
}

export type UploadSliceList = Array<UploadSliceInfo>
export interface UploadCheckReq {
    file_md5: string
    interface: string
    slice_list: UploadSliceList
}

export interface UploadCheckRes {
    is_upload: boolean
    list: Array<UploadSliceInfo>
    path: string
}

export interface UploadMergeReq {
    file_name: string
    interface: string
    slice_list: UploadSliceList
}

export type UploadMergeRes = string

export interface RegisterMediaReq {
    type: string
    path: string
}
export type RegisterMediaRes = string