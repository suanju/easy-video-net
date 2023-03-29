<template>
  <div class="common-layout">
    <el-container>
      <el-header>
      <topNavigation/>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>


<script setup lang="ts">

import topNavigation from "@/components/topNavigation/topNavigation.vue"

components : {
    topNavigation
}



</script>
<style scoped lang="scss">

@import "@/assets/styles/views/Layout.scss"

</style>
