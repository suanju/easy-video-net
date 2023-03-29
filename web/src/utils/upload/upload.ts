import { FileUpload } from "@/types/idnex";
import { localUpload } from "./local";
import { ossUpload } from "./oss";
export const uploadFile = (config: FileUpload, rawFile: File, fragment?: boolean): Promise<{ path: string }> => {
    let res
    //默认false
    if (fragment == undefined) fragment = false
    switch (config.uploadType) {
        case "aliyunOss":
            res = ossUpload(rawFile, config, fragment)
            break;
        case "local":
            res = localUpload(rawFile, config, fragment)
            break;
    }
    return res as Promise<{ path: string }>
}

