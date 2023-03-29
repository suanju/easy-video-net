import { editArticle, editVideo } from "@/types/store/creation"
import { defineStore } from "pinia"
import { reactive } from "vue"

export const useEditVideoStore = defineStore("editVideo", () => {

    const editVideoData = reactive(<editVideo>{
        videoID: 0,
        cover: "",
        cover_url : "",
        coverUploadType: "",
        reprinted : false,
        title: "",
        label: [],
        introduce: "",
        videoDuration: 0
    })

    const setEditVideoData = (info: editVideo) => {
        editVideoData.videoID = info.videoID
        editVideoData.cover = info.cover
        editVideoData.coverUploadType = info.coverUploadType
        editVideoData.cover_url = info.cover_url
        editVideoData.title = info.title
        editVideoData.reprinted = info.reprinted
        editVideoData.label = info.label
        editVideoData.introduce = info.introduce
        editVideoData.videoDuration = info.videoDuration
    }
    return {
        editVideoData,
        setEditVideoData
    }
}, {
    persist: true,
})



export const useEditArticleStore = defineStore("editArticle", () => {
    const editArticleData = reactive(<editArticle>{
        articleID : 0,
        cover: "",
        title : "",
        cover_url : "",
        cover_upload_type: "",
        content: "",
        is_comments: false,
        label: [],
        classification_id: 0,
    })

    const setEditArticleData = (info: editArticle) => {
        editArticleData.articleID = info.articleID
        editArticleData.title = info.title
        editArticleData.cover = info.cover
        editArticleData.cover_upload_type = info.cover_upload_type
        editArticleData.cover_url = info.cover_url
        editArticleData.content = info.content
        editArticleData.is_comments = info.is_comments
        editArticleData.label = info.label
        editArticleData.classification_id = info.classification_id
        
    }
    return {
        editArticleData,
        setEditArticleData
    }
}, {
    persist: true,
})