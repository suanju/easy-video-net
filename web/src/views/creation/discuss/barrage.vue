<template>
    <div class="overall">
        <pageHeader title="弹幕管理" icon-nmae="barrage" :animate="false" :whiteWhale="false"></pageHeader>
        <div class="content">
            <div class="comment-box" v-loading="isLoading" :infinite-scroll-delay="1000" :infinite-scroll-distance="40" v-infinite-scroll="scrollBottom"
                :infinite-scroll-disabled="isTheEnd" ref="scrollRef" :style="{ height: scrollHeight + 'px' }">
                <div class="comment-item" v-for="item in commentList" :key="item.id">
                    <div class="item-left">
                        <div class="avatar"><el-avatar :size="52" :src="item.photo" />
                        </div>
                        <div class="info">
                            <div class="top">
                                <div class="username"><span>{{ item.username }}</span></div>
                                <div class="time"><span>{{ dayjs(item.created_at).format('YYYY.MM.DD.hh.mm') }}</span></div>
                            </div>
                            <div class="comment-text">
                                <div class="comment-content">弹幕内容 : </div>
                                <VueEllipsis3 :visibleLine="1" :text="item.comment">
                                </VueEllipsis3>
                            </div>
                        </div>
                    </div>
                    <div class="item-right">
                        <div class="video-info">
                            <el-image class="item-img" :src="item.cover" fit="cover" />
                            <div class="title">
                                <VueEllipsis3 :visibleLine="1" :text="item.title">
                                </VueEllipsis3>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="load-more" v-show="commentList.length > 0 && isLoadMore" v-loading="true">
                </div>
                <!-- 撑开底部 -->
                <div class="spread-bottom">
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>

import { onMounted, ref , nextTick } from 'vue';
import globalScss from "@/assets/styles/global/export.module.scss"
import { VueEllipsis3 } from 'vue-ellipsis-3';
import { PageInfo } from '@/types/idnex';
import { getDiscussBarrageList } from '@/apis/contribution';
import dayjs from "dayjs"
import { GetDiscussBarrageListReq, GetDiscussBarrageListRes } from '@/types/creation/discuss/barrage';
import Swal from 'sweetalert2';

components: {
    VueEllipsis3
}

const commentList = ref(<GetDiscussBarrageListRes>[])
const pageInfo = ref(<PageInfo>{
    page: 1,
    size: 9,
})
const scrollHeight = ref(0)
const scrollRef = ref()
//是否首次加载
const isLoading = ref(true)
//是否正在加载更多
const isLoadMore = ref(false)
//是否全部加载完成
const isTheEnd = ref(false)


const loadData = async () => {
    try {
        const response = await getDiscussBarrageList(<GetDiscussBarrageListReq>{
            page_info: pageInfo.value
        })
        if (!response.data) return false
        if (response.data.length == 0) isTheEnd.value = true
        commentList.value = [...commentList.value, ...response.data]
        pageInfo.value.page++
        console.log(response)

    } catch (err) {
        console.log(err)
        Swal.fire({
            title: "加载数据失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}

//加载底部
const scrollBottom = async () => {
    // //无数据时取消加载更多
    if (isLoading.value == true) return false
    isLoadMore.value = true
    if (commentList.value.length <= 0) return false
    await loadData()
    isLoadMore.value = false

}


onMounted(async () => {
    await loadData()
    isLoading.value = false
    nextTick(()=>{
        scrollHeight.value = document.documentElement.clientHeight - scrollRef.value.offsetTop + 4
    })
})


</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/creation/discuss/barrage.scss";
</style>