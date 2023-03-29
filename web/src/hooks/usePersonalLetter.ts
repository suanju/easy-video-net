import { personalLetter } from "@/apis/personal";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useChatListStore } from "@/store/chat";
import { PersonalLetterReq } from "@/types/personal/chat/chat";
import Swal from "sweetalert2";

export default async (id: number) => {
    const chatListStore = useChatListStore()
    try {
        await personalLetter(<PersonalLetterReq>{
            id
        })
        //更新聊天列表
        useChatListStore()
        chatListStore.tid = id
        chatListStore.isShow = true

    } catch (err) {
        Swal.fire({
            title: "私信失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}