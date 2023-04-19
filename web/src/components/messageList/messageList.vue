<template>
    <div class="message-box" v-click-outside="clickOutside">
        <div class="box-left">
            <div class="title">
                私信列表
            </div>
            <div class="message-list">
                <div class="message-item" v-for="item in messageList" :key="item.to_id" @click="toChat(item.to_id)">
                    <div class="item-left">
                        <el-badge is-dot :hidden="item.unread == 0" class="item">
                            <div class="avatar"> <el-avatar :size="38" :src="item.photo" />
                            </div>
                        </el-badge>

                    </div>
                    <div class=" item-right">
                        <div class="on">
                            <div class="name">{{ item.username }}</div>
                        </div>
                        <div class="info">
                            <VueEllipsis3 class="text" :visibleLine="1" :text="item.last_message">
                            </VueEllipsis3>
                            <div class="time">{{ dayjs(item.last_at).format("HH:mm") }}</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="box-right" v-if="chatListStore.tid">
            <div class="title">
                <div class="title-left">{{ chatListStore.tUsername }}</div>
                <div class="title-right">
                    <el-button class="exit" type="primary" size="small" round
                        @click="chatListStore.tid = 0">退出会话</el-button>
                    <el-popover :width="100" :teleported="false"
                        popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 8px;">
                        <template #reference>
                            <el-icon :size="18" class="icon">
                                <MoreFilled />
                            </el-icon>
                        </template>
                        <template #default>
                            <div class="more-box">
                                <div class="more-item" @click="deleteChat">删除聊天</div>
                            </div>
                        </template>
                    </el-popover>

                    <el-icon :size="18" class="icon" @click="chatListStore.isShow = false">
                        <Close />
                    </el-icon>

                </div>
            </div>
            <ChatBox :tid="chatListStore.tid" :msg-list="tMessageList"></ChatBox>
        </div>
        <div class="empty" v-if="!chatListStore.tid">
            <el-empty :image="emptyImg" description="快选择一位好友聊天叭~" />
        </div>
    </div>
</template>

<script  lang="ts" setup>

import { deleteChatItem } from "@/apis/personal";
import useLoadChatList from "@/hooks/useLoadChatList";
import { useChatListStore } from "@/store/chat";
import { DeleteChatItemReq } from "@/types/personal/chat/chat";
import { Close, MoreFilled } from '@element-plus/icons-vue';
import dayjs from "dayjs";
import { ClickOutside as vClickOutside } from 'element-plus';
import { computed, onMounted } from 'vue';
import { VueEllipsis3 } from 'vue-ellipsis-3';
import ChatBox from "./chatBox.vue";
import emptyImg from '@/components/messageList/static/img/unselectedMsg.png'
components: { 
    ChatBox
    VueEllipsis3
}

const chatListStore = useChatListStore()

const tMessageList = computed(() => {
    if (chatListStore.tid != 0) {
        let tInfo = chatListStore.chatListData.filter((item) => {
            return item.to_id == chatListStore.tid
        })
        if (tInfo[0]) {
            chatListStore.tUsername = tInfo[0]?.username
            return tInfo[0]?.message_list
        }
        return []
    }
})

const messageList = computed(() => {
    return chatListStore.chatListData.sort(function (a, b) { return new Date(b.last_at).getTime() - new Date(a.last_at).getTime() });
})


const loadData = async () => {
    useLoadChatList()
}

const clickOutside = () => {
    chatListStore.isShow = false
}

const deleteChat = async () => {
    try {
        await deleteChatItem(<DeleteChatItemReq>{
            id: chatListStore.tid
        })
        chatListStore.deleteItemByid(chatListStore.tid)
        chatListStore.tid = 0
    } catch (err) {
        console.log(err)
    }
}

const toChat = (id: number) => {
    chatListStore.tid = id
}


onMounted(() => {
    loadData()
})


</script>

<style lang="scss" scoped>
@import "./static/style/messageList.scss";
</style>