<template>
  <div :class="{ 'page-header': true, 'animate__animated': animate, 'animate__slideInDown': animate, 'animate__faster': animate }">
    <SvgIcon :name="iconNmae" class="icon"></SvgIcon>
    <span>{{ title }}</span>
    <div v-if="whiteWhale" class="whiteWhale">
    </div>
  </div>
</template>


<script setup lang="ts">
defineProps({
  title: {
    type: String,
    required: true,
  },
  iconNmae: {
    type: String,
    required: true,
  },
  animate: {
    type: Boolean,
    default: true
  },
  whiteWhale :
  {
    type: Boolean,
    default: true
  },
})
</script>
<style scoped lang="scss">
@import "./static/style/pageHeader.scss";
</style> 