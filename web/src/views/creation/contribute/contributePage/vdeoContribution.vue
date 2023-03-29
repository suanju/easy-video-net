<template>
    <!-- 上传 -->
    <div class="vdeo-contribution ">
        <div class="upload-box animate__animated animate__bounceIn" v-show="!form.isShow">
            <el-upload class="upload" drag action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                multiple :on-success="handle.handleFileSuccess" :on-error="handle.handleFileError"
                :show-file-list="false" :before-upload="handle.beforeFileUpload" :auto-upload="true"
                :http-request="handle.RedefineUploadFile" accept=".mp4,.avi,.mov">
                <el-icon class="el-icon--upload">
                    <upload-filled />
                </el-icon>
                <div class="el-upload__text">
                    拖拽到此处也可上传
                </div>
                <div class="upload-btn">
                    <el-button type="primary">
                        Upload<el-icon class="el-icon--right">
                            <Upload />
                        </el-icon>
                    </el-button>
                </div>
                <div class="upload-audit">
                    当前审核队列 <el-tag class="tag" round>快速</el-tag>
                </div>
            </el-upload>
        </div>
        <!-- 配置上传信息 -->
        <div class="form-box animate__animated animate__bounceIn" v-show="form.isShow ">
            <!-- 视频预览 -->
            <div class="video-preview">
            <video class="video" ref="video" src=""  ></video>
            </div>
            <p>文件上传进度</p>
            <el-progress :text-inside="true" :stroke-width="16" :percentage="uploadFileformation.progress" />
            <h3> 基本设置</h3>
            <el-form :model="form" ref="ruleFormRef" label-width="120px" label-position="left"
                :rules="videoContributionRules" accept=".png,.jpg,.jpge">
                <el-form-item class="form-item-middle" label="封面">
                    <el-upload class="cover-uploader" :action="uploadCoveration.action" :show-file-list="false"
                        :on-success="handleCover.handleFileSuccess" :on-error="handleCover.handleFileError"
                        :before-upload="handleCover.beforeFileUpload" :auto-upload="true"
                        :http-request="handleCover.RedefineUploadFile">
                        <img v-if="uploadCoveration.FileUrl" :src="uploadCoveration.FileUrl" class="cover" />
                        <el-icon v-else class="cover-uploader-icon">
                            <Plus />
                        </el-icon>
                    </el-upload>
                </el-form-item>
                <el-form-item label="标题" prop="title">
                    <el-input v-model="form.title" placeholder="给视频起个标题吧~" />
                </el-form-item>
                <el-form-item label="类型">
                    <el-radio-group v-model="form.type">
                        <el-radio :label="false">自制</el-radio>
                        <el-radio :label="true">转载</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="定时发布" v-show="props.type != 'edit'">
                    <el-switch v-model="form.timing" />
                </el-form-item>
                <el-form-item label="选择时间" v-show="form.timing" class="animate__animated animate__fadeIn">
                    <el-col :span="7">
                        <el-date-picker v-model="form.date1time" type="datetime" placeholder="请选择定时发布时间" />
                    </el-col>
                </el-form-item>
                <el-form-item label="标签" class="label-box">
                    <el-tag v-for="tag in form.label" :key="tag" closable :disable-transitions="false"
                        class="label-item" @close="labelHandl.handleClose(tag)">
                        {{ tag }}
                    </el-tag>
                    <el-input v-if="form.labelInputVisible" ref="labelInputRef" v-model="form.labelText" size="small"
                        class="label-input" @keyup.enter="labelHandl.handleInputConfirm"
                        @blur="labelHandl.handleInputConfirm" />
                    <el-button class="label-btn" v-else size="small" @click="labelHandl.showInput">
                        + New Tag
                    </el-button>
                </el-form-item>
                <el-form-item label="介绍" class="form-item-middle" prop="introduce">
                    <el-input resize="none" maxlength="2000" rows="4" v-model="form.introduce" type="textarea"
                        placeholder="填写更全面的相关信息，让更多的人能找到你的视频吧" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary"
                        @click="useSaveData(form, uploadFileformation, uploadCoveration, ruleFormRef, router ,props)">{{props.type == "edit" ? "更新视频" : "发布视频"}}</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useVdeoContributionProp, useHandleFileMethod, useInit, useHandleCoverMethod, userLabelHandlMethod, useSaveData, useRules } from "@/logic/creation/contribute/contributePage/vdeoContribution"
import { UploadFilled, Upload } from '@element-plus/icons-vue'
import {  onMounted } from 'vue';
import { Plus } from '@element-plus/icons-vue'
import { ElInput} from 'element-plus'

const props = defineProps({
    type : {
        type : String
    }
})

const { form, uploadFileformation, uploadCoveration, labelInputRef, router, ruleFormRef ,video ,editVideoStore } = useVdeoContributionProp()
const handle = useHandleFileMethod(uploadFileformation, form ,video)
const handleCover = useHandleCoverMethod(uploadCoveration, form)
const labelHandl = userLabelHandlMethod(form, labelInputRef)
const { videoContributionRules } = useRules()


onMounted(() => {
    useInit(uploadFileformation, uploadCoveration ,form, props,editVideoStore)
})
</script>

<style scoped lang="scss">
@import "@/assets/styles/views/creation/contribute/contributePage/vdeoContribution.scss";
</style>
