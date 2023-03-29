<template>
    <div :class="{ item: true }" @click="jump()">
        <div :class="{ normal: true, mouseleave: !isMouseover }" v-show="!isMouseover">
            <div class="head">
                <div class="item-image" :style="{ backgroundImage: `url(${props.img})` }"></div>
                <div class="classification">
                 
                </div>
            </div>
            <div class="info">
                <div class="video-title">
                    <VueEllipsis3 :text="props.title" :visibleLine="1" />
                </div>
                <div class="video-information">
                    <div class="information-left">
                        <el-avatar :size="34" :src="props.photo" />
                        <span class="video-information-username">{{ props.username }} </span>
                    </div>
                    <span class="people"></span><span class="people-span">{{ props.online }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

import { ref, defineProps } from "vue"
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
    title: {
        type: String,
        required: true,
    },
    img: {
        type: String,
        required: true,
    },
    online: {
        type: Number,
        required: true,
    },
    photo: {
        type: String,
        required: true,
    },
    username: {
        type: String,
        required: true,
    }
}
)

const router = useRouter()
const isMouseover = ref(false)

const jump = () => {
    router.push({ name: "liveRoom", query: { roomID: props.id } })
}

</script>

<style scoped lang="scss">
@import "./static/style/card.scss";
</style>