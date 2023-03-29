<template>
  <svg aria-hidden="true" :style="{ color: color }">
    <use :href="symbolId" :fill="color" />
    <slot></slot>
  </svg>
</template>
  
<script  lang="ts">
import { computed, defineComponent } from 'vue';

export default defineComponent({
  name: 'SvgIcon',
  props: {
    prefix: {
      type: String,
      default: 'icon',
    },
    name: {
      type: String,
      required: true,
    },
    color: {
      type: String,
      default: '#333',
    },
  },
  setup(props) {
    const symbolId = computed(() => `#${props.prefix}-${props.name}`)
    return { symbolId }
  },
})
</script>