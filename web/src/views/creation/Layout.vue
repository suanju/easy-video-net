<template>
    <div class="common-layout">
        <el-container>
            <el-header class="animate__animated animate__slideInDown">
                <topNavigation />
            </el-header>
            <el-main >
                <div class="common-layout">
                    <el-container>
                        <el-aside width="200px"   class=" animate__animated animate__bounceInLeft"><creationlNavigation></creationlNavigation></el-aside>
                        <el-main class="creation-main gray">
                        <div class="creation-content">
                            <router-view></router-view>
                        </div>
                        </el-main>
                    </el-container>
                </div>
            </el-main>
        </el-container>
    </div>
</template>
  
  
<script setup lang="ts">

import topNavigation from "@/components/creationTopNavigation/topNavigation.vue"
import creationlNavigation from "@/components/creationNavigation/creationlNavigation.vue"
components: {
    topNavigation
    creationlNavigation
}



</script>
<style scoped lang="scss">
@import "@/assets/styles/views/Layout.scss"
</style>
  