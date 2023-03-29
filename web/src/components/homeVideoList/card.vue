<template>
    <div :class="{ item: true }" @click="jump()">
        <div :class="{ normal: true, mouseleave: !isMouseover }" v-show="!isMouseover">
            <div class="head">
                <div class="item-image" :style="{backgroundImage : `url(${props.cover})` }"></div>
                <div class="classification">
                    <div class="classification-left">
                        <span class="play-volume">
                            <SvgIcon name="video" class="icon" color="#fff"></SvgIcon>
                            <span>{{ props.heat }}</span>
                        </span>
                        <span class="barrage-number">
                            <SvgIcon name="barrageNumber" class="icon" color="#fff"></SvgIcon>
                            <span>{{ props.barrageNumber }}</span>
                        </span>
                    </div>
                    <div class="classification-right">
                        <span>{{ formattingSecondTime(props.video_duration) }}</span>
                    </div>
                </div>
            </div>
            <div class="info">
                <div class="video-title"><VueEllipsis3 :text="props.title" :visibleLine="2" /></div>
                <div class="video-information">
                    <SvgIcon name="up" class="icon" color="#999"></SvgIcon>
                    <span class="video-information-username">{{ props.username }} </span> <span class="video-information-time">·
                        {{ timestampFormat(props.created_at) }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

import { ref, defineProps } from "vue"
import { formattingSecondTime, timestampFormat } from "@/utils/conversion/timeConversion"
import { useRouter } from "vue-router"
import { VueEllipsis3 } from 'vue-ellipsis-3';


components: {
    VueEllipsis3
}


const props = defineProps({
    id: {
        type: Number,
        required: true,
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
@import "./static/style/card.scss";
</style>