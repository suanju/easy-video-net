import { PageInfo } from "@/types/idnex"

export interface GetDiscussBarrageListReq {
    page_info: PageInfo
}

export interface GetDiscussBarrageListItem {
    id: number
    username: string
    photo: string
    comment: string
    cover: string
    title: string
    created_at: string
}
export type GetDiscussBarrageListRes = Array<GetDiscussBarrageListItem>