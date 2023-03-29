import { App, DirectiveBinding } from 'vue'
//防抖
export const  vAntiShake = {
    mounted(el: HTMLElement, binding: DirectiveBinding<any>) {
      let timer: NodeJS.Timeout | null = null
      el.addEventListener('click', () => {
      let firstClick: Boolean = !timer;
      if (firstClick) {
        binding.value.fun(...binding.value.params)
      }
      if (timer) {
        clearTimeout(timer)
      }
      timer = setTimeout(() => {
        timer = null
        if (!firstClick) {
          binding.value.fun(...binding.value.params)
        }
      },  binding.value.time);
      })
    }
}

