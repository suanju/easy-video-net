import { uploadCheck, uploadFile, uploadMerge, UploadSliceFile } from "@/apis/commonality";
import { UploadCheckReq, UploadMergeReq, UploadSliceList } from "@/types/commonality/commonality";
import { FileSliceUpload, FileUpload } from "@/types/idnex";
import { ref, watch } from "vue";
import { fileHash, fileSuffix } from "./fileManipulation";
import { getSliceFile } from "./getSliceFile";
/**
 * 往本地上传文件
 * @param file File对象
 * @returns {Promise<{name:string,host:string}>}
 */
export const localUpload = async (file: File, uploadConfig: FileUpload, fragment?: boolean): Promise<any> => {
    if (!fragment) {
        //直接上传
        return new Promise(async (resolve, reject) => {
            // 计算文件Hash 避免多余的文件上传，这样做的目的是尽量少占用的空间
            const name = await fileHash(file) + fileSuffix(file.name)
            const formData = new FormData()
            const key = `${name}`
            formData.append('interface', uploadConfig.interface)
            formData.append('name', name)
            formData.append('file', file)
            try {
                const response = await uploadFile(formData, uploadConfig)
                resolve({ path: response.data as string })
                console.log(response)
            } catch (err) {
                console.log(err)
                reject({ msg: '上传失败' })
            }
        })
    } else {
        const uploadCheckFun = async () => {
            // 计算文件Hash 避免多余的文件上传，这样做的目的是尽量少占用的空间
            const name = await fileHash(file) + fileSuffix(file.name)
            //总切片
            let sliceArr = await getSliceFile(file, 1, name)
            let sliceList = ref(<UploadSliceList>[])
            sliceArr.filter((item) => {
                sliceList.value.push({
                    index: item.index,
                    hash: item.hash
                })
            })
            const uploadCheckResponse = await uploadCheck(<UploadCheckReq>{
                file_md5: name,
                interface: uploadConfig.interface,
                slice_list: sliceArr
            })
            if (uploadCheckResponse.data?.is_upload) {
                uploadConfig.progress = 100
                return new Promise((resolve) => { resolve({ path: uploadCheckResponse.data?.path }) })
            }

            let notUploadIndex: number[] = []
            uploadCheckResponse.data?.list.filter((item) => {
                notUploadIndex.push(item.index)
            })
            //取未上传的分片
            const notUploadSlice = sliceArr.filter((item) => {
                return (notUploadIndex.includes(item.index))
            })
            //将已经上传的进度设置成100%
            sliceArr = sliceArr.filter((item) => {
                if (!notUploadIndex.includes(item.index)) {
                    item.progress = 100
                }
                return item
            })
            console.log("所以需要上传的切片", sliceArr)
            console.log("未上传的切片", notUploadSlice)

            for (let i = 0; i < notUploadSlice.length; i++) {
                const formData = new FormData()
                formData.append('interface', uploadConfig.interface)
                formData.append('name', notUploadSlice[i].hash)
                formData.append('file', notUploadSlice[i].file)
                try {
                    const slice = ref(<FileSliceUpload>{
                        index: i,
                        progress: 0, //上传进度
                        size: notUploadSlice[i].size
                    })
                    let w = watch(() => { slice.value.progress }, () => {
                        //计算
                        sliceArr.filter((item, index, arr) => {
                            if (item.index === notUploadSlice[i].index) {
                                sliceArr[index].progress = slice.value.progress
                                updataProgress(sliceArr, uploadConfig)
                            }
                            return item;
                        })
                        if (slice.value.progress === 100) {
                            w()
                            return
                        }
                    }, { deep: true })
                    await UploadSliceFile(formData, slice.value)
                    notUploadSlice.reverse()
                } catch (err) {
                    console.log(err)
                    //如果失败等待3秒进行重试
                    setTimeout(() => { uploadCheckFun() }, 3000)
                    break
                }
            }
            try {
                //分片全部上传成功进行合并
                const uploadMergeResponse = await uploadMerge(<UploadMergeReq>{
                    file_name: name,
                    interface: uploadConfig.interface,
                    slice_list: sliceArr
                })
                uploadConfig.progress = 100
                return new Promise((resolve, _) => {
                    resolve({ path: uploadMergeResponse.data })
                })
            } catch (err) {
                return new Promise((_, reject) => {
                    reject({ msg: "上传失败" })
                })
            }

        }

        const updataProgress = (sliceArr: Array<any>, uploadConfig: FileUpload) => {
            const totalSize = file.size
            let loadSize = 0

            sliceArr.filter((item) => {
                //计算每片上传大小
                loadSize += (item.sliceSizeInByte * item.progress) / 100
            })
            let progress = Math.round(loadSize / totalSize * 100)
            uploadConfig.progress = progress
        }
        return uploadCheckFun()
    }

} 