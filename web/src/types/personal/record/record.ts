import { PageInfo } from "@/types/idnex";

export interface GetRecordListReq {
    page_info: PageInfo
}


export  interface GetRecordListItem {
	id :number
    to_id : number
	title :string
	cover :string
	username :string
	photo :string
	type :string
    updated_at :string

	is_delete : boolean //用于伪删除
}

export type GetRecordListRes = Array<GetRecordListItem>

export interface DeleteRecordByIDReq{
	id :number
}