<template>
  <div class="side-navigation">
    <div class="contribute">
      <el-button class="btn" v-removeFocus type="primary" @click="jump('Contribute')" round>投稿</el-button>
    </div>
    <el-menu class="el-menu-vertical" :collapse-transition="false" @open="handleOpen" @close="handleClose"
      @select="handleSelect">
      <!-- <el-menu-item index="pictureSetting">
        <div class="icon-item">
          <SvgIcon name="home" class="icon"></SvgIcon>
          <span>首页</span>
        </div>
      </el-menu-item> -->
      <el-sub-menu index="Manuscript">
        <template #title>
          <SvgIcon name="eg" class="icon"></SvgIcon>
          <span>稿件管理</span>
        </template>
        <el-menu-item index="VideoManagement">
          <div class="icon-item">
            <SvgIcon name="video" class="icon"></SvgIcon>
            <span>视频稿件</span>
          </div>
        </el-menu-item>
        <el-menu-item index="ArticleManagement">
          <div class="icon-item">
            <SvgIcon name="column" class="icon"></SvgIcon>
            <span>专栏稿件</span>
          </div>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="Discuss">
        <template #title>
          <SvgIcon name="discuss" class="icon"></SvgIcon>
          <span>互动管理</span>
        </template>
        <el-menu-item index="CommentDiscuss">
          <div class="icon-item">
            <SvgIcon name="comment" class="icon-xl"></SvgIcon>
            <span>评论管理</span>
          </div>
        </el-menu-item>
        <el-menu-item index="BarrageDiscuss">
          <div class="icon-item">
            <SvgIcon name="barrage" class="icon-small"></SvgIcon>
            <span>弹幕管理</span>
          </div>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script lang="ts" setup>
import { useUserStore } from "@/store/main";
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { useRouter } from "vue-router"

const userInfo = useUserStore();
const router = useRouter();
let arr = ["Manuscript","Discuss"];

const handleOpen = (key: string, keyPath: string[]) => {
  if (key == "" || arr.indexOf(key) >= 0) return false
  router.push({ name: key })
};
const handleClose = (key: string, keyPath: string[]) => {
  if (key == "" || arr.indexOf(key) >= 0) return false
  router.push({ name: key })
};

const handleSelect = (key: string, keyPath: string[]) => {
  if (key == "" || arr.indexOf(key) >= 0) return false
  router.push({ name: key })

}

//跳转
const jump = (name: string) => {
  router.push({
    name
  })
}

</script>

<style scoped lang="scss">
@import "./static/style/creationNavigation.scss";
</style>
