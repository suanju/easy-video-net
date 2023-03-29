<template>
    <div class="overall">
        <div class=" principal  personal-layout">
            <div class="form-box">
                <el-upload class="cover-uploader" :action="createFavoriteRmationForm.action" :show-file-list="false"
                    :on-success="handle.handleFileSuccess" :on-error="handle.handleFileError"
                    :before-upload="handle.beforeFileUpload" :auto-upload="true" :http-request="handle.RedefineUploadFile"
                    accept=".png,.jpg,.jpeg">
                    <img v-if="createFavoriteRmationForm.FileUrl" :src="createFavoriteRmationForm.FileUrl" class="cover" />
                    <el-icon v-else class="cover-uploader-icon">
                        <Plus />
                    </el-icon>
                </el-upload>
                <div>
                    <div class="form-show">
                        <el-form :model="createFavoriteRmationForm" ref="saveDateFormRef" :rules="liveInformationRules"
                            label-position="left" label-width="5rem">
                            <el-form-item label="标题" prop="title">
                                <el-input v-model="createFavoriteRmationForm.title" />
                            </el-form-item>
                            <el-form-item label="介绍">
                                <el-input type="textarea" resize="none" :rows="4"
                                    v-model="createFavoriteRmationForm.content" />
                            </el-form-item>
                        </el-form>
                    </div>
                </div>
            </div>
            <div class="bottom-box">
                <span class="text"> 请设置您的封面和标题,以便更好查询收藏内容嗷~ </span>
                <div class="button">
                    <el-button v-removeFocus
                        @click="useSaveData(createFavoriteRmationForm, saveDateFormRef, rawData, router, emits)"
                        type="primary" round> {{createFavoriteRmationForm.id == 0 ? "创建文件夹" : "更新文件夹"}}
                    </el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>

import { Plus } from "@element-plus/icons-vue";
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { useHandleFileMethod, useSaveData, useRules, useInit, useCreateFavoritesProp } from "@/logic/personal/create/createFavorites"
import { onMounted, onUpdated } from "vue";
import { GetFavoritesListItem } from "@/types/personal/collect/favorites";

const props = defineProps({
    info: {
        type: Object as () => GetFavoritesListItem,
    },
    type : {
        type : Boolean, //true 为插入模式false 为更新模式
        required: true,
    }
})
const emits = defineEmits(["shutDown"])

const { createFavoriteRmationForm, saveDateFormRef, rawData, router } = useCreateFavoritesProp()
const handle = useHandleFileMethod(createFavoriteRmationForm)
const { liveInformationRules } = useRules();

onMounted(() => {
    useInit(createFavoriteRmationForm, rawData,props.type,props.info)

})

//数据变化再次更新
onUpdated(() => {
    useInit(createFavoriteRmationForm, rawData,props.type,props.info)
})
</script>

<style lang="scss" scoped>@import "./static/style/createFavorites.scss";</style>
