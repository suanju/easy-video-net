<template>
    <div class="overall">
        <pageHeader title="视频稿件" icon-nmae="video" :animate="false" :whiteWhale="false"></pageHeader>
        <div class="content" v-loading="isLoading">
            <div class="video-list" v-show="videoList.length > 0" v-infinite-scroll="scrollBottom"
                infinite-scroll-delay="1000">
                <div :class="{ 'animate__animated': true, 'animate__fadeOutLeftBig': item.is_delete }"
                    v-for="(item, index) in videoList" :key="item.id" placement="top">
                    <div class="video-item">
                        <div class="item-left" @click="jump(item)">
                            <el-image class="item-img" :src="item.cover" fit="cover" />
                        </div>
                        <div class="item-right">
                            <div class="data">
                                <div class="item-title" @click="jump(item)">
                                    <VueEllipsis3 :visibleLine="1" :text="item.title">
                                    </VueEllipsis3>
                                </div>
                                <div class="item-info">
                                    <div class="icon-item">
                                        <SvgIcon name="video" class="icon" color="000"></SvgIcon> <span>{{ item.heat
                                        }}</span>
                                    </div>
                                    <div class="icon-item">
                                        <SvgIcon name="barrageNumber" class="icon" color="000"></SvgIcon> <span>{{
                                            item.barrageNumber }}</span>
                                    </div>
                                    <div class="icon-item">
                                        <SvgIcon name="comment" class="icon-max" color="000"></SvgIcon> <span>{{
                                            item.comments_number }}</span>
                                    </div>
                                </div>
                            </div>
                            <div class="function">
                                <el-button type="primary" v-removeFocus :icon="Edit" circle @click="editRecord(item)" />
                                <el-button type="info" v-removeFocus :icon="Delete" @click="delRecord(item.id)" circle />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="record-empty" v-show="videoList.length == 0 && isLoading == false">
                <el-empty description="还未发布视频,快去发布吧~" />
            </div>
            <div class="load-more" v-show="videoList.length > 0 && isLoadMore" v-loading="true">
            </div>
            <!-- 撑开底部 -->
            <div class="spread-bottom">
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { useDelRecord, useEditRecord, useJump, useLoadData, useVideoProp } from '@/logic/creation/manuscript/video';
import { GetVideoManagementListItem } from '@/types/creation/manuscript/video';
import { vRemoveFocus } from "@/utils/customInstruction/focus";
import { Delete, Edit } from '@element-plus/icons-vue';
import { watch } from 'vue';
import { VueEllipsis3 } from 'vue-ellipsis-3';

components: {
    VueEllipsis3
}

const { route, router, videoList, isLoading, pageInfo, loading, isLoadMore, isTheEnd, editVideoStore } = useVideoProp()

const delRecord = (id: number) => {
    useDelRecord(videoList, id)
}

const jump = (item: GetVideoManagementListItem) => {
    useJump(item, router)
}

const editRecord = (item: GetVideoManagementListItem) => {
    useEditRecord(item, loading, editVideoStore, router)
}

//加载底部
const scrollBottom = async () => {
    //无数据时取消加载更多
    if (isTheEnd.value) return false
    isLoadMore.value = true
    if (videoList.value.length <= 0) return false
    useLoadData(videoList, isLoading, pageInfo, isTheEnd)
    isLoadMore.value = false
}

watch(() => route.path, async () => {
    videoList.value = []
    isLoading.value = true
    useLoadData(videoList, isLoading, pageInfo, isTheEnd)

}, { immediate: true, deep: true })

</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/creation/manuscript/video.scss";
</style>
