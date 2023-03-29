<template>
    <div class="article-contribution">
        <!-- 基本设置 -->
        <div class="form-box animate__animated animate__bounceInRight" v-show="form.isShow">
            <h3> 基本设置</h3>
            <el-form :model="form" ref="ruleFormRef" label-width="120px" label-position="left"
                :rules="rules.articleContributionRules">
                <el-form-item label="标题" prop="title">
                    <el-input v-model="form.title" placeholder="给文章起个标题吧~" />
                </el-form-item>
                <el-form-item class="form-item-middle" label="封面">
                    <el-upload class="cover-uploader" :action="uploadCoveration.action" :show-file-list="false"
                        :on-success="handleCover.handleFileSuccess" :on-error="handleCover.handleFileError"
                        :before-upload="handleCover.beforeFileUpload" :auto-upload="true"
                        :http-request="handleCover.RedefineUploadFile" accept=".png,.jpg,.jpeg">
                        <img v-if="uploadCoveration.FileUrl" :src="uploadCoveration.FileUrl" class="cover" />
                        <el-icon v-else class="cover-uploader-icon">
                            <Plus />
                        </el-icon>
                    </el-upload>
                </el-form-item>
                <el-form-item label="内容编辑">
                    <el-button v-removeFocus size="small" type="primary" :icon="Edit" round
                        @click="form.isShow = false">编辑</el-button>
                </el-form-item>
                <el-form-item label="定时发布" v-show="props.type != 'edit'">
                    <el-switch v-model="form.timing" />
                </el-form-item>
                <el-form-item label="选择时间" v-show="form.timing" class="animate__animated animate__fadeIn">
                    <el-col :span="7">
                        <el-date-picker v-model="form.date1time" type="datetime" placeholder="请选择定时发布时间" />
                    </el-col>
                </el-form-item>
                <el-form-item label="开启评论">
                    <el-switch v-model="form.comments" />
                </el-form-item>
                <el-form-item label="分类">
                    <el-cascader v-model="cascaderValue" placeholder="请选择分类" :props="{
                        value: 'id',
                        label: 'label',
                        children: 'children'
                    }" :options="cascader" :show-all-levels="false" @change="cascaderHandleChange" />
                </el-form-item>
                <el-form-item label="标签" class="label-box">
                    <el-tag v-for="tag in form.label" :key="tag" closable :disable-transitions="false" class="label-item"
                        @close="labelHandl.handleClose(tag)">
                        {{ tag }}
                    </el-tag>
                    <el-input v-if="form.labelInputVisible" ref="labelInputRef" v-model="form.labelText" size="small"
                        class="label-input" @keyup.enter="labelHandl.handleInputConfirm"
                        @blur="labelHandl.handleInputConfirm" />
                    <el-button class="label-btn" v-else size="small" @click="labelHandl.showInput">
                        + New Tag
                    </el-button>
                </el-form-item>
                <el-button size="small" type="primary"
                    @click="useSaveData(form, ruleFormRef, router, uploadFileformation, uploadCoveration,props)">{{ props.type ==
                        "edit" ? "更新专栏" : "发布专栏" }}</el-button>
            </el-form>
        </div>
        <!-- 上传组件 -->
        <div class="upload-box" v-show="false">
            <el-upload class="upload" drag action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                ref="uploadRef" multiple :on-success="handle.handleFileSuccess" :on-error="handle.handleFileError"
                :show-file-list="false" :before-upload="handle.beforeFileUpload" :auto-upload="true"
                :http-request="handle.RedefineUploadFile">
                <el-button size="small" type="primary" ref="uploadBtnRef">上传文件</el-button>
            </el-upload>
        </div>
        <!-- 文件上传进度 -->
        <el-progress ref="uploadProgressRef" v-show="false" :text-inside="true" :stroke-width="16" type="dashboard"
            :percentage="uploadFileformation.progress" :color="colors" />
        <!-- 富文本组件 -->
        <div class="quill-box" v-show="!form.isShow">
            <el-page-header @back="form.isShow = true" class="page-header">
                <template #content>
                    <span class=""> {{ form.title }} </span>
                </template>
            </el-page-header>
            <QuillEditor id="editorId" ref="myQuillEditor" v-model:content="form.content" contentType="html"
                :options="options" theme="snow" />
        </div>
    </div>
</template>

<script setup lang="ts">
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import 'highlight.js/styles/agate.css'
import { QuillEditor } from '@vueup/vue-quill'
import { Plus, Edit } from '@element-plus/icons-vue'
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { useArticleContributionProp, useHandleFileMethod, useHandleCoverMethod, useInit, userLabelHandlMethod, useSaveData, useRules } from '@/logic/creation/contribute/contributePage/articleContribution';
import { onMounted } from 'vue';


const props = defineProps({
    type: {
        type: String
    }
})

components: {
    QuillEditor
}

const { myQuillEditor, form, options, uploadFileformation, uploadCoveration, uploadRef, uploadBtnRef, uploadProgressRef, colors, labelInputRef, ruleFormRef, router, cascader, cascaderValue, editArticleStore } = useArticleContributionProp()
const handle = useHandleFileMethod(uploadFileformation, form, myQuillEditor, uploadProgressRef)
const handleCover = useHandleCoverMethod(uploadCoveration)
const labelHandl = userLabelHandlMethod(form, labelInputRef)
const rules = useRules()
onMounted(() => {
    useInit(uploadFileformation, uploadCoveration, cascader, form, editArticleStore, props, myQuillEditor, cascaderValue)
})

//分类选择
const cascaderHandleChange = (value: any) => {
    form.classification_id = value[value.length - 1]
    console.log(form.classification_id)
}


</script>

<style scoped lang="scss">
@import "@/assets/styles/views/creation/contribute/contributePage/articleContribution.scss"
</style>
