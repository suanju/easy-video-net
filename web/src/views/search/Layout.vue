<template>
    <div class="main">
        <div class="head">
            <topNavigation color="#000" :displaySearch="false"></topNavigation>
        </div>
        <div class="search-box">
            <el-input v-model="searchText" placeholder="最近新鲜事" size="large" @keyup.enter.native="searchChange">
                <template #suffix>
                    <el-icon @click="search">
                        <Search />
                    </el-icon>
                </template>
            </el-input>
        </div>
        <div class="menu">
            <el-menu :default-active="menuIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item index="VideoSearch">视频</el-menu-item>
                <el-menu-item index="UserSearch">用户</el-menu-item>
            </el-menu>
        </div>
        <div class="content">
            <router-view></router-view>
        </div>
    </div>
</template>

<script lang="ts" setup>
import topNavigation from "@/components/topNavigation/topNavigation.vue";
import { Search } from "@element-plus/icons-vue";
import { ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";


components: {
    topNavigation
}
const route = useRoute()
const router = useRouter();
const searchText = ref("");
const menuIndex = ref('VideoSearch');


const search = () => {
    if (!searchText.value) return false;
    router.push({
        name: 'VideoSearch', params: {
            text: searchText.value
        }
    })
}

const searchChange = () => {
    if (!searchText.value) return false;
    router.push({
        name: 'VideoSearch', params: {
            text: searchText.value
        }
    })
}

const handleSelect = (key: string, keyPath: string[]) => {
    router.push({ name: key })
}



watch(() => route.path, async () => {
    searchText.value = route.params.text as string
    //光标切换
    if (route.path.split("/")[route.path.split("/").length - 1] == "video") {
        menuIndex.value = "VideoSearch"
    } else if (route.path.split("/")[route.path.split("/").length - 1] == "user") {
        menuIndex.value = "UserSearch"
    }
}, { immediate: true, deep: true })
</script>


<style lang="scss" scoped>
@import "@/assets/styles/views/search/Layout.scss";
</style>