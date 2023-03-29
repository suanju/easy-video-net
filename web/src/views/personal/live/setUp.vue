<template>
  <div class="overall">
    <pageHeader title="直播设置" icon-nmae="live"></pageHeader>
    <div class=" principal  personal-layout animate__animated animate__slideInRight">
      <div class="form-box">
        <el-upload class="cover-uploader" :action="liveInformationForm.action" :show-file-list="false"
          :on-success="handle.handleFileSuccess" :on-error="handle.handleFileError"
          :before-upload="handle.beforeFileUpload" :auto-upload="true" :http-request="handle.RedefineUploadFile"
          accept=".png,.jpg,.jpeg">
          <img v-if="liveInformationForm.FileUrl" :src="liveInformationForm.FileUrl" class="cover" />
          <el-icon v-else class="cover-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
        <div>
          <div class="form-show">
            <el-form :model="liveInformationForm" ref="saveDateFormRef" :rules="liveInformationRules"
              label-position="left" label-width="5rem">

              <el-form-item label="直播标题" prop="title">
                <el-input v-model="liveInformationForm.title" />
              </el-form-item>
              <el-form-item label="Adders">
                {{ userStore.userInfoData.liveData.address }}
              </el-form-item>
              <el-form-item label="key">
                {{ liveKeyDesensitization(userStore.userInfoData.liveData.key) }} <el-button
                  @click="useCopy(userStore.userInfoData.liveData.key)" class="copy" color="#626aef" size="small" plain
                  round>
                  copy
                </el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
      <div class="bottom-box">
        <span class="text"> 请设置您的直播封面和标题,以便更好吸引观众嗷~</span>
        <div class="button">
          <el-button v-removeFocus @click="useSaveData(liveInformationForm, saveDateFormRef, rawData)" type="primary"
            round>保存资料
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { Plus } from "@element-plus/icons-vue";
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { useLiveInfoProp, useHandleFileMethod, useSaveData, useCopy, useRules, useInit } from "@/logic/personal/live/setUpLive"
import { liveKeyDesensitization } from "@/utils/conversion/stringConversion";
import { onMounted } from "vue";


const { userStore, liveInformationForm, saveDateFormRef, rawData } = useLiveInfoProp()
const handle = useHandleFileMethod(liveInformationForm)
const { liveInformationRules } = useRules();

onMounted(() => {
  useInit(liveInformationForm, rawData)
})

</script>
<style scoped lang="scss">
@import "@/assets/styles/views/personal/live/setUp.scss";
</style>
