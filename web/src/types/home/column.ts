import { PageInfo } from "../idnex";

export interface GetArticleContributionListReq {
    page_info: PageInfo
}


export interface GetArticleContributionListItem {
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

export type GetArticleContributionListRes = Array<GetArticleContributionListItem>
