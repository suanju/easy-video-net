<template>
    <div class="carousel">
        <el-carousel ref="carouselRef" arrow="never" height="490px" :autoplay="false" @change="change">
            <el-carousel-item v-for="item in rotograph" :key="item.cover">
                <div class="carousel-item" @click="jump(item)">
                    <div class="carousel-img-box">
                        <el-image class="carousel-img" :src="item.cover" fit="cover" />
                    </div>
                </div>
            </el-carousel-item>
            <div class="gradient" :style="{ background: `linear-gradient(rgba(0,0,0,0),${color})` }"></div>
            <!-- 轮播图底部 -->
            <div class="carousel-bottom" :style="{ backgroundColor: `${color}` }">
                <div class="carousel-title">{{ carouselTitle ? carouselTitle : props.rotograph[0]?.title }}</div>
                <div class="toggle-button">
                    <div class="button-item" v-throttle="{ fun: carouselSwitch, params: [false], time: 500 }">
                        <SvgIcon name="rightArrow" class="arrow-icon rotation" color="#fff">
                        </SvgIcon>
                    </div>
                    <div class="button-item" v-throttle="{ fun: carouselSwitch, params: [true], time: 500 }">
                        <SvgIcon name="leftArrow" class="arrow-icon" color="#fff"></SvgIcon>
                    </div>
                </div>
            </div>
        </el-carousel>
    </div>
</template>

<script lang="ts" setup>
import { Rotograph, RotographList } from "@/types/home/home";
import { vThrottle } from "@/utils/customInstruction/throttle";
import { ElCarousel, ElCarouselItem, ElImage } from 'element-plus';
import { onMounted, ref } from 'vue';
import { useRouter } from "vue-router";

const props = defineProps({
    rotograph: {
        type: Array as () => RotographList,
        required: true,
        default: () => []
    }
})
const router = useRouter()
const carouselRef = ref()
const carouselTitle = ref("")
const color = ref("")
//切换轮播图 true 下一张 false 上一张
const carouselSwitch = (is: boolean) => {
    if (is) {
        carouselRef.value.next()

    } else {
        carouselRef.value.prev()
    }
}


//轮播图变化事件
const change = (index: number) => {
    carouselTitle.value = props.rotograph[index].title
    color.value = props.rotograph[index].color
}

const jump = (item: Rotograph) => {
    switch (item.type) {
        case "video":
            router.push({ name: "VideoShow", params: { id: item.to_id } })
            break
        case "article":
            router.push({ name: "ArticleShow", params: { id: item.to_id } })
            break
        case "live":
            router.push({ name: "liveRoom", params: { id: item.to_id } })
            break
    }
}

onMounted(() => {
    color.value = props.rotograph[0].color
})
</script>

<style lang="scss" scoped>
@import "./static/css/homeRotograph.scss";
</style>