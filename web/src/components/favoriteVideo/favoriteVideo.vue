<template>
    <div class="favorite-video">
        <div class="checkbox" v-loading="isLoading">
            <div class="check-item" v-for="item in favoritesList" :key="item.id">
                <el-checkbox v-model="item.choose" size="large" :disabled="item.selected && item.present < item.max">
                    <div class="title">{{ item.title }}</div>
                    <div class="num"> {{ item.present }}/{{ item.max }}</div>
                </el-checkbox>
            </div>
        </div>
        <div class="create">
            <el-input v-model="createInput" placeholder="创建收藏夹" :prefix-icon="Plus">
                <template #append>
                    <el-button type="primary" round @click="createFavorite()">创建</el-button>
                </template>
            </el-input>
        </div>
        <div class="function">
            <el-button type="primary" v-removeFocus round @click="confirmedCollection">确认收藏</el-button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { createFavorites, favoriteVideo, getFavoritesListByFavoriteVideo } from '@/apis/personal';
import globalScss from "@/assets/styles/global/export.module.scss";
import { createFavoritesReq } from '@/types/personal/collect/createFavorites';
import { FavoriteVideoReq, GetFavoritesListByFavoriteVideoItem, GetFavoritesListByFavoriteVideoReq, getFavoritesListByFavoriteVideoRes, GetFavoritesListRes } from '@/types/personal/collect/favorites';
import { vRemoveFocus } from "@/utils/customInstruction/focus";
import { Plus } from '@element-plus/icons-vue';
import Swal from 'sweetalert2';
import { onMounted, Ref, ref } from 'vue';

const props = defineProps({
    id: {
        type: Number,
        required: true,
    }
})

const emits = defineEmits(["shutDown", "success"])

const createInput = ref("")
const favoritesList = ref(<getFavoritesListByFavoriteVideoRes>{})
const isLoading = ref(true)
const isChoneIds = ref(<Array<number>>[])

const loadData = async (favoritesList: Ref<GetFavoritesListRes>, isLoading: Ref<boolean>) => {
    try {
        //获取收藏夹列表 
        isLoading.value = true
        const response = await getFavoritesListByFavoriteVideo(<GetFavoritesListByFavoriteVideoReq>{
            video_id: props.id
        });
        if (!response.data) return false
        response.data = response.data.filter((item: GetFavoritesListByFavoriteVideoItem) => {
            if (item.selected) {
                isChoneIds.value.push(item.id)
                item.choose = true
            } else {
                item.choose = false
            }
            return true
        })
        favoritesList.value = response.data
        isLoading.value = false
    } catch (err) {
        console.log(err)
    }
}

//创建收藏夹
const createFavorite = async () => {
    try {
        await createFavorites(<createFavoritesReq>{
            title: createInput.value
        })
        loadData(favoritesList, isLoading)
        createInput.value = ""
    } catch (err: any) {
        Swal.fire({
            title: "创建失败",
            confirmButtonColor: globalScss.colorButtonTheme,
            heightAuto: false,
            icon: "error",
        })
    }
}

//确认收藏
const confirmedCollection = async () => {
    try {
        //得到确认收藏的收藏夹
        let ids: Array<number> = []
        favoritesList.value.filter((item) => {
            if (item.choose) ids.push(item.id);
        })
        if (JSON.stringify(isChoneIds.value) === JSON.stringify(ids)) {
            throw new Error("未选择收藏夹");
        }
        console.log(ids)
        await favoriteVideo(<FavoriteVideoReq>{
            ids,
            video_id: props.id
        })
        emits("success")
        emits("shutDown")
        Swal.fire({
            title: "收藏成功",
            confirmButtonColor: globalScss.colorButtonTheme,
            heightAuto: false,
            icon: "success",
        })
        loadData(favoritesList, isLoading)
    } catch (err: any) {
        emits("shutDown")
        Swal.fire({
            title: err.message,
            confirmButtonColor: globalScss.colorButtonTheme,
            heightAuto: false,
            icon: "error",
        })
    }
}

onMounted(() => {
    loadData(favoritesList, isLoading)
})

</script>

<style lang="scss" scoped>
@import "./static/style/favoriteVideo.scss";
</style>