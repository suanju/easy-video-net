import { PageInfo } from "@/types/idnex"

export interface GetArticleManagementListReq {
    page_info: PageInfo
}


export interface GetArticleManagementListItem {
    id: number
    classification_id: number
    title: string
    cover: string
    cover_url : string
    cover_upload_type : string
    label: Array<string>
    content: string
    is_comments: boolean
    heat: number

    is_delete: boolean //用于伪删除
}

export type GetArticleManagementListRes = Array<GetArticleManagementListItem>

export interface DeleteArticleByIDReq {
    id: number
}