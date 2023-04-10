import { gteossSTS } from "@/apis/commonality"
import { useGlobalStore } from "@/store/main"
import { FileUpload, OssSTSInfo } from "@/types/idnex"
import OSS from 'ali-oss'
import { fileHash, fileSuffix } from "./fileManipulation"


//初始化sts
export const initOssSTS = async (_interface: string): Promise<OssSTSInfo> => {
    return new Promise((resolve, reject) => {
        // 从本地localstore从获取配置
        const globalStore = useGlobalStore()
        const conf = globalStore.ossData
        if (conf) {
            //  配置存在并且距离过期时间还大一分钟秒则返回此配置
            const now = new Date().getTime() / 1000
            console.log(conf)
            if (conf.expirationTime - 600 > now) {
                resolve(conf)
                return
            }
        }
        // 请求接口返回配置数据
        gteossSTS()
            .then((res) => {
                if (res.code == 200) {
                    if (!res.data) return false
                    let info = res.data
                    // 配置数据写入本地store
                    globalStore.setOssInfo(<OssSTSInfo>{
                        region: info.region,
                        accessKeyId: info.access_key_id,
                        accessKeySecret: info.access_key_secret,
                        stsToken: info.sts_token,
                        bucket: info.bucket,
                        expirationTime: info.expiration_time
                    })
                    resolve(globalStore.ossData);
                } else {
                    reject(res)
                }
            })
            .catch((err) => {
                console.log(err);
                reject(err)
            })
    })
}


/**
 * 往Oss上传文件
 * @param file File对象
 * @returns {Promise<{name:string,host:string}>}
 */
export const ossUpload = (file: File, uploadConfig: FileUpload, dir: string, fragment: boolean): Promise<{ path: string }> => {
    return new Promise((resolve, reject) => {
        initOssSTS(uploadConfig.interface)
            .then(async (ossSts) => {
                //得到名称及其初始化
                const name = await fileHash(file) + fileSuffix(file.name)
                const key = `${dir}${name}`
                // 初始化阿里云oss客户端
                const client = new OSS({
                    region: ossSts.region,
                    accessKeyId: ossSts.accessKeyId,
                    accessKeySecret: ossSts.accessKeySecret,
                    stsToken: ossSts.stsToken,
                    bucket: ossSts.bucket,
                });

                if (!fragment) {
                    console.log("普通上传")
                    //为了能够显示进度条的取舍也进行了分片上传
                    var checkpoint = getCheckpoint(name);
                    const options = {
                        checkpoint: checkpoint,
                        progress: (p: any, cpt: any) => {
                            console.log("上传进度", p)
                            uploadConfig.progress = Math.round(p * 100)
                            saveCheckpoint(name, cpt);
                        },
                        mime: "text/plain",
                        // 设置并发上传的分片数量。
                        parallel: 4,
                        // 设置分片大小。默认值为1 MB，最小值为100 KB。
                        partSize: 200 * 1024,
                    };

                    try {
                        const res = await client.multipartUpload(`${dir}${name}`, file, {
                            ...options,
                        });
                        console.log(res);
                        deleteCheckpoint(name);
                        resolve({ path: key })
                    } catch (err) {
                        console.log(err);
                        deleteCheckpoint(name);
                        reject({ msg: '上传失败' })
                    }
                } else {
                    console.log("分片上传")
                    //分片上传加断点续传
                    var checkpoint = getCheckpoint(name);
                    const options = {
                        checkpoint: checkpoint,
                        // 获取分片上传进度、断点和返回值。
                        progress: (p: any, cpt: any) => {
                            saveCheckpoint(name, cpt);
                            console.log(cpt)
                            uploadConfig.progress = Math.round(p * 100)
                        },
                        // 设置并发上传的分片数量。
                        parallel: 4,
                        // 设置分片大小。默认值为1 MB，最小值为100 KB。
                        partSize: 1 * 1024 * 1024,
                        mime: "text/plain",
                    };
                    try {
                        const res = await client.multipartUpload(`${dir}${name}`, file, {
                            ...options,
                        });
                        deleteCheckpoint(name);
                        resolve({ path: key })

                        console.log(res)
                    } catch (err) {
                        deleteCheckpoint(name);
                        console.log(err);
                        reject({ msg: '上传失败' })
                    }
                }
            })
            .catch((err) => {
                console.log(err);
                reject({ msg: '上传失败' })
            })
    })
}

// 保存上传断点
function saveCheckpoint(key: string, checkpoint: any) {
    localStorage.setItem(key, JSON.stringify(checkpoint));
}

// 获取上传断点
function getCheckpoint(key: string,) {
    var checkpoint = localStorage.getItem(key);
    if (checkpoint) {
        return JSON.parse(checkpoint);
    } else {
        return null;
    }
}

// 删除上传断点
function deleteCheckpoint(key: string) {
    localStorage.removeItem(key);
}