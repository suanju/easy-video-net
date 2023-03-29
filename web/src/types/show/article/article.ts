import { PageInfo } from "@/types/idnex"

export interface GetArticleContributionByIDReq {
    articleID: number

}
export interface Comments {
    id: number
    comment_id: number
    created_at: string
    context: string
    comment_user_id : number
    comment_user_name : string
    uid: number
    username: string
    photo: string
    lowerComments: any
}


type commentsList = Array<Comments>


export interface GetArticleContributionByIDRes {
    id: number
    uid: number
    title: string
    cover: string
    label: Array<string>
    content: string
    is_comments: boolean
    heat: number
    likes_number: number
    comments: commentsList
    comments_number: number
    created_at: string
    is_stay: boolean
}

export interface ArticlePostCommentReq {
    article_id: number
    content: string
    content_id: number
}

//页面中评论相关
export interface CommentsInfo{
    comments : string,
    commentsID : number
}

export interface GetArticleCommentReq {
    pageInfo : PageInfo
    articleID: number

}

export interface GetArticleCommentRes {
    id: number
    comments: commentsList
    comments_number: number
}
