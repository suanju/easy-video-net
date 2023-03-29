<template>
    <div class="video-card" @click="jump()">
        <div class="card-left" >
            <div class="item-image" :style="{ backgroundImage: `url(${props.cover})` }"></div>
            <div class="classification">
                <div class="classification-right">
                    <span>{{ formattingSecondTime(props.video_duration) }}</span>
                </div>


            </div>
        </div>
        <div class="card-right">
            <div class="video-title"><VueEllipsis3 :text="props.title" :visibleLine="2" /></div>
            <div class="video-information-top">
                <SvgIcon name="up" class="icon" color="#999"></SvgIcon>
                <span class="video-information-username">{{ props.username }} </span> 
            </div>
            <div class="video-information-bottom">
                <SvgIcon name="video" class="icon" color="#999"></SvgIcon>
                <span class="video-information-plays-num">{{ props.heat }} </span> 
                <SvgIcon name="barrageNumber" class="icon" color="#999"></SvgIcon>
                <span class="video-information-username">{{ props.barrageNumber }} </span> 
            </div>
        </div>
    </div>

</template>

<script setup lang="ts">
import { VueEllipsis3 } from 'vue-ellipsis-3';
import { ref, defineProps } from "vue"
import { formattingSecondTime } from "@/utils/conversion/timeConversion"
import { useRouter } from "vue-router"

components: {
    VueEllipsis3
}
const props = defineProps({
    id: {
        type: Number,
        required: true,
        default : 0
    },
    video_duration: {
        type: Number,
        required: true,
    },
    title: {
        type: String,
        required: true,
    },
    cover: {
        type: String,
        required: true,
    },
    heat: {
        type: Number,
        required: true,
    },
    barrageNumber: {
        type: Number,
        required: true,
    },
    username: {
        type: String,
        required: true,
    },
    created_at: {
        type: String,
        required: true,
    },
}
)

const router = useRouter()
const isMouseover = ref(false)

const jump = () => {
    router.push({ name: "VideoShow", query: { videoID: props.id } }) 
}

</script>

<style scoped lang="scss">
@import "./static/style/videoPageVideoCard.scss";
</style>