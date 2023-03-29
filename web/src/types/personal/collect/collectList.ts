export interface GetFavoriteVideoListReq {
    favorite_id : number
}


//视频信息
export interface VideoInfo {
    id : number
	uid :number 
	title : string
	video : string
	cover : string
    video_duration : number
    created_at : string
}

export type VideoInfoList = Array<VideoInfo>

export interface GetFavoriteVideoListRes {
    videoList :VideoInfoList
}