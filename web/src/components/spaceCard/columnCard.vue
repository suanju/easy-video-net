<template>
    <div class="column">
        <div class="column-list">

        <!-- 骨架屏 -->
        <el-skeleton style="width: 100%; height: 18rem; margin-bottom: 8rem; " class="video-card"  v-for="(item,index) in list?.length ? list : quickCreationArr(6) "
                :key="item.id" :loading="!list?.length"
          animated>
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
                        <div class="lable-item" style="margin-right: 12px;" v-for="label in item.label" :key="label">
                            <SvgIcon name="label" class="icon-small" /> {{ label }}
                        </div>
                    </div>
                </div>
            </div>
          </template>
        </el-skeleton>
        </div>
    </div>

</template>

<script setup lang="ts">
import {GetArticleContributionListByUserResItem } from "@/types/live/liveRoom";
import dayjs from "dayjs"
import { useRouter } from 'vue-router';
import { VueEllipsis3 } from 'vue-ellipsis-3';
import { ArticleInfoList } from '@/types/space/space';

components: {
    VueEllipsis3
}
const props = defineProps({
    list: {
        type: Array as () => ArticleInfoList,
        required: true,
    }
})

const router = useRouter()

const mouseOver = (item: GetArticleContributionListByUserResItem) => {
    item.is_stay = true
}
const mouseleave = (item: GetArticleContributionListByUserResItem) => {
    item.is_stay = false
}

const quickCreationArr = (num :number) : Array<GetArticleContributionListByUserResItem> =>{
  let arr = []
  for(let i = 0; i < num ; i++){
    arr.push(<GetArticleContributionListByUserResItem>{
      
    })
  }
  return arr
}

const jumpArticle = (articleID: number) => {
    router.push({ name: "ArticleShow", query: { articleID } })
}
</script>
<style scoped lang="scss">
@import "./static/style/columnCard.scss"
</style>