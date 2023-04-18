//form 表单结构
export interface ArticleContribution {
    id: number,
    isShow: boolean
    title: string
    content: string
    comments: boolean,
    labelInputVisible: boolean,
    labelText: string,
    label: Array<string>,
    classification_id: number
}

//api createArticleContribution 需要结构
export interface AreateArticleContributionReq {
    cover: string,
    coverUploadType: string,
    articleContributionUploadType: string,
    title: string,
    timing: boolean,
    timingTime?: string,
    label: Array<string>,
    content: string,
    comments: boolean
    classification_id: number
}

export interface UpdateArticleContributionReq {
    id: number
    cover: string,
    coverUploadType: string,
    articleContributionUploadType: string,
    title: string,
    label: Array<string>,
    content: string,
    comments: boolean
    classification_id: number
}

//api GetArticleContributionListByUser 需要结构
export interface GetArticleContributionListByUserReq {
    userID: Number
}


export interface GetArticleClassificationListResItem {
    id: number
    aid: number
    label: string
    children: Array<GetArticleClassificationListResItem>
}
export type GetArticleClassificationListRes = Array<GetArticleClassificationListResItem>


export interface getArticleTotalInfoRes {
    classification: GetArticleClassificationListRes
    article_num: number
    classification_num: number
}


