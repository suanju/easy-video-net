import { UserInfoRes } from '@/types/personal/userInfo/userInfo';
import { reactive } from "vue";
import { getUserInfoRequist, setUserInfoRequist } from "@/apis/personal"
import { validateRepeatName } from '@/utils/validate/validate';
import globalScss from "@/assets/styles/global/export.module.scss"
import Swal from 'sweetalert2';

export const useUserInfoProp = () => {
    const form = reactive<UserInfoRes>({
        username: "",
        gender: 0,
        birth_date: "",
        is_visible: true,
        signature: "",
    });

    return {
        form
    }
}

export const useUserInfoMethod = (form: UserInfoRes) => {
    const onSubmit = () => {
        setUserInfo()
    };

    const getUserInfo = async () => {
        try {
            let data = await getUserInfoRequist()
            data.data = <UserInfoRes>data.data
            form.username = data.data.username
            form.gender = data.data.gender
            form.birth_date = data.data.birth_date
            form.is_visible = data.data.is_visible
            form.signature = data.data.signature
        } catch (err) {
            console.log(err)
        }
    }

    const setUserInfo = async () => {
        try {
            let data = await setUserInfoRequist(form)
            Swal.fire({
                title: "修改成功",
                heightAuto: false,
                icon: "success",
                confirmButtonColor: globalScss.colorButtonTheme,
            })
            console.log(data)
        } catch (err) {
            console.log(err)
        }
    }

    return {
        getUserInfo,
        setUserInfo,
        onSubmit
    }
}


//表单验证
export const useRules = () => {
    const userInfoRules = reactive({
        username: [{ validator: validateRepeatName, trigger: 'change' }],
    });
    return {
        userInfoRules,
    };
}; 
