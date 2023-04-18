<template>
    <div id="chat-box" class="chat-box" ref="boxRef" @scroll="boxScroll">
        <div class="chat-list">
            <div class="msg-list">
                <div v-for="msgItem in msgList" :key="msgItem.id">
                    <div class="msg-left" v-if="userStore.userInfoData.id != msgItem.uid">
                        <el-avatar shape="circle" :size="34" :src="msgItem.photo" />
                        <div class="msg-text">
                            <span class="msg-span">{{ msgItem.message }}</span>
                        </div>
                    </div>
                    <div class="msg-right" v-if="userStore.userInfoData.id == msgItem.uid">
                        <div class="msg-text">
                            <span class="msg-span">{{ msgItem.message }}</span>
                        </div>
                        <el-avatar shape="circle" :size="34" :src="userStore.userInfoData.photo" />
                    </div>
                    <div></div>
                </div>
            </div>
        </div>
        <div class="chat-input">
            <el-input v-model="input" resize="none" :input-style="{ width: '400px' }" autosize type="textarea"
                placeholder="聊点啥呢~" />
            <el-button v-removeFocus @click="sendText(input)" class="send" :type="input ? 'primary' : 'info'" :icon="Check"
                circle />
        </div>
    </div>
</template>

<script lang="ts" setup>
import { getChatHistoryMsg } from "@/apis/personal";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useChatListStore } from "@/store/chat";
import { useUserStore } from "@/store/main";
import { ResultDataWs } from "@/types/idnex";
import { GetChatHistoryMsgReq } from "@/types/personal/chat/chat";
import { ChatSendTextMsg } from "@/types/socket/chat";
import { MessageInfo } from '@/types/store/chat';
import { vRemoveFocus } from "@/utils/customInstruction/focus";
import { Check } from '@element-plus/icons-vue';
import { ElMessage } from "element-plus";
import Swal from "sweetalert2";
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const props = defineProps({
    tid: {
        type: Number,
        required: true,
    },
    msgList: {
        type: Array as () => Array<MessageInfo> | undefined,
        required: true,
    }
})

var socket: WebSocket | undefined
const input = ref("")
const userStore = useUserStore()
const chatListStore = useChatListStore()
const router = useRouter()
const tid = ref(0)
const boxRef = ref()

const loadSocket = () => {
    let socket: WebSocket
    const open = () => {
        console.log("用户聊天websocket连接成功 ")
    }
    const error = (err: any) => {
        console.log(err)
        console.error("用户聊天聊天websocket连接失败")
    }
    const getMessage = async (msg: any) => {
        let data = <ResultDataWs>JSON.parse(msg.data)
        switch (data.type) {
            case "error":
                console.error("用户聊天聊天socket返回错误")
                ElMessage({
                    message: data.message,
                    type: 'error',
                    appendTo: document.getElementById("chat-box") as HTMLElement,
                })
                break;
            case "chatSendTextMsg":
                data as ResultDataWs<ChatSendTextMsg>
                chatSendTextMsg(data.data)
                break
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
        socket = new WebSocket(import.meta.env.VITE_SOCKET_URL + "/ws/chatUserSocket?token=" + userStore.userInfoData.token + "&tid=" + tid.value)
        // 监听socket连接
        socket.onopen = open
        // 监听socket错误信息
        socket.onerror = error
        // 监听socket消息
        socket.onmessage = getMessage
    }
    return socket
}

const send = (type: string, msg: string) => {
    let data = JSON.stringify({
        type,
        "data": msg,
    })
    socket?.send(data)
}

const sendText = (msg: string) => {
    if (msg == "") return false
    send("sendChatMsgText", msg)
    //添加记录
    input.value = ""
    chatListStore.addMessage(tid.value, <MessageInfo>{
        uid: userStore.userInfoData.id,
        username: userStore.userInfoData.username,
        photo: userStore.userInfoData.photo,
        tid: tid.value,
        message: msg,
        type: "text",
        created_at: Date().toString()
    })
    //滚动到底部
    nextTick(() => {
        rollingBottom()
    })
}

const chatSendTextMsg = (data: ChatSendTextMsg) => {
    chatListStore.addMessage(data.uid, data)
    nextTick(() => {
        rollingBottom()
    })

}


const rollingBottom = () => {
    //执行了
    if (props.msgList) {
        nextTick(() => {
            if (boxRef?.value?.scrollHeight) {
                boxRef.value.scrollTop = boxRef?.value?.scrollHeight
            }

        })
    }

}

const boxScroll = async () => {
    console.log("触碰顶部")

    if (boxRef.value.scrollTop == 0) {
        //触碰顶加载更多‘
        const h = boxRef.value.scrollHeight
        try {
            let mixTime: number | string = new Date().getTime() //最小值默认当前时间
            chatListStore.chatListData.filter((item) => {
                if (item.to_id == chatListStore.tid) {
                    item.message_list.filter((ml) => {
                        console.log(new Date(ml.created_at).getTime() - new Date(mixTime).getTime())
                        if (new Date(ml.created_at).getTime() - new Date(mixTime).getTime() < 0) {
                            mixTime = ml.created_at
                        }
                    })
                }
            })
            console.log("最小值", mixTime)
            const response = await getChatHistoryMsg(<GetChatHistoryMsgReq>{
                tid: props.tid,
                last_time: mixTime
            })
            if (!response.data) return false
            const chatList = response.data.reverse()
            chatListStore.chatListData = chatListStore.chatListData.filter((item) => {
                if (item.to_id == chatListStore.tid) {
                    item.message_list = [...chatList, ...item.message_list]
                }
                return item
            })
            nextTick(() => {
                boxRef.value.scrollTop = boxRef.value.scrollHeight - h
            })

            console.log(response)
        } catch (err) {
            console.log(err)
        }
    }
}

const watchTid = watch(() => { return chatListStore.tid }, (newVal, oldVal) => {
    socket?.close()
    if (newVal != oldVal) {
        tid.value = newVal
        chatListStore.clearUnread(tid.value)
        socket = loadSocket()
        rollingBottom()
    }
}, { immediate: true })



onMounted(() => {

})

//关闭时结束监听和socket
onUnmounted(() => {
    watchTid()
    socket?.close()
})

</script>

<style lang="scss" scoped>
@import "./static/style/chatBox.scss";
</style>