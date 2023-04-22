<template>
    <div class="overall">
        <pageHeader title="历史记录" icon-nmae="playRecording"></pageHeader>
        <div class="content" v-loading="isLoading">
            <div class="timeline" v-show="recordList.length > 0" :infinite-scroll-delay="1000"
                :infinite-scroll-disabled="isTheEnd" :infinite-scroll-immediate="false" v-infinite-scroll="scrollBottom">
                <el-timeline>
                    <el-timeline-item :class="{ 'animate__animated': true, 'animate__fadeOutLeftBig': item.is_delete }"
                        v-for="(item, index) in recordList" :key="item.id" :timestamp="recordTimeFormat(item.updated_at)"
                        placement="top">
                        <div class="timeline-item">
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
                                        <el-avatar :size="38" :src="item.photo" />
                                        <span class="username">{{ item.username }}</span>
                                        <span class="partition">|</span>
                                        <span class="type">{{ item.type }}</span>
                                    </div>
                                </div>
                                <div class="function">
                                    <el-button type="info" v-removeFocus :icon="Delete" @click="delRecord(item.id)"
                                        circle />
                                </div>
                            </div>
                        </div>
                    </el-timeline-item>
                </el-timeline>
            </div>
            <div class="record-empty" v-show="recordList.length == 0 && isLoading == false">
                <el-empty description="暂无历史记录,快去逛逛叭~" />
            </div>
            <div class="load-more" v-show="recordList.length > 0 && isLoadMore" v-loading="true">
            </div>
            <!-- 撑开底部 -->
            <div class="spread-bottom">
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { useDelRecord, useJump, useLoadData, useRecordProp } from '@/logic/personal/record/record';
import { GetRecordListItem } from '@/types/personal/record/record';
import { recordTimeFormat } from "@/utils/conversion/timeConversion";
import { vRemoveFocus } from "@/utils/customInstruction/focus";
import { Delete } from '@element-plus/icons-vue';
import { watch } from 'vue';
import { VueEllipsis3 } from 'vue-ellipsis-3';


components: {
    VueEllipsis3
}

const { route, router, recordList, isLoading, pageInfo, loading, isLoadMore, isTheEnd } = useRecordProp()

const delRecord = (id: number) => {
    useDelRecord(recordList, id, loading)
}

const jump = (item: GetRecordListItem) => {
    useJump(item, router)
}

//加载底部
const scrollBottom = async () => {
    //无数据时取消加载更多
    if (isTheEnd.value) return false
    isLoadMore.value = true
    if (recordList.value.length <= 0) return false
    useLoadData(recordList, isLoading, pageInfo, isTheEnd)
    isLoadMore.value = false
}

watch(() => route.path, async () => {
    recordList.value = []
    isLoading.value = true
    useLoadData(recordList, isLoading, pageInfo, isTheEnd)

}, { immediate: true, deep: true })

</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/personal/record/record.scss";
</style>
