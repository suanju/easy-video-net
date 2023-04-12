import { getuploadingMethod } from "@/apis/commonality";
import { createFavorites } from "@/apis/personal";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useUserStore } from "@/store/main";
import { GetUploadingMethodReq, GetUploadingMethodRes } from "@/types/commonality/commonality";
import { CreateCollectRmation, SaveCreateFavoritesDataReq } from "@/types/personal/collect/createFavorites";
import { GetFavoritesListItem } from "@/types/personal/collect/favorites";
import { GetLiveDataRes } from "@/types/personal/live/setUp";
import { uploadFile } from '@/utils/upload/upload';
import { validateCollectTitle } from "@/utils/validate/validate";
import { FormInstance, UploadProps, UploadRequestOptions } from 'element-plus';
import Swal from 'sweetalert2';
import { reactive, ref } from "vue";
import { Router, useRouter } from 'vue-router';

export const useCreateFavoritesProp = () => {
    const userStore = useUserStore()
    const router = useRouter()
    const saveDateFormRef = ref<FormInstance>()
    const createFavoriteRmationForm = reactive<CreateCollectRmation>({
        id: 0,
        FileUrl: '',
        uploadUrl: "",
        interface: "createFavoritesCover",
        title: "",
        content: "",
        uploadType: "",
        action: "#",
        progress: 0
    });

    //定义请求结果原始数据
    const rawData = reactive<GetLiveDataRes>({
        title: "",
        img: ""
    })

    return {
        userStore,
        createFavoriteRmationForm,
        saveDateFormRef,
        rawData,
        router
    }
}


//文件上传处理
export const useHandleFileMethod = (createFavoriteRmationForm: CreateCollectRmation) => {
    const handleFileSuccess: UploadProps['onSuccess'] = (
        response,
        uploadFile
    ) => {
        createFavoriteRmationForm.FileUrl = URL.createObjectURL(uploadFile.raw!)
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


    const beforeFileUpload: UploadProps['beforeUpload'] = async (rawFile) => {
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
    const RedefineUploadFile = async (params: UploadRequestOptions) => {
        try {
            const response = await uploadFile(createFavoriteRmationForm, params.file)
            console.log(response)
            createFavoriteRmationForm.uploadUrl = response?.path
            console.log(response)
        } catch (err) {
            console.log(err)
            Swal.fire({
                title: "获取上传节点失败",
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

export const useSaveData = async (createFavoriteRmationForm: CreateCollectRmation, formEl: FormInstance | undefined, rawData: GetLiveDataRes, router: Router, emits: any) => {
    if (!formEl) return;

    await formEl.validate(async (valid, fields) => {
        if (valid) {
            try {
                if (createFavoriteRmationForm.uploadUrl == rawData.img && createFavoriteRmationForm.title == rawData.title) throw "未修改信息";
                if (!createFavoriteRmationForm.uploadUrl) throw "请先上传图片"
                const requistData = <SaveCreateFavoritesDataReq>{
                    id: createFavoriteRmationForm.id,
                    type: createFavoriteRmationForm.uploadType,
                    cover: createFavoriteRmationForm.uploadUrl,
                    title: createFavoriteRmationForm.title,
                    content: createFavoriteRmationForm.content
                }
                await createFavorites(requistData)
                emits("shutDown");
                if (createFavoriteRmationForm.id == 0) {
                    //清空内容
                    createFavoriteRmationForm.FileUrl = ""
                    createFavoriteRmationForm.title = ""
                    createFavoriteRmationForm.content = ""
                    //创建模式
                    Swal.fire({
                        title: "创建成功",
                        confirmButtonColor: globalScss.colorButtonTheme,
                        heightAuto: false,
                        icon: "success",
                        preConfirm: () => {
                            router.push({ name: "Favorites", query: { type: 'createTime' + Date.now() } })
                        }
                    })
                } else {
                    //更新模式
                    Swal.fire({
                        title: "更新成功",
                        confirmButtonColor: globalScss.colorButtonTheme,
                        heightAuto: false,
                        icon: "success",
                    })
                }

            } catch (err: any) {
                console.log(err)
                emits("shutDown");
                Swal.fire({
                    title: (err as Error).message,
                    confirmButtonColor: globalScss.colorButtonTheme,
                    heightAuto: false,
                    icon: "warning",
                })
            }
        } else {
            console.log('error submit!', fields)
        }
    })
}


export const useInit = async (createFavoriteRmationForm: CreateCollectRmation, rawData: GetLiveDataRes, type: boolean, item: GetFavoritesListItem | undefined) => {
    try {
        //获取当前接口的请求方法
        const updataMenhod = (await getuploadingMethod(<GetUploadingMethodReq>{
            method: createFavoriteRmationForm.interface
        })).data as GetUploadingMethodRes
        createFavoriteRmationForm.uploadType = updataMenhod.type
        if (!type) {
            if (item == undefined) return false
            if (createFavoriteRmationForm.FileUrl) {
                //已上传图片才更改上传类型
                createFavoriteRmationForm.uploadType = item.type
            }
            createFavoriteRmationForm.title = item.title
            createFavoriteRmationForm.content = item.content
            createFavoriteRmationForm.uploadUrl = item.src
            createFavoriteRmationForm.FileUrl = item.cover
            createFavoriteRmationForm.id = item.id
        }

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
    const liveInformationRules = reactive({
        title: [{ validator: validateCollectTitle, trigger: 'change' }],
    });
    return {
        liveInformationRules,
    };
};