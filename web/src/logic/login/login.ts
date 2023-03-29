import { forgetRequist, loginRequist, regist, sendEmailVerificationCode, sendEmailVerificationCodeByForget } from "@/apis/login";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useGlobalStore, useUserStore } from "@/store/main";
import { Result } from "@/types/idnex";
import { forgetReq, loginReq, registReq, sendEmailInfo, sendEmailReq, sendEmailType, userInfoRes } from "@/types/login/login";
import { validateEmail, validatePassword, validateUsername, validateVarCode } from "@/utils/validate/validate";
import { ElMessage, FormInstance } from "element-plus";
import Swal from "sweetalert2";
import { reactive, ref, Ref } from "vue";
import { Router, useRouter } from 'vue-router';
import { useInitChatSocket } from "../global/chat";
import { useInitNoticeSocket } from "../global/notice";
//提供数据

const loading = useGlobalStore().globalData.loading

export const useLoginProp = () => {
    //公共部分
    const currentModel = ref(false) //当前模块 true 注册 false 登入
    const userStore = useUserStore()
    const router = useRouter();
    const isShowSendEmail = ref(false)
    const loginFormRef = ref<FormInstance>()
    const isForget = ref(false)


    const loginForm = reactive<loginReq>({
        username: "",
        password: ""
    })

    const regFormRef = ref<FormInstance>()
    const regForm = reactive<registReq>({
        email: "",
        username: "",
        password: "",
        verificationCode: "",
    })

    const forgetFormRef = ref<FormInstance>()
    const forgetForm = reactive<forgetReq>({
        email: "",
        password: "",
        verificationCode: "",
    })

    let sendEmailInfo = reactive<sendEmailInfo>({
        btnText: "send",
        isPleaseClick: true
    })

    const sendEmail = async (email: string, emailType: sendEmailType) => {
        if (!emailType) return
        if (sendEmailInfo.isPleaseClick == false) return
        try {
            const emailReq = <sendEmailReq>{
                email: email
            }
            loading.loading = true
            if (emailType == "regist") {
                const result = await sendEmailVerificationCode(emailReq)
            } else {
                const result = await sendEmailVerificationCodeByForget(emailReq)
            }

            sendEmailInfo.isPleaseClick = false
            loading.loading = false
            ElMessage({
                message: "验证码已发送！",
                type: 'success',
            })
            sendEmailInfo.btnText = "60"
            const interval = setInterval(() => {
                if (Number(sendEmailInfo.btnText) <= 0) {
                    sendEmailInfo.btnText = "send"
                    sendEmailInfo.isPleaseClick = true
                    clearInterval(interval)
                } else {
                    sendEmailInfo.btnText = String(Number(sendEmailInfo.btnText) - 1)
                }
            }, 1000)


        } catch (err) {
            loading.loading = false
            console.log(err)

        }
    }

    return {
        currentModel,
        userStore,
        router,
        loginForm,
        forgetForm,
        loginFormRef,
        regForm,
        regFormRef,
        isShowSendEmail,
        sendEmailInfo,
        sendEmail,
        isForget,
        forgetFormRef
    }
}
//登入
export const useLoginMethod = (store: any, router: Router, loginForm: loginReq) => {

    const login = async (formEl: FormInstance | undefined) => {
        console.log(formEl)
        if (!formEl) return;
        await formEl.validate(async (valid, fields) => {
            if (valid) {
                try {
                    loading.loading = true
                    const result = await loginRequist(loginForm)
                    store.setUserInfo(<userInfoRes>result.data)
                    //初始化socket
                    useInitChatSocket()
                    useInitNoticeSocket()
                    ElMessage({
                        message: "登入成功",
                        type: 'success',
                    })
                    loading.loading = false
                    router.push("/")

                }
                catch (err: any) {
                    err as unknown as Result
                    loading.loading = false
                    Swal.fire({
                        title: err.message,
                        heightAuto: false,
                        confirmButtonColor: globalScss.colorButtonTheme,
                        icon: "error",
                    })
                    console.log(err)
                }
            } else {
                console.log('error submit!', fields)
            }
        })

    }
    return {
        login
    }
}


//注册
export const useRegisterMethod = (store: any, router: Router, registForm: registReq) => {

    const register = async (formEl: FormInstance | undefined) => {
        console.log(formEl)
        if (!formEl) return;
        await formEl.validate(async (valid, fields) => {
            if (valid) {
                try {
                    loading.loading = true
                    const result = await regist(registForm)
                    store.setUserInfo(<userInfoRes>result.data)
                    ElMessage({
                        message: "注册成功",
                        type: 'success',
                    })
                    loading.loading = false
                    router.push("/")
                }
                catch (err: any) {
                    err as unknown as Result
                    loading.loading = false
                    Swal.fire({
                        title: err.message,
                        heightAuto: false,
                        confirmButtonColor: globalScss.colorButtonTheme,
                        icon: "error",
                    })
                    console.log(err)
                }
            } else {
                console.log('error submit!', fields)
            }
        })
    }
    return {
        register
    }
}


//找回密码
export const useForgetrMethod = (forgetForm: forgetReq, currentModel: Ref<boolean>) => {

    const forfet = async (formEl: FormInstance | undefined) => {
        if (!formEl) return;
        await formEl.validate(async (valid, fields) => {
            if (valid) {
                try {
                    loading.loading = true
                    const result = await forgetRequist(forgetForm)
                    ElMessage({
                        message: "修改成功",
                        type: 'success',
                    })
                    forgetForm.email = ""
                    forgetForm.password = ""
                    forgetForm.verificationCode = ""
                    //清除表单校验 需要延迟
                    setTimeout(() => {
                        formEl.clearValidate()
                    }, 500)
                    currentModel.value = false
                    loading.loading = false
                    return

                }
                catch (err: any) {
                    err as unknown as Result
                    loading.loading = false
                    Swal.fire({
                        title: err.message,
                        heightAuto: false,
                        confirmButtonColor: globalScss.colorButtonTheme,
                        icon: "error",
                    })
                    console.log(err)
                }
            } else {
                formEl.clearValidate()
                console.log('error submit!', fields)
            }
        })
    }
    return {
        forfet
    }
}

//发送验证码
export const useSendEmail = async (email: string) => {
    try {
        const emailReq = <sendEmailReq>{
            email: email
        }
        const result = await sendEmailVerificationCode(emailReq)
        return result
    } catch (err) {
        console.log(err)
        return err

    }
}

//表单验证
export const useRules = () => {
    const loginRules = reactive({
        username: [{ validator: validateUsername, trigger: 'change' }],
        password: [{ validator: validatePassword, trigger: 'change' }]
    });
    // 注册表单校验规则
    const registerRules = reactive({
        username: [{ validator: validateUsername, trigger: 'change' }],
        email: [{ validator: validateEmail, trigger: 'change' }],
        password: [{ validator: validatePassword, trigger: 'change' }],
        verificationCode: [{ validator: validateVarCode, trigger: 'change' }]
    });

    const forgetRules = reactive({
        email: [{ validator: validateEmail, trigger: 'change' }],
        password: [{ validator: validatePassword, trigger: 'change' }],
        verificationCode: [{ validator: validateVarCode, trigger: 'change' }]
    });
    return {
        loginRules,
        registerRules,
        forgetRules
    };
};
