<template>
    <div class="column">
        <div class="column-list" v-show="!isLoading || columnList.length > 0" v-infinite-scroll="scrollBottom"
            infinite-scroll-delay="1000">
            <!-- 骨架屏 -->
            <el-skeleton style="width: 100%; height: 18rem; margin-bottom: 8rem; " class="video-card"
                v-for="(item, index) in columnList.length ? columnList : quickCreationArr(6) " :key="item.id"
                :loading="!columnList.length" animated>
                <template #template>
                    <el-skeleton-item variant="text" style="  width: 100%;height: 100%;" />
                    <div style="text-align: start; margin-top: 0.2rem;">
                        <el-skeleton-item variant="h3" style="width: 100%" />
                        <div>
                            <el-skeleton-item variant="h3" style="width: 100%" />
                        </div>
                    </div>
                </template>
                <template #default>
                    <!-- 单个卡片 -->
                    <div :class="{ mouseover: item.is_stay, mouseleave: !item.is_stay }" class="column-item shadow-box "
                        @mouseover="mouseOver(item)" @mouseleave="mouseleave(item)" @click="jumpArticle(item.id)">
                        <div :class="{ 'item-image': true, 'right': index % 2 == 1 }">
                            <img :src="item.cover" class="el-image__inner image" style="object-fit: cover;">
                        </div>
                        <div class="item-text">
                            <div class="post-meta">
                                <SvgIcon name="camera" class="icon-small"></SvgIcon> " 发布于
                                {{ dayjs(item.created_at).format('YYYY.MM.DD.hh.mm') }} "
                            </div>
                            <h3>{{ item.title }}</h3>
                            <div class="post-meta" style="margin-bottom: 15px;">
                                <SvgIcon name="hot" class="icon-small"></SvgIcon>
                                <span>
                                    {{ item.heat }} 热度
                                </span>
                                <SvgIcon name="comments" class="icon-small"></SvgIcon>
                                <span>
                                    {{ item.comments_number }} 条评论
                                </span>
                                <SvgIcon name="like" class="icon-small"></SvgIcon>
                                <span>
                                    {{ item.likes_number }}点赞
                                </span>
                            </div>
                            <div class="recent-post-desc">
                                <VueEllipsis3 :text="item.content" :visibleLine="4">
                                    <template v-slot:ellipsisNode>
                                        <span>...</span>
                                    </template>
                                </VueEllipsis3>
                            </div>
                            <div class="sort-label">
                                <div class="lable-item" style="margin-right: 12px;">
                                    <SvgIcon name="classification" class="icon-small" /> {{ item.classification }}
                                </div>
                                <div class="lable-item" style="margin-right: 12px;" v-for="label in item.label" :key="label"
                                    v-show="label">
                                    <SvgIcon name="label" class="icon-small" /> {{ label }}
                                </div>
                            </div>
                        </div>
                    </div>
                </template>
            </el-skeleton>
        </div>
        <div class="load-more" v-show="isLoadMore" v-loading="isLoadMore">
        </div>
        <!-- 撑开底部 -->
        <div class="no-more" v-show="isTheEnd">
            没有更多了~
        </div>
        <div class="spread-bottom">
        </div>
    </div>
    <div class="column-empty" v-show="columnList.length == 0 && isLoading == true">
        <el-empty description="还没有人发布专栏,快去发布叭~" />
    </div>
</template>

<script setup lang="ts">

import { getArticleContributionList } from "@/apis/contribution";
import { GetArticleContributionListReq } from "@/types/home/column";
import { PageInfo } from '@/types/idnex';
import { GetArticleContributionListByUserRes, GetArticleContributionListByUserResItem } from "@/types/live/liveRoom";
import dayjs from "dayjs";
import { onMounted, ref } from 'vue';
import { VueEllipsis3 } from 'vue-ellipsis-3';
import { useRouter } from 'vue-router';

components: {
    VueEllipsis3
}

//专栏列表
const columnList = ref<GetArticleContributionListByUserRes>([])
const router = useRouter()
const pageInfo = ref(<PageInfo>{
    page: 1,
    size: 6
})
//是否首次加载
const isLoading = ref(false)
//是否正在加载更多
const isLoadMore = ref(false)
//是否全部加载完成
const isTheEnd = ref(false)


//加载底部
const scrollBottom = async () => {
    if (!isLoading.value) return false
    if (isTheEnd.value) return false
    //无数据时取消加载更多
    if (columnList.value.length <= 0) return false
    isLoadMore.value = true
    await loadData()

}

const mouseOver = (item: GetArticleContributionListByUserResItem) => {
    item.is_stay = true
}
const mouseleave = (item: GetArticleContributionListByUserResItem) => {
    item.is_stay = false
}

const loadData = async () => {
    setTimeout(async () => {
        try {
            const response = await getArticleContributionList(<GetArticleContributionListReq>{
                page_info: pageInfo.value
            })
            if (!response.data) return false
            if (response.data.length == 0) isTheEnd.value = true
            //加上是否停留鼠标
            response.data?.filter((item) => {
                item.is_stay = false
                return true
            })
            columnList.value = [...columnList.value, ...response.data]
            console.log(columnList)
            //请求成功后下次分页+1
            pageInfo.value.page++
            isLoading.value = true
            isLoadMore.value = false
        } catch (err) {
            console.log(err)
        }
    }, 500);

}

const quickCreationArr = (num: number): Array<GetArticleContributionListByUserResItem> => {
    let arr = []
    for (let i = 0; i < num; i++) {
        arr.push(<GetArticleContributionListByUserResItem>{

        })
    }
    return arr
}

const jumpArticle = (articleID: number) => {
    router.push({ name: "ArticleShow", query: { articleID } })
}
onMounted(() => {
    loadData()
})
</script>
<style scoped lang="scss">
@import "./static/style/column.scss"
</style>