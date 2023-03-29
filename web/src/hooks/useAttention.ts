import { attention } from "@/apis/personal";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useUserStore } from "@/store/main";
import { AttentionReq } from "@/types/personal/userInfo/userInfo";
import Swal from "sweetalert2";

export default async (tid: number): Promise<boolean> => {
    try {
        const userStore = useUserStore()
        console.log(userStore.userInfoData.id, tid)
        if (userStore.userInfoData.id == tid) {
            Swal.fire({
                title: "不能对自己操作哟!",
                heightAuto: false,
                confirmButtonColor: globalScss.colorButtonTheme,
                icon: "error",
            })
            return Promise.reject(false)
        }
        await attention(<AttentionReq>{
            uid: tid
        })
        return Promise.resolve(true)
    } catch (err: any) {
        Swal.fire({
            title: "操作失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
        return Promise.reject(false)
    }
}