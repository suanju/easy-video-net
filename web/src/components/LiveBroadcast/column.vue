<template>
    <div class="column">

        <div class="column-list">
            <!-- 单个卡片 -->
            <div :class="{ mouseover: item.is_stay, mouseleave: !item.is_stay }" class="column-item shadow-box "
                @mouseover="mouseOver(item)" @mouseleave="mouseleave(item)" v-for="(item,index) in columnList" :key="item.id"
                @click="jumpArticle(item.id)">
                <div :class="{'item-image': true ,'right' : index%2 == 1}">
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
                        <span class="lable-item" style="margin-right: 12px;" v-for="label in item.label" :key="label">
                            <SvgIcon name="label" class="icon-small" /> {{ label }}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>

</template>

<script setup lang="ts">

import { getArticleContributionListByUser } from '@/apis/contribution';
import { GetArticleContributionListByUserReq } from '@/types/creation/contribute/contributePage/articleContribution';
import { GetArticleContributionListByUserRes, GetArticleContributionListByUserResItem } from "@/types/live/liveRoom";
import { onMounted, reactive, ref } from 'vue';
import dayjs from "dayjs"
import { useRouter } from 'vue-router';
import { VueEllipsis3 } from 'vue-ellipsis-3';

components: {
    VueEllipsis3
}


const props = defineProps({
    roomID: {
        type: Number,
        required: true,
    }
}
)
const router = useRouter()

//专栏列表
const columnList = ref<GetArticleContributionListByUserRes>([])
const mouseOver = (item: GetArticleContributionListByUserResItem) => {
    item.is_stay = true
}
const mouseleave = (item: GetArticleContributionListByUserResItem) => {
    item.is_stay = false
}

const init = async () => {
    try {
        console.log(props)
        let requistData = <GetArticleContributionListByUserReq>{
            userID: props.roomID
        }
        const response = await getArticleContributionListByUser(requistData)
        //加上是否停留鼠标
        response.data?.filter((item) => {
            item.is_stay = false
            return true
        })
        columnList.value = response.data as GetArticleContributionListByUserRes
        console.log(response)
    } catch (err) {
        console.log(err)
    }
}

const jumpArticle = (articleID: number) => {

    router.push({ name: "ArticleShow", query: { articleID } })

}
onMounted(() => {
    init()
})
</script>
<style scoped lang="scss">
@import "./static/style/column.scss"
</style>