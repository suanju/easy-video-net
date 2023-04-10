<template>
  <div class="info">
    <el-popover :teleported="false" :width="120" trigger="hover" @show="chatListStore.isShow = false"
      popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px;">
      <template #reference>
        <div class="avatar-box">
          <router-link v-if="userInfo.userInfoData.token" to="/personal" class="avatar">
            <el-avatar :size="36" :src="userInfo.userInfoData.photo" />
          </router-link>
        </div>
      </template>
      <template #default>
        <div class="user-selection">
          <div class="selection-item" @click="jump('Personal')">
            <SvgIcon name="user" class="icon" color="#000"></SvgIcon>用户中心
          </div>
          <div class="selection-item  mt" @click="loginOut">
            <SvgIcon name="exit" class="icon" color="#000"></SvgIcon>退出登入
          </div>
        </div>
      </template>
    </el-popover>
    <div v-if="!userInfo.userInfoData.token" @click="login()" class="login">
      <span>登入</span>
    </div>


    <!-- iocn -->
    <div class="icon-list">
      <el-popover ref="popover" :teleported="false" :width="400" trigger="hover" @hide="noticePopoverHide()"
        @show="noticePopoverShow()"
        popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px;padding: 14px 0px 14px 14px;">
        <template #reference>
          <el-badge is-dot :hidden="userInfo.userInfoData.unreadNotice == 0">
            <div class="icon-item">
              <SvgIcon name="notice" class="icon" :color="iconColor">
                <div class="red-num-message">1</div>
              </SvgIcon>
              <p :style="{ color: iconColor }">消息</p>
            </div>
          </el-badge>
        </template>
        <template #default>
          <NoticeList ref="noticeListRef"></NoticeList>
        </template>
      </el-popover>


      <div class="icon-item" @click="notOpen()">
        <SvgIcon name="dynamic" class="icon" :color="iconColor"></SvgIcon>
        <p :style="{ color: iconColor }">动态</p>
      </div>

      <el-popover :visible="chatListStore.isShow" :teleported="false" :width="720" trigger="hover"
        popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px;padding: 14px 0px 14px 14px;">
        <template #reference>
          <el-badge is-dot :hidden="chatUnreadMessage == 0">
            <div class="icon-item" @mouseover="chatListStore.isShow = true">
              <SvgIcon name="message" class="icon" :color="iconColor"></SvgIcon>
              <p :style="{ color: iconColor }">私信</p>
            </div>
          </el-badge>
        </template>
        <template #default>
          <MessageList v-if="chatListStore.isShow"></MessageList>
        </template>
      </el-popover>

      <div class="icon-item" @click="jump('Record')">
        <SvgIcon name="history" class="icon" :color="iconColor"></SvgIcon>
        <p :style="{ color: iconColor }">历史</p>
      </div>

      <div class="icon-item" @click="jump('Contribute')">
        <SvgIcon name="creation" class="icon" :color="iconColor"></SvgIcon>
        <p :style="{ color: iconColor }"> 创作</p>
      </div>

    </div>

    <el-button type="primary" round @click="startLive()">开始直播</el-button>
    <el-dialog v-model="dialogTableVisible" :lock-scroll="false" class="dialog" center title="Begin to live">
      <el-steps :active="nextIndex">
        <el-step title="准备工作" description="Download tool" />
        <el-step title="设置参数" description="Set the parameters" />
        <el-step title="开始直播" description="Start to live" />
      </el-steps>

      <el-row class="steps">
        <el-col :span="18">
          <div class="steps-left">
            <p v-show="nextIndex == 1">下载 OBS Studio 直播工具, 安装到我的电脑打开</p>
            <div v-show="nextIndex == 2">
              <h4>Adders : {{ userInfo.userInfoData.liveData.address }}</h4>
              <h4>
                key : {{ liveKeyDesensitization(userInfo.userInfoData.liveData.key) }}
                <el-button @click="copy(userInfo.userInfoData.liveData.key)" class="copy" color="#626aef" size="small"
                  plain round>
                  copy
                </el-button>
              </h4>
            </div>
            <p v-show="nextIndex == 3">配置完成！开始直播咯~</p>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="steps-right">
            <el-button type="primary" @click="nextStep">
              <span v-show="nextIndex < 3">下一步</span>
              <span v-show="nextIndex == 3">完成</span>
              <el-icon class="el-icon--right">
                <ArrowRight />
              </el-icon>
            </el-button>
          </div>
        </el-col>
      </el-row>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { getLiveRoom } from "@/apis/live";
import globalScss from "@/assets/styles/global/export.module.scss";
import MessageList from "@/components/messageList/messageList.vue";
import NoticeList from "@/components/notice/noticeList.vue";
import { useChatListStore } from "@/store/chat";
import { useUserStore } from "@/store/main";
import { liveKeyDesensitization } from "@/utils/conversion/stringConversion";
import { ArrowRight } from "@element-plus/icons-vue";
import { ElNotification } from "element-plus";
import Swal from "sweetalert2";
import { onMounted, onUnmounted, ref, watch } from "vue";
import useClipboard from "vue-clipboard3";
import { useRouter } from "vue-router";

components: {
  NoticeList
  MessageList
}
const chatListStore = useChatListStore()
const userInfo = useUserStore();
const router = useRouter()
const dialogTableVisible = ref(false);
let nextIndex = ref(1);

const emit = defineProps({
  iconColor: {
    type: String,
    default: '#18191C',
  }
}
)

const nextStep = () => {
  if (nextIndex.value >= 3) {
    dialogTableVisible.value = !dialogTableVisible.value;
    //关窗动画结束后修改
    setTimeout(() => {
      nextIndex.value = 1;
    }, 1000);
  } else {
    nextIndex.value = nextIndex.value + 1;
  }
};

//跳转
const jump = (name: string) => {
  router.push({
    name
  })
}

const getLiveRoomInfo = async () => {
  if (userInfo.userInfoData.token) {
    try {
      const data = await getLiveRoom();
      userInfo.userInfoData.liveData.address = data.data?.address || "";
      userInfo.userInfoData.liveData.key = data.data?.key || "";
    } catch (err) {
      console.log(err);
    }
  }

};

const { toClipboard } = useClipboard();

const startLive = () => {
  if (userInfo.userInfoData.token) {
    dialogTableVisible.value = !dialogTableVisible.value
    getLiveRoomInfo();
  } else {
    router.push({
      name: 'Login'
    })
  }
}

const login = () => {
  router.push({
    name: 'Login'
  })
}

const copy = async (text: string) => {
  try {
    await toClipboard(text); //实现复制
    ElNotification({
      title: "Success",
      message: "复制成功",
      type: "success",
    });
    console.log("Success");
  } catch (e) {
    console.error(e);
  }
};

// //消息通知相关
const noticeListRef = ref()

const noticePopoverShow = () => {
  chatListStore.isShow = false
  noticeListRef.value.init()
}
const noticePopoverHide = () => {
  noticeListRef.value.end()
}

//未读消息数量
const chatUnreadMessage = ref(0)
//监听未读消息
const watchChatUnreadMessage = watch(() => { return chatListStore.chatListData }, (newVal, oldVal) => {
  chatUnreadMessage.value = 0
  newVal.filter((item) => {
    chatUnreadMessage.value += item.unread
  })
}, { immediate: true, deep: true })

//退出登入
const loginOut = () => {
  userInfo.loginOut()
  router.push({
    name: "Login"
  })
}


const notOpen = () => {
  Swal.fire({
    title: "敬请期待",
    heightAuto: false,
    confirmButtonColor: globalScss.colorButtonTheme,
    icon: "info",
  })
}


onMounted(() => {
  //刷新时关闭
  chatListStore.isShow = false
})

onUnmounted(() => {
  //清除监听
  watchChatUnreadMessage()
})

</script>

<style scoped lang="scss">
@import "../static/style/rightSide.scss";
</style>
