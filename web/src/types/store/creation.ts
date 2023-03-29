export interface editVideo {
    videoID : number
    cover: string,
    cover_url : string
    coverUploadType: string,
    reprinted : boolean,
    title: string,
    label:  Array<string>,
    introduce: string,
    videoDuration : number
}

export interface editArticle {
    articleID :number
    title : string
    cover: string,
    cover_url : string
    cover_upload_type: string,
    content: string
    is_comments: boolean,
    label: Array<string>,
    classification_id: number
}