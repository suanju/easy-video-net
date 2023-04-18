<template>
    <div class="overall">
        <pageHeader title="空间" icon-nmae="planet"></pageHeader>
        <div :class="{ content: true, scroll: releaseInformation.videoList?.length }" ref="scrollRef"
            :style="{ width: '100%', height: scrollHeight + 'px' }">
            <div class="video-list">
                <div class="title">
                    <SvgIcon name="video" class="icon"></SvgIcon>视频作品
                </div>
                <div class="vidoe-box" v-show="releaseInformation?.videoList?.length > 0 || isLoading == false">
                    <div class="video-item"
                        v-for="videoInfo in releaseInformation?.videoList?.length ? releaseInformation.videoList : quickCreationArr(8)">
                        <el-skeleton style="width: 100%; height: 154px;" class="video-card"
                            :loading="!releaseInformation?.videoList?.length" animated>
                            <template #template>
                                <el-skeleton-item variant="text" style="  width: 100%;height: 100%;" />
                                <div style="text-align: start; margin-top: 0.6rem;">
                                    <el-skeleton-item variant="text" style="width: 100%" />
                                    <div>
                                        <el-skeleton-item variant="text" style="width: 100%" />
                                        <el-skeleton-item variant="text" style="width: 100%" />
                                    </div>
                                </div>
                            </template>
                            <template #default>
                                <VideoCard :id="videoInfo.id" :title="videoInfo.title" :username="videoInfo.username"
                                    :video_duration="videoInfo.video_duration"
                                    v-bind:barrage-number="videoInfo.barrageNumber" :heat="videoInfo.heat"
                                    :cover="videoInfo.cover" :created_at="videoInfo.created_at">
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
                    <Column :uid="userStore.userInfoData.id"
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
    </div>
</template>

<script lang="ts" setup>
import Column from "@/components/userShowCard/columnCard.vue";
import VideoCard from "@/components/userShowCard/videoCard.vue";
import { useInit, useUseUserShowProp } from "@/logic/personal/userInfo/userShow";
import { VideoInfo } from "@/types/space/space";
import { onMounted, ref } from "vue";

const scrollHeight = ref(0)
const scrollRef = ref()
const { userStore, releaseInformation, isLoading } = useUseUserShowProp()

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
    scrollHeight.value = document.documentElement.clientHeight - scrollRef.value.offsetTop
    useInit(userStore.userInfoData.id, releaseInformation, isLoading)
})

</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/personal/userInfo/userShow.scss";
</style>
