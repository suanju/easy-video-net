import { FileUpload } from "@/types/idnex"
export interface CreateCollectRmation extends FileUpload {
    id : number
    title: string, //标题
    content: string //内容
}

export interface SaveCreateFavoritesDataReq {
    id : number 
    type: string
    title: string
    content: string
    cover: string
}

//用于单标题请求
export  interface createFavoritesReq {
    title : string
}