<template>
    <div class="live-info">
        <div class="avatar">
            <el-avatar :size="38" :src="src" />
        </div>
        <div class="info">
            <div class="info-left">
                <p class="live-titel"> {{ titel }} </p>
            </div>
            <div class="info-right">
                <el-button type="primary" v-removeFocus :icon="Star" circle />
                <el-button type="primary" v-removeFocus :icon="Warning" circle />
                <el-button type="primary"  v-removeFocus :icon="Position" circle />
            </div>
        </div>
    </div>
</template>


<script setup  lang="ts">
import {vRemoveFocus} from "@/utils/customInstruction/focus"
import {
  Position,
  Warning,
  Star,
} from '@element-plus/icons-vue'

defineProps({
    src: {
        type: String,
        required: true,
    },
    titel: {
        type: String,
        required: true,
    },
 }
)

</script>


<style scoped lang="scss">
@import "./static/style/liveHeader.scss"
</style>