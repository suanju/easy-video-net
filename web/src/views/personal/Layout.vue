<template>
  <div class="personal">
  <div class="personal-layout animate__animated animate__rotateInUpLeft animate__faster">
        <el-container>
        <el-aside width="250px">
        <personalNavigation></personalNavigation>
        </el-aside>
        <el-main  style="overflow:hidden;">
            <router-view ></router-view>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<script setup lang="ts">
import personalNavigation from "@/components/personalNavigation/personalNavigation.vue"

components:{
    personalNavigation
}

</script>

<style scoped lang="scss">
@import "@/assets/styles/views/personal/Layout.scss";
</style>
