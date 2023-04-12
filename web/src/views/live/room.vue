<template>
  <div class="overall">
    <div class="head animate__animated animate__slideInDown">
      <div class="common-row">
        <el-row :gutter="20">
          <el-col :span="18">
            <div class="live-box">
              <LiveHeader :src="liveInfo?.photo ? liveInfo?.photo : ''"
                :titel="liveInfo?.live_title ? liveInfo?.live_title : ''" />
              <div ref="videoRef" class="player" id="dplay" />
            </div>
          </el-col>
          <el-col :span="6">
            <Side ref="sideRef" @sendMessage="sendMessage" :userInfo="userStore.userInfoData"></Side>
          </el-col>
        </el-row>
      </div>
    </div>

    <div class="content">
      <div class="content-row">
        <h3>Ta 的专栏</h3>
        <el-row :gutter="20">
          <el-col :span="16">
            <div class="dynamic-box">
              <div class="dynamic">
                <Column :roomID="Number(route.params.id)" />
              </div>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="recommended">
              <Announcement />
            </div>
          </el-col>
        </el-row>
      </div>

    </div>
  </div>
</template>
<script lang="ts"  setup >
import Announcement from "@/components/LiveBroadcast/announcement.vue";
import Column from "@/components/LiveBroadcast/column.vue";
import LiveHeader from "@/components/LiveBroadcast/liveHeader.vue";
import Side from "@/components/LiveBroadcast/side.vue";
import { useInit, useLiveRoomProp, useWebSocket } from "@/logic/live/liveRoom";
import DPlayer from "dplayer";
import { onMounted, ref } from "vue";

components: {
  LiveHeader
  Side
  Column
  Announcement
}
const sideRef = ref()
var dp: DPlayer //播放器配置对象
const { videoRef, userStore, route, router, roomID, liveInfo } = useLiveRoomProp()


const sendMessage = ref((tset: string) => {
})

onMounted(async () => {
  dp = await useInit(videoRef, route, router, roomID, liveInfo) as DPlayer
  sendMessage.value = useWebSocket(dp, userStore, sideRef, roomID, router).sendMessage
})



</script>

<style scoped lang="scss">
@import "@/assets/styles/views/live/room.scss"
</style>
