import { getuploadingMethod } from "@/apis/commonality";
import { getLiveDataRequist, saveLiveDataRequist } from "@/apis/personal";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useUserStore } from "@/store/main";
import { GetUploadingMethodReq, GetUploadingMethodRes } from "@/types/commonality/commonality";
import { GetLiveDataRes, LiveInformation, SaveLiveDataReq } from "@/types/personal/live/setUp";
import { getLocation } from "@/utils/conversion/stringConversion";
import { uploadFile } from '@/utils/upload/upload';
import { validateLiveTitle } from "@/utils/validate/validate";
import { FormInstance, UploadProps, UploadRequestOptions } from 'element-plus';
import Swal from 'sweetalert2';
import { reactive, ref } from "vue";
import useClipboard from "vue-clipboard3";

export const useLiveInfoProp = () => {
    const userStore = useUserStore()
    const saveDateFormRef = ref<FormInstance>()
    const liveInformationForm = reactive<LiveInformation>({
        FileUrl: '',
        uploadUrl: "",
        interface: "liveCover",
        title: "",
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
        liveInformationForm,
        saveDateFormRef,
        rawData
    }
}


export const useHandleFileMethod = (liveInformationForm: LiveInformation) => {

    const handleFileSuccess: UploadProps['onSuccess'] = (
        response,
        uploadFile
    ) => {
        liveInformationForm.FileUrl = URL.createObjectURL(uploadFile.raw!)
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
            const response = await uploadFile(liveInformationForm, params.file)
            liveInformationForm.uploadUrl = response.path
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

export const useSaveData = async (liveInformationForm: LiveInformation, formEl: FormInstance | undefined, rawData: GetLiveDataRes) => {
    if (!formEl) return;

    await formEl.validate(async (valid, fields) => {
        if (valid) {
            try {
                if (liveInformationForm.uploadUrl == rawData.img && liveInformationForm.title == rawData.title) throw "未修改信息";
                if (!liveInformationForm.uploadUrl) throw "请先上传图片"
                const requistData = <SaveLiveDataReq>{
                    type: liveInformationForm.uploadType,
                    imgUrl: liveInformationForm.uploadUrl,
                    title: liveInformationForm.title,
                }
                const data = await saveLiveDataRequist(requistData)
                console.log(data)
                Swal.fire({
                    title: "修改成功",
                    confirmButtonColor: globalScss.colorButtonTheme,
                    heightAuto: false,
                    icon: "success",

                })
                console.log("上传成功")
            } catch (err) {
                console.log(err)
                Swal.fire({
                    title: err as string,
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

export const useInit = async (liveInformationForm: LiveInformation, rawData: GetLiveDataRes) => {
    try {
        //获取用户信息
        const data = (await getLiveDataRequist()).data as GetLiveDataRes;
        liveInformationForm.FileUrl = data.img
        const imgPathInfo = getLocation(data.img)

        //如何后端返回全路径取域名后路径
        if (imgPathInfo?.pathname) {
            let pathname = imgPathInfo?.pathname.slice(1)
            liveInformationForm.uploadUrl = pathname
            rawData.img = pathname
        }
        //保存原始数据
        rawData.title = data.title
        liveInformationForm.title = data.title
        //获取当前接口的请求方法
        const updataMenhod = (await getuploadingMethod(<GetUploadingMethodReq>{
            method: liveInformationForm.interface
        })).data as GetUploadingMethodRes
        liveInformationForm.uploadType = updataMenhod.type
        console.log(updataMenhod)

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



export const useCopy = async (text: string) => {
    try {
        const { toClipboard } = useClipboard();
        await toClipboard(text); //实现复制
        Swal.fire({
            title: "复制成功",
            confirmButtonColor: globalScss.colorButtonTheme,
            heightAuto: false,
            icon: "success",

        })
    } catch (e) {
        console.error(e);
    }
};

//表单验证
export const useRules = () => {
    const liveInformationRules = reactive({
        title: [{ validator: validateLiveTitle, trigger: 'change' }],
    });
    return {
        liveInformationRules,
    };
};