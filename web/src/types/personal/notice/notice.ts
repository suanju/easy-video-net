import { PageInfo } from "@/types/idnex"

export interface GetNoticeListReq {
    type: string
    page_info: PageInfo
}

export interface GetNoticeListItem {
    id: number
    username: string
    type: string
    to_id: number
    photo: string
    comment: string
    cover: string
    title: string
    created_at: string

    type_rompt: string
}
export type GetNoticeListRes = Array<GetNoticeListItem>