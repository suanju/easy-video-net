<template>
    <div class="title"><span >粉丝列表</span></div>
    <div class="list">
        <div class="border"></div>
        <div class="item" v-for="item in list" :key="item.id">
            <div class="avatar"> <el-avatar :size="52" :src="item.photo" />
            </div>
            <div class="info">
                <div class="username"> {{ item.name }}</div>
                <div class="signature"> {{ item.signature }}</div>
            </div>
            <div class="function">
                <el-button class="attention" v-if="!item.is_attention" v-removeFocus type="primary" size="small" round
                    :icon="Plus" @click="attentionUser(item.id)">关注</el-button>
                <el-button class="attention" v-if="item.is_attention" v-removeFocus type="primary" size="small" round
                    :icon="MoreFilled" color="#F1F2F3" @click="attentionUser(item.id)">已关注</el-button>
            </div>
        </div>
</div>
</template>

<script lang="ts" setup>
import { attention } from '@/apis/personal';
import { useSpaceStore, useUserStore } from '@/store/main';
import globalScss from "@/assets/styles/global/export.module.scss"
import { AttentionReq } from '@/types/personal/userInfo/userInfo';
import { Plus, MoreFilled } from '@element-plus/icons-vue'
import Swal from 'sweetalert2';
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { getVermicelliList } from '@/apis/space';
import { GetVermicelliListReq, GetVermicelliListRes } from '@/types/space/space';
import { onMounted, ref } from 'vue';

const userStore = useUserStore()
const space = useSpaceStore()
const list = ref(<GetVermicelliListRes>[])

const attentionUser = async (id: number) => {
    try {
        if (userStore.userInfoData.id == id) {
            Swal.fire({
                title: "不能对自己操作哟!",
                heightAuto: false,
                confirmButtonColor: globalScss.colorButtonTheme,
                icon: "error",
            })
            return
        }
        await attention(<AttentionReq>{
            uid: id
        })

        list.value = list.value.filter((item) =>{
            if(item.id == id){
                item.is_attention = !item.is_attention
            } 
            return item
        })
        
    } catch (err: any) {
        Swal.fire({
            title: "操作失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}

const loadData = async () => {
    try {
        const response = await getVermicelliList(<GetVermicelliListReq>{
            id: space.spaceInfoData.id
        })
        if (!response.data) return false
        list.value = response.data
        console.log(response)
    } catch (err: any) {
        console.error(err)
    }
}


onMounted(() => {
    loadData()
})
</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/space/myAttention.scss";
</style>