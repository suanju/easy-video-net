import { GetFullPathOfImageRrq, GetuploadingDirReq, GetuploadingDirRes, GetUploadingMethodReq, GetUploadingMethodRes, GteossStsres, SearchReq, SearchRes, UploadCheckReq, UploadCheckRes, UploadMergeReq, UploadMergeRes } from "@/types/commonality/commonality";
import { FileSliceUpload, FileUpload } from "@/types/idnex";
import httpRequest from "@/utils/requst";

//获取oss sts 信息
export const gteossSTS = () => {
    return httpRequest.post<GteossStsres>('/commonality/ossSTS');
}

//获取上传方法
export function getuploadingMethod(params: GetUploadingMethodReq) {
    return httpRequest.post<GetUploadingMethodRes>('/commonality/uploadingMethod', params);
}

//获取保存地址
export function getuploadingDir(params: GetuploadingDirReq) {
    return httpRequest.post<GetuploadingDirRes>('/commonality/uploadingDir', params);
}

//获取图片完整路径
export function getFullPathOfImage(params: GetFullPathOfImageRrq) {
    return httpRequest.post<string>('/commonality/getFullPathOfImage', params);
}

//搜索功能
export function search(params: SearchReq) {
    return httpRequest.post<SearchRes>('/commonality/search', params);
}

//上传文件
export const uploadFile = (params: any, uploadConfig: FileUpload) => {
    return httpRequest.upload('/commonality/upload', params, uploadConfig);
}

//分片上传文件
export const UploadSliceFile = (params: any, uploadConfig: FileSliceUpload) => {
    return httpRequest.uploadSlice('/commonality/UploadSlice', params, uploadConfig);
}

//上传文件验证(验证操作)
export const uploadCheck = (params: UploadCheckReq) => {
    return httpRequest.post<UploadCheckRes>('/commonality/uploadCheck', params);
}

//上传文件验证(合并操作)
export const uploadMerge = (params: UploadMergeReq) => {
    return httpRequest.post<UploadMergeRes>('/commonality/uploadMerge', params);
}