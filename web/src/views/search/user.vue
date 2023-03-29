<template>
    <div class="main">
        <div class="middle">
            <div class="video-card">
                <Card v-for="(item, index) in searchInfo" :id="item.id" :avatar="item.photo" :username="item.username"
                    :signature="item.signature" :is_attention="item.is_attention" @attention="attention" />
            </div>
        </div>
        <div class="empty" v-show="isLoading && !searchInfo.length">
            <el-empty description="这里空空如也" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import { search } from "@/apis/commonality";
import Card from "@/components/userSearchCard/card.vue";
import { SearchReq, SearchRes, UserInfo, UserInfoList } from "@/types/commonality/commonality";
import { reactive, Ref, ref, UnwrapNestedRefs, watch } from "vue";
import { useRoute } from "vue-router";
component: {
    Card
}

const isLoading = ref(false)
const route = useRoute()
const pageInfo = reactive(<SearchReq>{
    //待完成分页功能
    page_info: {
        page: 1,
        size: 15,
        keyword: route.params.text
    },
    type: "user"
})

let searchInfo = ref<SearchRes>([])

const loadData = async (searchInfo: Ref<SearchRes>, pageInfo: UnwrapNestedRefs<SearchReq>) => {
    const response = await search(pageInfo)
    if (!response.data) return
    response.data as UserInfoList
    searchInfo.value = [...searchInfo.value, ...response.data]
    //请求成功后下次分页+1
    pageInfo.page_info.page = pageInfo.page_info.page + 1
    isLoading.value = true
}

//子组件关注成功回调
const attention = (id: number) => {
    searchInfo.value = searchInfo.value.filter((item: UserInfo) => {
        if (item.id == id) {
            item.is_attention = !item.is_attention
        }
        return item
    })
}


watch(() => route.path, async () => {
    pageInfo.page_info.page = 1
    pageInfo.page_info.keyword = route.params.text as string
    searchInfo.value = []
    isLoading.value = false
    loadData(searchInfo, pageInfo)
    console.log(route)
}, { immediate: true, deep: true })
</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/search/user.scss";
</style>