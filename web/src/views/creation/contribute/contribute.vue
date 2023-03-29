<template>
    <div class="content-box">
        <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
            <el-tab-pane label="视频投稿" name="video" v-if="type == 'editVideo' || type == '' ">
                <videoContribution :type="type == 'editVideo' ? 'edit' : 'create'"></videoContribution>
            </el-tab-pane>
            <el-tab-pane label="专栏投稿" name="article"  v-if="type == 'editArticle' || type == '' ">
                <articleContribution :type="type == 'editArticle' ? 'edit' : 'create'"></articleContribution>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>


<script lang="ts" setup>
import { onMounted, ref, nextTick } from 'vue'
import type { TabsPaneContext } from 'element-plus'
import { useRoute } from 'vue-router'
import videoContribution from '@/views/creation/contribute/contributePage/vdeoContribution.vue'
import articleContribution from "@/views/creation/contribute/contributePage/articleContribution.vue"

components: {
    videoContribution
    articleContribution
}

const activeName = ref('video')
const route = useRoute()
const type = ref("")

const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}

onMounted(() => {
    nextTick(() => {
        type.value = route.query.type as string
        if(type.value == undefined) type.value = ""
        if (type.value == "editVideo") {
            activeName.value = "video"
        } else if (type.value == "editArticle") {
            activeName.value = "article"
        } else {
            activeName.value = "video"
        }
    })
})


</script>
<style scoped lang="scss">
@import "@/assets/styles/views/creation/contribute/contribute.scss";
</style>