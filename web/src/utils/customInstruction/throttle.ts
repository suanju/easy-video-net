import { DirectiveBinding } from 'vue'
//节流
export const vThrottle = {
    mounted(el: HTMLElement, binding: DirectiveBinding<any>) {
        let lastTime: number = 0
        el.addEventListener('click', () => {
            const nowTime = new Date().getTime()
            const remainTime = nowTime - lastTime
            if (remainTime - binding.value.time >= 0) {
                binding.value.fun(...binding.value.params)
                lastTime = nowTime
            }
        })
    }
}
