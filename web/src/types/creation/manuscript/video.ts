import { PageInfo } from "@/types/idnex"

export interface GetVideoManagementListReq {
    page_info: PageInfo
}


export  interface GetVideoManagementListItem {
    id : number
	uid : number
	title : string
	video : string
	cover : string
	reprinted : boolean,
	cover_url : string
	cover_upload_type : string
	video_duration : number
	label : Array<string>
	introduce : string
 	heat : number
	barrageNumber : number
	comments_number :number
	created_at : string

	is_delete : boolean //用于伪删除
}

export type GetVideoManagementListRes = Array<GetVideoManagementListItem>

export interface DeleteVideoByIDReq {
	id :number
}