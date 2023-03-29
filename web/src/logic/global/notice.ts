import globalScss from "@/assets/styles/global/export.module.scss";
import { useUserStore } from "@/store/main";
import Swal from "sweetalert2";
import { useRouter } from "vue-router";



export const useInitNoticeSocket = () => {
    const router = useRouter()
    const userStore = useUserStore()
    let socket: WebSocket
    const open = () => {
        console.log("通知websocket连接成功 ")
    }
    const error = () => {
        console.error("通知websocket连接失败")
    }
    const getMessage = async (msg: any) => {
        let data = JSON.parse(msg.data)
        switch (data.type) {
            case "error":
                console.error("通知socket返回错误")
                break;
            case "messageNotice":
                messageNotice(data.data.unread, userStore)
                break;
        }
    }

    if (typeof (WebSocket) === "undefined") {
        Swal.fire({
            title: "您的浏览器不支持socket",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
        router.back()
        return
    } else {
        // 实例化socket
        socket = new WebSocket(import.meta.env.VITE_SOCKET_URL + "/ws/noticeSocket?token=" + userStore.userInfoData.token)
        // 监听socket连接
        socket.onopen = open
        // 监听socket错误信息
        socket.onerror = error
        // 监听socket消息
        socket.onmessage = getMessage
    }
}


//消息处理函数

const messageNotice = (num: number, userStore: any) => {
    userStore.setUnreadNotice(num)
}