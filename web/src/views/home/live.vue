<template>
    <div class="mian">
        <div class="head">
            <topNavigation color="#fff" :scroll="true" :displaySearch="false"></topNavigation>
        </div>
        <!-- 封面图 -->
        <div class="cover-picture">
        </div>
        <div class="title">
            <SvgIcon name="live" class="icon"></SvgIcon>推荐直播
        </div>
        <!-- 主体 -->
        <div class="middle">
            <!-- 直播 -->
            <el-skeleton style="width: 100%; height: 13rem;  margin-bottom: 6rem;" class="video-card" v-show="!isLoading"
                :loading="!beList.length" animated
                v-for="liveInfo in beList.length || isLoading ? beList : quickCreationArr(15) ">
                <template #template>
                    <el-skeleton-item variant="text" style="  width: 100%;height: 100%;" />
                    <div style="text-align: start; margin-top: 0.6rem;">
                        <el-skeleton-item variant="h3" style="width: 100%" />
                        <el-skeleton-item variant="h3" style="width: 100%" />
                    </div>
                </template>
                <template #default>
                    <div class="live-card">
                        <Card :id="liveInfo.id" :title="liveInfo.title" :username="liveInfo.username" :img="liveInfo.img"
                            :online="liveInfo.online" :photo="liveInfo.photo"></Card>
                    </div>
                </template>
            </el-skeleton>
        </div>
        <!-- 空状态 -->
        <div class="empty" v-show="beList.length == 0 && isLoading == true">
            <el-empty description="还没有人开播" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { getBeLiveList } from "@/apis/live";
import homeHeaderChannel from "@/components/homeHeaderChannel/homeHeaderChannel.vue";
import Card from "@/components/homeLiveCard/card.vue";
import homeRotograph from "@/components/homeRotograph/homeRotograph.vue";
import topNavigation from "@/components/topNavigation/topNavigation.vue";
import { BeLiveInfo, GetBeLiveListRes } from "@/types/home/live";
import { Ref, onMounted, ref } from "vue";

components: {
    homeRotograph
    Card
    homeHeaderChannel
    topNavigation
}

const beList = ref<GetBeLiveListRes>([])
const isLoading = ref(false)

const loadData = async (beList: Ref<GetBeLiveListRes>) => {
    const list = await getBeLiveList()
    isLoading.value = true
    if (!list.data) return
    beList.value = list.data
}


//生成占位骨架屏
const quickCreationArr = (num: number): Array<BeLiveInfo> => {
    let arr = []
    for (let i = 0; i < num; i++) {
        arr.push(<BeLiveInfo>{
        })
    }
    return arr
}

onMounted(() => {
    loadData(beList)
})

</script>

<style scoped lang="scss">
@import "@/assets/styles/views/home/live.scss";
</style>
