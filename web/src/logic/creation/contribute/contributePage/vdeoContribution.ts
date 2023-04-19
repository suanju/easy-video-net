import { getuploadingMethod, registerMedia } from "@/apis/commonality";
import { createVideoContribution, updateVideoContribution } from "@/apis/contribution";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useEditVideoStore } from '@/store/creation';
import { useUserStore } from "@/store/main";
import { GetUploadingMethodReq, GetUploadingMethodRes, RegisterMediaReq } from "@/types/commonality/commonality";
import { CreateVideoContributionReq, UpdateVideoContributionReq, uploadFileformation, vdeoContributionForm } from "@/types/creation/contribute/contributePage/vdeoContribution";
import { fileReader } from "@/utils/fun/fun";
import { uploadFile } from '@/utils/upload/upload';
import { validateVideoIntroduce, validateVideoTitle } from "@/utils/validate/validate";
import { ElInput, FormInstance, UploadProps, UploadRawFile, UploadRequestOptions } from 'element-plus';
import Swal from 'sweetalert2';
import { Ref, UnwrapNestedRefs, nextTick, reactive, ref } from "vue";
import { Router, useRouter } from 'vue-router';

export const useVdeoContributionProp = () => {
    const userStore = useUserStore()
    const editVideoStore = useEditVideoStore()
    const formRef = ref<FormInstance>()
    const ruleFormRef = ref<FormInstance>()
    const router = useRouter()
    const video = ref() //上传的视频信息
    const form = reactive(<vdeoContributionForm>{
        id: 0,
        isShow: false,
        title: '',
        type: false,
        labelInputVisible: false,
        labelText: "",
        label: [],
        introduce: "",
        videoDuration: 0,
    })
    const uploadFileformation = reactive(<uploadFileformation>{
        progress: 0,
        FileUrl: '',
        uploadUrl: "",
        interface: "videoContribution",
        uploadType: "",
        action: "#",
    })

    const uploadCoveration = reactive(<uploadFileformation>{
        progress: 0,
        FileUrl: '',
        uploadUrl: "",
        interface: "videoContributionCover",
        uploadType: "",
        action: "#",
    })
    const labelInputRef = ref<InstanceType<typeof ElInput>>()
    return {
        ruleFormRef,
        userStore,
        formRef,
        form,
        router,
        uploadFileformation,
        uploadCoveration,
        labelInputRef,
        video,
        editVideoStore
    }
}

//上传视频处理
export const useHandleFileMethod = (uploadFileformation: uploadFileformation, form: vdeoContributionForm, video: Ref) => {

    const handleFileSuccess: UploadProps['onSuccess'] = async (
        response,
        uploadFile
    ) => {
        uploadFileformation.FileUrl = URL.createObjectURL(uploadFile.raw!)
        //视频准备好事件
        video.value.onloadedmetadata = () => {
            //修改视频时长
            form.videoDuration = Math.round(video.value.duration)
        }
        const readerInfo = await fileReader(uploadFile.raw!)
        video.value.src = readerInfo?.result
    }

    const handleFileError: UploadProps['onError'] = (
        response,
    ) => {
        console.log("上传失败")
        Swal.fire({
            title: "上传失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",

        })
        console.log(response)

    }

    //上传前处理
    const beforeFileUpload: UploadProps['beforeUpload'] = async (rawFile: UploadRawFile) => {
        return true
    }

    //修改默认请求
    const RedefineUploadFile = async (params: UploadRequestOptions) => {
        try {
            //大于30mb分片
            let fragment = params.file.size > 30 * 1024 * 1024 ? true : false
            form.isShow = !form.isShow
            const response = await uploadFile(uploadFileformation, params.file, fragment)
            console.log(response)
            uploadFileformation.uploadUrl = response.path
            //进行媒体资源注册
            if (uploadFileformation.uploadType == "aliyunOss") {
                let media = await registerMedia(<RegisterMediaReq>{
                    type: uploadFileformation.uploadType,
                    path: response.path
                })
                uploadFileformation.media = media.data
            }
        } catch (err) {
            console.log(err)
            form.isShow = false
            Swal.fire({
                title: "上传视频失败",
                heightAuto: false,
                confirmButtonColor: globalScss.colorButtonTheme,
                icon: "error",
            })
        }
    }

    return {
        handleFileSuccess,
        beforeFileUpload,
        handleFileError,
        RedefineUploadFile
    }

}

//上传封面处理
export const useHandleCoverMethod = (uploadCoveration: uploadFileformation, form: vdeoContributionForm) => {

    const handleFileSuccess: UploadProps['onSuccess'] = (
        response,
        uploadFile
    ) => {
        uploadCoveration.FileUrl = URL.createObjectURL(uploadFile.raw!)
    }

    const handleFileError: UploadProps['onError'] = (
        response,
    ) => {
        console.log("上传失败")
        Swal.fire({
            title: "上传失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",

        })
        console.log(response)

    }

    //上传前处理
    const beforeFileUpload: UploadProps['beforeUpload'] = async (rawFile: UploadRawFile) => {
        return await new Promise<boolean>((resolve, reject) => {
            //判断大小
            if (rawFile.size / 1024 / 1024 > 2) {
                Swal.fire({
                    title: "封面大小不能大于2M",
                    heightAuto: false,
                    icon: "error",

                })
                reject(false);
            }
            //判断尺寸
            let reader = new FileReader();
            reader.readAsDataURL(rawFile);
            reader.onload = () => {
                // 让页面中的img标签的src指向读取的路径
                let img = new Image();
                img.src = reader.result as string;
                img.onload = () => {
                    console.log(img.width);
                    console.log(img.height);
                    if (img.width < 960 || img.height < 600) {
                        Swal.fire({
                            title: "请上传 960*600 尺寸以上图片",
                            heightAuto: false,
                            confirmButtonColor: globalScss.colorButtonTheme,
                            icon: "error",
                        });
                        reject(false);
                    } else {
                        resolve(true);
                    }
                };
            };
        })
    }

    //修改默认请求
    const RedefineUploadFile = async (params: UploadRequestOptions) => {
        try {
            const response = await uploadFile(uploadCoveration, params.file)
            console.log(response)
            uploadCoveration.uploadUrl = response.path
        } catch (err) {
            console.log(err)
            Swal.fire({
                title: "上传视频封面失败",
                heightAuto: false,
                confirmButtonColor: globalScss.colorButtonTheme,
                icon: "error",
            })
        }
    }

    return {
        handleFileSuccess,
        beforeFileUpload,
        handleFileError,
        RedefineUploadFile
    }

}


export const userLabelHandlMethod = (form: vdeoContributionForm, labelInputRef: Ref) => {
    const handleClose = (tag: string) => {
        form.label.splice(form.label.indexOf(tag), 1)
    }

    const showInput = () => {
        form.labelInputVisible = true
        nextTick(() => {
            labelInputRef.value!.input!.focus()
        })
    }

    const handleInputConfirm = () => {
        if (form.labelText) {
            form.label.push(form.labelText)
        }
        form.labelInputVisible = false
        form.labelText = ''
    }

    return {
        handleClose,
        showInput,
        handleInputConfirm
    }

}
export const useSaveData = async (form: vdeoContributionForm, uploadFileformation: uploadFileformation, uploadCoveration: uploadFileformation, formEl: FormInstance | undefined, router: Router, props: any) => {
    if (!formEl) return;
    await formEl.validate(async (valid, fields) => {
        if (valid) {
            try {
                if (!uploadCoveration.uploadUrl) throw "请先上传封面"
                //判断操作类型
                if (props.type == "edit") {
                    //更新视频
                    var updateRequistData = <UpdateVideoContributionReq>{
                        id: form.id,
                        cover: uploadCoveration.uploadUrl,
                        coverUploadType: uploadCoveration.uploadType,
                        title: form.title,
                        reprinted: form.type,
                        label: form.label,
                        introduce: form.introduce,
                    }
                    await updateVideoContribution(updateRequistData)
                } else {
                    //创建视频
                    if (uploadFileformation.progress != 100) throw "上传未完成"
                    var createRequistData = <CreateVideoContributionReq>{
                        id: form.id,
                        video: uploadFileformation.uploadUrl,
                        videoUploadType: uploadFileformation.uploadType,
                        cover: uploadCoveration.uploadUrl,
                        coverUploadType: uploadCoveration.uploadType,
                        title: form.title,
                        reprinted: form.type,
                        label: form.label,
                        introduce: form.introduce,
                        videoDuration: form.videoDuration,
                        media: uploadFileformation.media
                    }
                    await createVideoContribution(createRequistData)
                }
                let swalTitle = props.type == "edit" ? "更新成功" : "发布成功"
                Swal.fire({
                    title: swalTitle,
                    confirmButtonColor: globalScss.colorButtonTheme,
                    heightAuto: false,
                    icon: "success",
                    preConfirm: () => {
                        router.push({ name: "VideoManagement" })
                    }
                })
            } catch (err) {
                console.log(err)
                Swal.fire({
                    title: (err as Error).message,
                    confirmButtonColor: globalScss.colorButtonTheme,
                    heightAuto: false,
                    icon: "error",

                })
            }
        } else {
            console.log('error submit!', fields)
        }
    })
}




export const useInit = async (uploadFileformation: uploadFileformation, uploadCoveration: uploadFileformation, form: UnwrapNestedRefs<vdeoContributionForm>, props: any, editVideoStore: any) => {
    try {
        //获取当前接口的请求方法
        const updataMenhod = (await getuploadingMethod(<GetUploadingMethodReq>{
            method: uploadFileformation.interface
        })).data as GetUploadingMethodRes
        uploadFileformation.uploadType = updataMenhod.type
        const updataMenhodCover = (await getuploadingMethod(<GetUploadingMethodReq>{
            method: uploadCoveration.interface
        })).data as GetUploadingMethodRes

        //判断当前模式
        if (props.type == "edit") {
            //编辑模式
            form.isShow = true
            uploadFileformation.progress = 100
            form.id = editVideoStore.editVideoData.videoID
            form.title = editVideoStore.editVideoData.title
            form.label = editVideoStore.editVideoData.label
            form.type = editVideoStore.editVideoData.reprinted
            form.introduce = editVideoStore.editVideoData.introduce
            uploadCoveration.FileUrl = editVideoStore.editVideoData.cover
            uploadCoveration.uploadUrl = editVideoStore.editVideoData.cover_url
            uploadCoveration.uploadType = editVideoStore.editVideoData.coverUploadType
        }



        uploadCoveration.uploadType = updataMenhod.type

    } catch (err) {
        console.log(err)
        Swal.fire({
            title: "获取上传方法失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}

//表单验证
export const useRules = () => {
    const videoContributionRules = reactive({
        title: [{ validator: validateVideoTitle, trigger: 'change' }],
        introduce: [{ validator: validateVideoIntroduce, trigger: 'change' }]
    });
    return {
        videoContributionRules,
    };
}; 
