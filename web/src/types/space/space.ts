import { type } from 'os';
export interface GetSpaceIndividualReq {
    id :number
}

export interface GetSpaceIndividualRes {
    id :number
    username :string
    photo :string 
    signature  :string
    is_attention :boolean
    attention_num :number
    vermicelli_num : number
}

export interface GetReleaseInformationReq {
    id : number
}

//视频信息
export interface VideoInfo {
    id : number
	uid :number 
	title : string
	video : string
	cover : string
	video_duration : number
	label : Array<string>
	introduce :  string
	heat : number 
	barrageNumber : number
	username : string
    created_at : string
}

export type VideoInfoList = Array<VideoInfo>

//专栏信息
export interface ArticleInfo {
    id: number
    uid: number
    title: string
    cover: string
    classification: string
    label: Array<string>
    content: string
    is_comments: boolean
    heat: number
    likes_number: number
    comments_number: number
    created_at: string
    is_stay: boolean
}

export type ArticleInfoList = Array<ArticleInfo>


export interface GetReleaseInformationRes{
    videoList : VideoInfoList
    articleList : ArticleInfoList
}

export interface GetAttentionListReq{
    id :number
}

export interface GetAttentionListItem{
    id: number 
    is_attention :boolean
    name : string
    photo :string
    signature :string
}

export type GetAttentionListRes = Array<GetAttentionListItem>



export interface GetVermicelliListReq{
    id :number
}

export interface GetVermicelliListItem{
    id: number 
    is_attention :boolean
    name : string
    photo :string
    signature :string
}

export type GetVermicelliListRes = Array<GetVermicelliListItem>