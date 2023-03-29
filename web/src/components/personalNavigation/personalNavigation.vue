<template>
  <div class="side-navigation">
    <div>
      <div class="avatar-box">
        <el-avatar :size="110" :src="userInfo.userInfoData.photo" />
      </div>
      <p>{{ userInfo.userInfoData.username }}</p>
    </div>
    <el-menu class="el-menu-vertical" :collapse-transition="false" @open="handleOpen" @close="handleClose"
      @select="handleSelect">
      <el-sub-menu index="UserShow">
        <template #title>
          <SvgIcon name="user" class="icon"></SvgIcon>
          <span>用户信息</span>
        </template>
        <el-menu-item index="UserInfo">
          <div class="icon-item">
            <SvgIcon name="userData" class="icon-small"></SvgIcon>
            <span>个人信息</span>
          </div>
        </el-menu-item>
        <el-menu-item index="PictureSetting">
          <div class="icon-item">
            <SvgIcon name="pictures" class="icon-small"></SvgIcon>
            <span>头像设置</span>
          </div>
        </el-menu-item>
        <el-menu-item index="Safety">
          <div class="icon-item">
            <SvgIcon name="security" class="icon-small"></SvgIcon>
            <span>用户安全</span>
          </div>
        </el-menu-item>
      </el-sub-menu>
      <el-menu-item index="LiveSetUp">
        <div class="icon-item">
          <SvgIcon name="live" class="icon"></SvgIcon>
          <span>直播设置</span>
        </div>
      </el-menu-item>
      <el-sub-menu index="Favorites">
        <template #title>
          <SvgIcon name="collection" class="icon"></SvgIcon>
          <span>我的收藏</span>
        </template>
        <el-menu-item @click="createCollectDialogShow = true" index="CreateFavorites">
          <div class="icon-item">
            <SvgIcon name="folder" class="icon-smallxl"></SvgIcon>
            <span>创建收藏夹</span>
          </div>
        </el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="Record">
        <template #title>
          <SvgIcon name="playRecording" class="icon"></SvgIcon>
          <span>历史记录</span>
        </template>
        <el-menu-item @click="clearHistory()" index="ClearHistory">
          <div class="icon-item">
            <SvgIcon name="trashCan" class="icon"></SvgIcon>
            <span>清空历史</span>
          </div>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>

    <!-- 消息弹出框 -->
    <div class="dialog">
      <el-dialog v-model="createCollectDialogShow" title="创建收藏夹" width="50%" center :close-on-click-modal="true"
        :append-to-body="true" :before-close="createCollectDialogShowClose" align-center>
        <CreateFavorite :type="true" @shutDown="shutDown"></CreateFavorite>
      </el-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useUserStore } from "@/store/main";
import { useRouter } from "vue-router"
import CreateFavorite from "@/components/createFavorites/createFavorites.vue"
import { ref } from "vue";
import Swal from "sweetalert2";
import globalScss from "@/assets/styles/global/export.module.scss"
import { clearRecord } from "@/apis/personal";

comments: {
  CreateFavorite
}
const userInfo = useUserStore();
const router = useRouter();

const createCollectDialogShow = ref(false)

const handleOpen = (key: string, keyPath: string[]) => {
  router.push({ name: key })
};

const handleClose = (key: string, keyPath: string[]) => {
  router.push({ name: key })
};

const handleSelect = (key: string, keyPath: string[]) => {
  let arr = ["CreateFavorites", "ClearHistory"]
  if (key == "" || arr.indexOf(key) >= 0) return false
  router.push({ name: key })
}

const createCollectDialogShowClose = (done: () => void) => {
  done()
}

const clearHistory = () => {
  Swal.fire({
    heightAuto: false,
    title: '确认清空历史记录嘛?',
    icon: 'warning',
    confirmButtonColor: globalScss.colorButtonTheme,
    showCancelButton: true,
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    reverseButtons: true
  }).then(async (result) => {
    if (result.isConfirmed) {
      try {
        await clearRecord()
        Swal.fire({
          title: "清空成功",
          confirmButtonColor: globalScss.colorButtonTheme,
          heightAuto: false,
          icon: "success",
          preConfirm: () => {
            router.push({ name: "Record", query: { type: 'createTime' + Date.now() } })
          }
        })
      } catch (err: any) {
        console.log(err)
        Swal.fire({
          title: "清空失败",
          heightAuto: false,
          confirmButtonColor: globalScss.colorButtonTheme,
          icon: "error",
        })
      }
    } else if (result.dismiss === Swal.DismissReason.cancel) {
      console.log("取消")
    }
  })
}
//关闭createCollectDialogShow
const shutDown = () => {
  createCollectDialogShow.value = false
}

</script>

<style scoped lang="scss">
@import "./static/style/personalNavigation.scss";
</style>
