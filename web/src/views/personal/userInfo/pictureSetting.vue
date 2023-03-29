<template>
  <div class="overall">
    <pageHeader title="更换头像" icon-nmae="pictures"></pageHeader>
    <div class=" principal  personal-layout animate__animated animate__slideInRight">
      <div class="form-box">
        <el-upload class="avatar-uploader" :action="uploadAvatarForm.action" :show-file-list="false"
          :on-success="handle.handleFileSuccess" :on-error="handle.handleFileError"
          :before-upload="handle.beforeFileUpload" :auto-upload="true" :http-request="handle.redefineUploadFile">
          <img v-if="uploadAvatarForm.FileUrl" :src="uploadAvatarForm.FileUrl" accept=".png,.jpg,.jpeg" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
        <div>
          <div class="avatar-show">
            <el-avatar :size="240" :src="userStore.userInfoData.photo" />
          </div>
        </div>
      </div>
      <div class="bottom-box">
        <span class="text"> 请选择图片上传:大小180 * 180像素支持JPG、PNG等格式,图片需小于2M</span>
        <div class="button">
          <el-button v-removeFocus @click="useUpdateAvatar(userStore, uploadAvatarForm)" type="primary" round>修改头像
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { Plus } from "@element-plus/icons-vue";
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { useAvatarProp, useHandleFileMethod, useUpdateAvatar, useInit } from "@/logic/personal/userInfo/pictureSetting"
import { onMounted } from "vue";

const { userStore, uploadAvatarForm } = useAvatarProp()
const handle = useHandleFileMethod(uploadAvatarForm)


onMounted(() => {
  useInit(uploadAvatarForm)
})

</script>
<style scoped lang="scss">
@import "@/assets/styles/views/personal/userInfo/pictureSetting.scss";
</style>
