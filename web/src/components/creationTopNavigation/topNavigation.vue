<template>
  <el-row class="row ">
    <el-col :span="2">
      <div class="left">
        <img class="center-for-creators" src="./static/img/centerForCreators.png" />
      </div>
    </el-col>
    <el-col :span="16">
      <div class="middle">
        <div class="info" @click="jump('Home')">
          <SvgIcon name="bilibili-line" class="icon"></SvgIcon>主站
        </div>

      </div>
    </el-col>
    <el-col :span="6">
      <RightSide></RightSide>
    </el-col>

  </el-row>
</template>

<script lang="ts" setup>
import RightSide from "./rightSide/rightSide.vue";
import { useRouter } from "vue-router";

components: {
  RightSide;
}

//跳转
const jump = (name: string) => {
  router.push({
    name
  })
}
const router = useRouter();

//选项卡点击
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
  router.push(key);
};


</script>

<style scoped lang="scss">
@import "./static/style/topNavigation.scss";
</style>
