<template>
    <div class="overall">
        <pageHeader title="评论管理" icon-nmae="comment" :animate="false" :whiteWhale="false"></pageHeader>
        <div class="content">
            <div class="tabs">
                <el-tabs v-model="activeName" class="demo-tabs">
                    <el-tab-pane label="视频评论" name="video"><Video></Video></el-tab-pane>
                    <el-tab-pane label="专栏评论" name="article"><Article></Article></el-tab-pane>
                </el-tabs>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import type { TabsPaneContext } from 'element-plus'
import Video from "@/views/creation/discuss/comment/video.vue"
import Article from "@/views/creation/discuss/comment/article.vue"


components:{
    Video
    Article
}

const activeName = ref('video')



</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/creation/discuss/comment.scss";
</style>