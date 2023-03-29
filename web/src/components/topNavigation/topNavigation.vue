<template>
  <el-row ref="elRowRef"
    :class="{ row: true, bgColorBlack: scrollTopNotTop && scroll, bgColorWhite: !scrollTopNotTop, boxShadow: boxShadow }"
    :style="{ maxWidth: screenWidth }">
    <el-col :span="1">
      <div class="grid-content ep-bg-purple" />
    </el-col>
    <el-col :span="6">
      <div class="grid-content ep-bg-purple-light" />
      <div class="menu-choose">
        <el-menu :default-active="globaStore.globalData.router.currentRouter" class="el-menu" mode="horizontal"
          @select="handleSelect">
          <el-menu-item :style="{ color: scrollTopNotTop ? color : '#18191C' }" index="/"> 首页 </el-menu-item>
          <el-menu-item :style="{ color: scrollTopNotTop ? color : '#18191C' }" index="/column"> 专栏 </el-menu-item>
          <el-menu-item :style="{ color: scrollTopNotTop ? color : '#18191C' }" index="/live"> 直播 </el-menu-item>
        </el-menu>
      </div>
    </el-col>
    <el-col :span="11">
      <div class="search" v-show="displaySearch">
        <el-input v-model="searchText" placeholder="最近新鲜事" size="large" @keyup.enter.native="searchChange">
          <template #suffix>
            <el-icon @click="search">
              <Search />
            </el-icon>
          </template>
        </el-input>
      </div>
    </el-col>
    <el-col :span="6">
      <RightSide :icon-color="scrollTopNotTop ? color : '#18191C'"></RightSide>
    </el-col>
    <el-col :span="1">
      <div class="grid-content ep-bg-purple" />
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { useGlobalStore } from "@/store/main";
import { Search } from "@element-plus/icons-vue";
import { onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";
import RightSide from "./rightSide/rightSide.vue";

const emit = defineProps({
  scroll: {
    type: Boolean,
    required: false,
    default: false,
  },
  //是否固定颜色
  color: {
    type: String,
    required: false,
    default: '#18191C',
  },
  displaySearch: {
    type: Boolean,
    required: false,
    default: true,
  },
  boxShadow: {
    type: Boolean,
    required: false,
    default: false,
  },
})
components: {
  RightSide;
}


const globaStore = useGlobalStore()
const router = useRouter();
const searchText = ref("");
const scrollTopNotTop = ref(true) // 滚动条是否在顶部
const elRowRef = ref()
//屏幕最大宽
const screenWidth = screen.width + 'px'
//选项卡点击
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
  router.push(key);
};

const search = () => {
  if (!searchText.value) return false;
  router.push({
    name: 'VideoSearch', params: {
      text: searchText.value
    }
  })
}

const searchChange = () => {
  if (!searchText.value) return false;
  router.push({
    name: 'VideoSearch', params: {
      text: searchText.value
    }
  })
}

onMounted(() => {
  // 监听滚动条位置
  if (emit.scroll) {
    watch(() => { return globaStore.globalData }, async () => {
      if (globaStore.globalData.scroll.scrollTop > 58) {
        scrollTopNotTop.value = false
      } else {
        scrollTopNotTop.value = true
      }
    }, { immediate: true, deep: true })
  }
})

</script>

<style scoped lang="scss">
@import "./static/style/topNavigation.scss";
</style>
