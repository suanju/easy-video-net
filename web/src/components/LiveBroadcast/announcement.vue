<template>
    <div class="announcement-box">
        <div  :class="{ mouseover : isMouseover ,mouseleave: !isMouseover}"  class="announcement shadow-box "   @mouseover="mouseOver" @mouseleave="mouseleave">
            <div class="img-box">
                <img src="@/assets/img/temporary/00018.jpg" />
            </div>
            <div class="info-box">
                <h3>枉有一壶醉人酒 难解半点俗世愁</h3>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const isMouseover = ref(false)

const mouseOver = () => {
    isMouseover.value = true
}
const mouseleave = () => {
    isMouseover.value = false
}
</script>

<style scoped lang="scss">
.announcement-box {

    .shadow-box {
        box-shadow: 0 1px 20px -6px rgba(0, 0, 0, 0.5);
        transition: all .3s ease;
    }

    .mouseleave {
        box-shadow: 0 1px 20px -6px rgba(0, 0, 0, 0.5);
        transition: all .3s ease;
    }

    .mouseover {
        box-shadow: 0 1px 20px -6px rgba(0, 0, 0, 1);
        transition: all .3s ease;
    }

    .announcement {
        width: 100%;
        height: 256px;
        border-radius: 10px;
        -webkit-animation: zoomIn .8s ease-in-out;
        animation: zoomIn .8s ease-in-out;

        .img-box {
            img {
                border-radius: 10px 10px 0px 0px;
                width: 100%;
                max-height: 40%;
            }

        }

        .info-box {
            display: flex;
            justify-content: center;

        }
    }



}
</style>