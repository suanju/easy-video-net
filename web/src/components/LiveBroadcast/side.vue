<template>
    <div class="left-box" width="200px">
        <el-tabs v-model="activeName" stretch class="demo-tabs" @tab-click="handleClick">
            <!-- 弹幕留言 -->
            <el-tab-pane label="弹幕留言" name="first">
                <div class="barrage" @scroll="onScroll" ref="barrageMsgListRef">
                    <div>
                        <div class="barrage-msg-list">
                            <div v-for="barrageItem in dataList.baeeage" :key="barrageItem.msg_id">
                                <div class="barrage-msg-left" v-if="userInfo?.id != barrageItem.user_id">
                                    <el-avatar shape="circle" :size="34" :src="barrageItem.avatar" />
                                    <div class="msg-text">
                                        <span class="msg-span">{{ barrageItem.text }}</span>
                                    </div>
                                </div>
                                <div class="barrage-msg-right" v-if="userInfo?.id == barrageItem.user_id">
                                    <div class="msg-text">
                                        <span class="msg-span">{{ barrageItem.text }}</span>
                                    </div>
                                    <el-avatar shape="circle" :size="34" :src="barrageItem.avatar" />
                                </div>
                                <div></div>
                            </div>
                        </div>
                    </div>
                    <div class="barrage-prompt" v-show="prompt.show" @click="theBottom">
                    <div class="prompt"> 新弹幕来辽 {{prompt.num}} <SvgIcon name="drop-down" class="icon"></SvgIcon></div>
                    </div>
                    <div class="barrage-send">
                        <el-input v-model="barrage" resize="none" />
                        <el-button v-removeFocus type="primary" @click="sendMesage" :icon="Check" :disabled="!barrage"
                            circle />
                    </div>
                </div>
            </el-tab-pane>
            <!-- 在线人数 -->
            <el-tab-pane label="在线用户" name="second">
                <div class="online">
                    <div class="online-list">
                        <div class="online-item" v-for="onlineItem in dataList.online" :key="onlineItem.user_id">
                            <div class="online-info">
                                <el-avatar shape="circle" :size="38" :src="onlineItem.avatar" />
                                <span class="username">{{ onlineItem.username }}</span>
                            </div>

                        </div>
                    </div>
                </div>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>

<script setup lang="ts">
import { sideData } from "./types/side";
import { Check } from "@element-plus/icons-vue";
import { ElNotification, TabsPaneContext } from "element-plus";
import { ref, defineEmits, reactive, onMounted, onBeforeUpdate } from "vue";
import { encodeProtoFormat } from "@/utils/proto/proto";
import { WebClientSendBarrageReq, encodeWebClientSendBarrageReq, EnterLiveRoom, WebClientHistoricalBarrageRes } from "@/proto/pb/live";
import { WebClientSendBarrageRes } from "@/proto/pb/live";
import { vRemoveFocus } from "@/utils/customInstruction/focus";

defineProps({
    userInfo: Object,
});

const emit = defineEmits(["sendMessage"]);
const activeName = ref("first");

const barrageMsgListRef = ref()
const barrageMsgListPositioning = reactive({
    scrollHeight: 0,
    windowHeight: 0,
    scrollTop: 0,
    offsetHeight: 0,
})
const prompt = reactive({
    show : false,
    num : 0 ,
})

const dataList = <sideData>reactive({
    baeeage: [],
    online: [],
});
const barrage = ref("");

const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event);
};

const sendMesage = () => {
    if (!barrage.value) {
        ElNotification.success({
            title: "Warning",
            message: "",
            offset: 70,
        });
        return;
    }
    const msg = encodeProtoFormat(
        "webClientBarrageReq",
        encodeWebClientSendBarrageReq(<WebClientSendBarrageReq>{
            text: barrage.value,
            color: "#fff",
            type: "right",
        })
    );

    barrage.value = "";
    // console.log(emit("sendMessage"))
    emit("sendMessage", msg);
    theBottom()
};

const theBottom = () => {
    barrageMsgListRef.value.scrollTop = barrageMsgListRef.value.scrollHeight;
    prompt.show = false
    prompt.num = 0
}


const onScroll = () => {
    barrageMsgListPositioning.scrollHeight = barrageMsgListRef.value.scrollHeight
    barrageMsgListPositioning.windowHeight = barrageMsgListRef.value.clientHeight
    barrageMsgListPositioning.scrollTop = barrageMsgListRef.value.scrollTop
    barrageMsgListPositioning.offsetHeight = barrageMsgListRef.value.offsetHeight
    if (barrageMsgListPositioning.windowHeight + barrageMsgListPositioning.scrollTop < barrageMsgListPositioning.scrollHeight){
        prompt.show = false;
        prompt.num = 0
    } 

}

defineExpose({
    addBarrage(msg: WebClientSendBarrageRes) {
        if (barrageMsgListPositioning.windowHeight + barrageMsgListPositioning.scrollTop < barrageMsgListPositioning.scrollHeight) {
            //未在顶部
            dataList.baeeage = [...dataList.baeeage, msg];
            prompt.show = true
            prompt.num++
        } else {
            //在底部的话显示下一行
            dataList.baeeage = [...dataList.baeeage, msg];
            setTimeout(()=>{
                barrageMsgListRef.value.scrollTop = barrageMsgListRef.value.scrollHeight;
            },100)
        }
    },
    async addHistoryBarrage(msg: WebClientHistoricalBarrageRes) {
        if (!msg.list) return false;
        dataList.baeeage = [...dataList.baeeage, ...msg.list.reverse()];
        // barrageMsgListRef.value.scrollTop = barrageMsgListRef.value.scrollHeight;
    },
    updataOnline(online: Array<EnterLiveRoom>) {
        dataList.online = online
    },
});

onMounted(() => {

})
</script>

<style scoped lang="scss">
@import "./static/style/side.scss";
</style>
