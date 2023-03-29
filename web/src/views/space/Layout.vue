<template>
    <div class="space">
        <div class="space-layout animate__animated animate__rotateInUpLeft animate__faster">
            <el-container>
                <div style="width:100% ;height:100%">
                    <spaceHead></spaceHead>
                    <div ref="scrollRef" class="scroll" :style="{width:'100%', height : scrollHeight + 'px'}" >
                        <router-view></router-view>
                    </div>
                </div>
            </el-container>
        </div>
</div>
</template>

<script setup lang="ts">

import personalNavigation from "@/components/personalNavigation/personalNavigation.vue"
import spaceHead from "@/components/spaceHead/spaceHead.vue"
import { onBeforeMount, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { useSpaceStore } from "@/store/main"
components: {
    personalNavigation
    spaceHead
}

const route = useRoute()
const uid = ref(0)
const space = useSpaceStore()
const scrollHeight = ref(0)
const scrollRef = ref()

//需要在组件挂载前绑定id
onBeforeMount(() => {
    uid.value = parseInt(route.params.id as string)
    space.setSpaceID(uid.value)
})

onMounted(() =>{
    scrollHeight.value = document.documentElement.clientHeight -scrollRef.value.offsetTop 
})

</script>

<style scoped lang="scss">
@import "@/assets/styles/views/personal/Layout.scss";

.scroll{
    overflow: scroll;
}
</style>

