<template>
    <div class="card-list">
        <div class="author-card">
            <div class="avatar"><el-avatar :size="100"
                    :src="authorInfo.avatar" /></div>
            <div class="author-name"><span>{{ authorInfo.username }}</span></div>
            <div class="column-info">
                <div class="column-info-item">
                    <div> <span>文章</span></div>
                    <div class="info-num"><span>{{ totalInfo.article_num }}</span></div>
                </div>
                <div class="column-info-item">
                    <div><span>分类</span></div>
                    <div class="info-num"><span>{{ totalInfo.classification_num }}</span></div>
                </div>
            </div>
            <a class="fo-github" @click="jumpGithub()">
                <SvgIcon name="github" class="icon" color="#fff" /> Follow Me
            </a>
        </div>
        <div class="classification-card">
            <div class="card-title"><SvgIcon name="file" class="icon" color="#4C4948"></SvgIcon>分类</div>
            <el-tree :data="totalInfo.classification" :default-expand-all="true" :props="defaultProps" />
        </div>
    </div>
</template>

<script lang="ts" setup>

import { getArticleTotalInfo } from '@/apis/contribution';
import { getArticleTotalInfoRes } from '@/types/creation/contribute/contributePage/articleContribution';
import { onMounted, reactive } from 'vue';

const authorInfo = reactive({
    avatar : "http://q1.qlogo.cn/g?b=qq&nk=2506152074&s=100",
    username : "橡皮擦",
    github : "https://github.com/suanju",
}) 
const totalInfo = reactive(<getArticleTotalInfoRes>{})

const defaultProps = {
  children: 'children',
  label: 'label',
}

const init = async () => {
    const articleTotalInfo = await getArticleTotalInfo()
    if(!articleTotalInfo.data) return false
    totalInfo.classification = articleTotalInfo.data?.classification
    totalInfo.article_num = articleTotalInfo.data?.article_num
    totalInfo.classification_num = articleTotalInfo.data?.classification_num
}

const jumpGithub = () =>{
    window.open(authorInfo.github); 
}

onMounted(() => {
    init()
   
})
</script>

<style lang="scss" scoped>
@import "./static/style/sideCard.scss";
</style>