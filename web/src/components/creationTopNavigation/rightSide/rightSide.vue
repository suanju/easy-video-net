<template>
  <div class="info">
    <div class="avatar">
      <router-link to="/personal">
        <el-avatar :size="36" :src="userInfo.userInfoData.photo" />
      </router-link>
    </div>
    <div class="tips-calendar-wrap">
     已经成为 UP {{ day }} 天
    </div>
    <div class="line-divid"></div>

    <SvgIcon name="message" class="icon left-icon"></SvgIcon>

    <SvgIcon name="download" class="icon"></SvgIcon>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router"
import { useUserStore } from "@/store/main";
import { onMounted, ref } from "vue";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime"
const userInfo = useUserStore();
const router = useRouter()
const day = ref(0)

//跳转
const jump = (name: string) => {
  router.push({
    name
  })
}

onMounted(() => {
  
  dayjs.extend(relativeTime)
  day.value = dayjs(Date()).diff(userInfo.userInfoData.created_at, 'day')
})

</script>

<style scoped lang="scss">
@import "../static/style/rightSide.scss";
</style>
