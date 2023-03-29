import globalScss from "@/assets/styles/global/export.module.scss";
import { useChatListStore } from "@/store/chat";
import { useUserStore } from "@/store/main";
import { ResultDataWs } from "@/types/idnex";
import { ChatOnlineUnreadNotice, ChatUnreadNotic } from "@/types/socket/chat";
import { ChatInfo, MessageInfo } from "@/types/store/chat";
import Swal from "sweetalert2";
import { useRouter } from "vue-router";



export const useInitChatSocket = () => {
    const router = useRouter()
    const userStore = useUserStore()
    console.log(useChatListStore())
    let socket: WebSocket
    const open = () => {
        console.log("聊天websocket连接成功 ")
    }
    const error = () => {
        console.error("聊天websocket连接失败")
    }
    const getMessage = async (msg: any) => {
        let data = <ResultDataWs>JSON.parse(msg.data)
        switch (data.type) {
            case "error":
                console.error("聊天socket返回错误")
                break;
            case "chatUnreadNotice":
                data as ResultDataWs<ChatUnreadNotic>
                chatUnreadNotice(data.data)
                break;
            case "chatOnlineUnreadNotice":
                data as ResultDataWs<ChatUnreadNotic>
                chatOnlineUnreadNotice(data.data)
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
        socket = new WebSocket(import.meta.env.VITE_SOCKET_URL + "/ws/chatSocket?token=" + userStore.userInfoData.token)
        // 监听socket连接
        socket.onopen = open
        // 监听socket错误信息
        socket.onerror = error
        // 监听socket消息
        socket.onmessage = getMessage
    }
}


//消息处理函数
const chatUnreadNotice = (data: ChatUnreadNotic) => {
    const chatListStore = useChatListStore()
    chatListStore.chatListData = chatListStore.chatListData.filter((item: ChatInfo) => {
        if (item.to_id == data.uid) {
            item.updated_at = data.last_message_info.created_at
            item.last_message = data.last_message
            item.unread = data.unread
            //添加到消息列表
            chatListStore.addMessage(data.last_message_info.uid, <MessageInfo>{
                uid: data.last_message_info.uid,
                username: data.last_message_info.username,
                photo: data.last_message_info.photo,
                tid: data.last_message_info.tid,
                message: data.last_message_info.message,
                type: data.last_message_info.type,
                created_at: data.last_message_info.created_at
            })
        }
        return item
    })
}

const chatOnlineUnreadNotice = (data: ChatOnlineUnreadNotice) => {
    const chatListStore = useChatListStore()
    try {
        const chatList = data.filter((item) => {
            item.message_list = item.message_list.reverse()
            return item
        })
        chatListStore.chatListData = chatList
    } catch (err: any) {
        console.error(err)
    }
}