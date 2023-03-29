<template>

    <div class="video-introduce">
        <div class="introduce">
            <VueEllipsis3 :text="introduce ? introduce : ''" :visibleLine="introduceVisibleLine" reflowOnResize>
                <template v-slot:ellipsisNode>
                    <span class="show-more" @click="introduceVisibleLine = 999">... 展开</span>
                </template>
            </VueEllipsis3>
        </div>
        <div class="tag" >
            <el-tag v-for="item in label" :key="item" class="tag-item" type="info" effect="dark" round>
                {{ item }}</el-tag>
        </div>
    </div>

</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { VueEllipsis3 } from 'vue-ellipsis-3';

const props = defineProps({
    introduce: {
        type: String,
        required: false,
    },
    label: {
        type: Array,
        required: false,
    }

})
components: {
    VueEllipsis3
}

const introduceVisibleLine = ref(3)
</script>

<style lang="scss" scoped>
@import "./static/css/videoIntroduce.scss";
</style>