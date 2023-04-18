<template>
    <div class="content">
        <div class="video-list">
            <div class="title">
                <SvgIcon name="video" class="icon"></SvgIcon>视频作品
            </div>
            <div class="vidoe-box" v-show="releaseInformation?.videoList?.length > 0 || isLoading == false">
                <div class="video-item"
                    v-for="videoInfo in releaseInformation?.videoList?.length ? releaseInformation.videoList : quickCreationArr(8)">
                    <el-skeleton style="width: 100%; height: 13rem;" class="video-card"
                        :loading="!releaseInformation?.videoList?.length" animated>
                        <template #template>
                            <el-skeleton-item variant="text" style="  width: 100%;height: 100%;" />
                            <div style="text-align: start; margin-top: 0.6rem;">
                                <el-skeleton-item variant="h3" style="width: 100%" />
                                <div>
                                    <el-skeleton-item variant="h3" style="width: 80%" />
                                    <el-skeleton-item variant="h3" style="width: 60%" />
                                </div>
                            </div>
                        </template>
                        <template #default>
                            <VideoCard :id="videoInfo.id" :title="videoInfo.title" :username="videoInfo.username"
                                :video_duration="videoInfo.video_duration" v-bind:barrage-number="videoInfo.barrageNumber"
                                :heat="videoInfo.heat" :cover="videoInfo.cover" :created_at="videoInfo.created_at">
                            </VideoCard>
                        </template>
                    </el-skeleton>
                </div>
            </div>
            <div class="vidoe-empty" v-show="releaseInformation?.videoList?.length == 0 && isLoading == true">
                <el-empty description="还未发布视频~" />
            </div>
        </div>
        <div class="article-list">
            <div class="title">
                <SvgIcon name="column" class="icon"></SvgIcon>专栏作品
            </div>
            <div class="article-box" v-show="releaseInformation?.articleList?.length > 0 || isLoading == false">
                <Column :uid="spaceData.spaceInfoData.id"
                    :list="releaseInformation.articleList ? releaseInformation.articleList : []"></Column>
            </div>
            <div class="article-empty" v-show="releaseInformation?.articleList?.length == 0 && isLoading == true">
                <el-empty description="还未发布专栏~" />
            </div>
        </div>
        <!-- 撑开底部 -->
        <div class="spread-bottom">
        </div>
    </div>
</template>

<script lang="ts" setup>
import Column from "@/components/spaceCard/columnCard.vue";
import VideoCard from "@/components/spaceCard/videoCard.vue";
import { useInit, useSpaceProp } from "@/logic/space/space";
import { VideoInfo } from "@/types/space/space";
import { onMounted } from "vue";

const { spaceData, releaseInformation, isLoading } = useSpaceProp()

//生成占位骨架屏
const quickCreationArr = (num: number): Array<VideoInfo> => {
    let arr = []
    for (let i = 0; i < num; i++) {
        arr.push(<VideoInfo>{
        })
    }
    return arr
}

components: {
    Column
    VideoCard
}


onMounted(() => {
    useInit(spaceData.spaceInfoData.id, releaseInformation, isLoading)
})

</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/space/space.scss";
</style>
