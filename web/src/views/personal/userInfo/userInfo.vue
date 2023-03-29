<template>
  <div class="overall">
    <pageHeader title="编辑信息" icon-nmae="userData"></pageHeader>
    <div class="form-box personal-layout animate__animated animate__slideInRight ">
      <el-form :model="form" :rules="userInfoRules" label-width="120px">
        <el-form-item label="昵称" prop="username">
          <el-input class="input-name" v-model="form.username" clearable />
        </el-form-item>
        <el-form-item label="性别">
          <el-select v-model="form.gender" placeholder="请选择您的性别">
            <el-option label="男生" :value="0" />
            <el-option label="女生" :value="1" />
            <el-option label="薛定谔的猫" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="出生日期">
          <el-col :span="11">
            <el-date-picker
              v-model="form.birth_date"
              type="date"
              placeholder="请选出生年月日"
              style="width: 100%"
            />
          </el-col>
        </el-form-item>
        <el-form-item label="他人可见">
          <el-switch class="switch-btn" v-model="form.is_visible" />
        </el-form-item>

        <el-form-item label="个性签名">
          <el-input
            class="input-hobby"
            v-model="form.signature"
            type="textarea"
            resize="none"
            rows="4"
          />
        </el-form-item>
        <el-form-item>
          <el-button v-removeFocus type="primary" round @click="UserInfoMethod.onSubmit">
            修改资料</el-button
          >
        </el-form-item>
      </el-form>
      <div class="figure-box">
        <div class="figure"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  useUserInfoProp,
  useUserInfoMethod,
  useRules,
} from "@/logic/personal/userInfo/userInfo";
import pageHeader from "@/components/personalNavigation/pageHeader.vue";
import {vRemoveFocus} from "@/utils/customInstruction/focus"
import { onMounted } from "vue";


components: {
  pageHeader;
}
const { form } = useUserInfoProp();
const UserInfoMethod = useUserInfoMethod(form);
const { userInfoRules } = useRules();

onMounted(() => {
  UserInfoMethod.getUserInfo();

})


</script>
<style scoped lang="scss">
@import "@/assets/styles/views/personal/userInfo/userInfo.scss";
@import "@/assets/styles/global/el-style.scss";
</style>
