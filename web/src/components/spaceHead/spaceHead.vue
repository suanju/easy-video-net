<template>
    <div class="header-content">
        <div class="userInfo">
            <div class="info">
                <div class="info-avatar">
                    <el-avatar :size="64" :src="userInfo.photo" />
                </div>
                <div class="into-text">
                    <div class="name"> {{ userInfo.username }}</div>
                    <div class="signature"> {{ userInfo.signature ? userInfo.signature : "这个人很勤快什么都没留下~" }}</div>
                </div>
            </div>
            <div class="function">
                <el-button class="attention" v-if="!userInfo.is_attention" v-removeFocus type="primary" size="small" round
                    :icon="Plus" @click="attention()">关注</el-button>
                <el-button class="attention" v-if="userInfo.is_attention" v-removeFocus type="primary" size="small" round
                    :icon="MoreFilled" color="#F1F2F3" @click="attention()">已关注</el-button>
                <div class="button"> <el-button type="primary" size="small" @click="usePersonalLetter(userInfo?.id)" round>
                        <SvgIcon name="message" class="icon" color="#fff"></SvgIcon>私信
                    </el-button></div>
            </div>
        </div>

        <div class="menu">
            <div class="left">
                <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                    <el-menu-item index="SpaceIndividual">Ta的作品</el-menu-item>
                    <el-menu-item index="MyAttention">关注</el-menu-item>
                    <el-menu-item index="MyVermicelli">粉丝</el-menu-item>
                </el-menu>
            </div>
            <div class="right">
                <div class="item-box">
                    <div>关注 : {{ userInfo.attention_num }}</div>
                    <div class="data-item">粉丝 : {{ userInfo.vermicelli_num }}</div>
                </div>
            </div>
        </div>
        <div class="dividing-line">

        </div>
    </div>
</template> 

<script lang="ts" setup>
import { getSpaceIndividual } from '@/apis/space';
import useAttention from '@/hooks/useAttention';
import usePersonalLetter from "@/hooks/usePersonalLetter";
import { useSpaceStore } from '@/store/main';
import { GetSpaceIndividualReq, GetSpaceIndividualRes } from '@/types/space/space';
import { vRemoveFocus } from "@/utils/customInstruction/focus";
import { MoreFilled, Plus } from '@element-plus/icons-vue';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const activeIndex = ref('1')

const space = useSpaceStore()
const userInfo = ref(<GetSpaceIndividualRes>{})
const router = useRouter();


const handleSelect = (key: string) => {
    router.push({ name: key })

}


const attention = async () => {
    if (await useAttention(space.spaceInfoData.id)) {
        userInfo.value.is_attention = !userInfo.value.is_attention
    }
}

const loadData = async () => {
    try {
        const response = await getSpaceIndividual(<GetSpaceIndividualReq>{
            id: space.spaceInfoData.id
        })
        if (!response.data) return
        userInfo.value = response.data


    } catch (err: any) {
        console.log(err)
    }
}


onMounted(() => {
    loadData()
})
</script>

<style lang="scss" scoped>
@import "./static/style/spaceHead.scss";
</style>