<template>
    <el-card shadow="never">
        <div class="user-card">
            <div class="card-left">
                <div class="avatar" @click="jumpSpace(id)"> <el-avatar :size="104" :src="avatar" />
                </div>
            </div>

            <div class="card-right">
                <div class="username">
                    <span>{{ username }}</span>
                    <div class="private-letter" @click="usePersonalLetter(id)">
                        <SvgIcon name="message" class="icon" color="#9499A0"></SvgIcon> 私信
                    </div>
                </div>
                <div class="signature">
                    <VueEllipsis3 :text="signature ? signature : '这个人很勤快什么都没留下~'">
                        <template v-slot:ellipsisNode>
                            <span>...</span>
                        </template>
                    </VueEllipsis3>
                </div>
                <div class="btn-list">
                    <el-button v-removeFocus type="primary" size="small" round :icon="House"
                        @click="jumpSpace(id)">主页</el-button>
                    <el-button class="attention" v-if="!is_attention" v-removeFocus type="primary" size="small" round
                        :icon="Plus" @click="attention(id)">关注</el-button>
                    <el-button class="attention" v-if="is_attention" v-removeFocus type="primary" size="small" round
                        :icon="MoreFilled" color="#F1F2F3" @click="attention(id)">已关注</el-button>
                </div>
            </div>
        </div>
    </el-card>
</template>

<script lang="ts" setup>
import useAttention from "@/hooks/useAttention";
import usePersonalLetter from "@/hooks/usePersonalLetter";
import { House, MoreFilled, Plus } from '@element-plus/icons-vue';
import { VueEllipsis3 } from 'vue-ellipsis-3';
import { useRouter } from "vue-router";

components: {
    VueEllipsis3
}
const props = defineProps({
    id: {
        type: Number,
        required: true,
    },
    avatar: {
        type: String,
        required: true,
    },
    username: {
        type: String,
        required: true,
    },
    signature: {
        type: String,
        required: true,
    },
    is_attention: {
        type: Boolean,
        required: true,
    }
})
const emits = defineEmits(["attention"])

const router = useRouter()


//跳转用户空间
const jumpSpace = (id: number) => {
    router.push({ name: "SpaceIndividual", params: { "id": id } })
}

const attention = async (id: number) => {
    if (await useAttention(id)) {
        emits("attention", id)
    }
}

</script>

<style lang="scss" scoped>
@import "./static/style/card.scss";
</style>