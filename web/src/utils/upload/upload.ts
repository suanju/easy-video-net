import { getuploadingDir } from "@/apis/commonality";
import { GetuploadingDirReq } from "@/types/commonality/commonality";
import { FileUpload } from "@/types/idnex";
import { compressFile, isImageFile } from "./fileManipulation";
import { localUpload } from "./local";
import { ossUpload } from "./oss";
export const uploadFile = async (config: FileUpload, rawFile: File, fragment?: boolean): Promise<{ path: string }> => {
    let res
    //默认false
    if (fragment == undefined) fragment = false
    //获取保存路径和图片质量
    const response = await getuploadingDir(<GetuploadingDirReq>{
        interface: config.interface
    })
    let dir = response.data?.path as string
    let quality = response.data?.quality as number
    //压缩图片
    if (isImageFile(rawFile)) {
        try {
            const compressedFile = await compressFile(rawFile, quality as number)
            // 压缩成功后的操作
            rawFile = compressedFile as File
        } catch (err) {
            // 压缩失败后的操作
            console.log('压缩失败！', err);
        }
    }
    switch (config.uploadType) {
        case "aliyunOss":
            res = ossUpload(rawFile, config, dir, fragment)
            break;
        case "local":
            res = localUpload(rawFile, config, dir, fragment)
            break;
    }
    return res as Promise<{ path: string }>
}

