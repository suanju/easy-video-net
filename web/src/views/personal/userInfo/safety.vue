<template>
    <div class="overall">
        <pageHeader title="用户安全" icon-nmae="security"></pageHeader>
        <div class="form-box personal-layout animate__animated animate__slideInRight ">
            <el-form :model="form" :rules="changePasswor" label-width="120px">
                <el-form-item label="邮箱验证码" prop="username">
                    <el-input class="input-name" v-model="form.verificationCode" clearable> <template #append>
                            <el-button type="primary" size="small" round @click="sendVerificationCode(sendVerificationCodeInfo)">{{ sendVerificationCodeInfo.btnText }}</el-button>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item label="新密码" prop="username">
                    <el-input class="input-name" v-model="form.password" clearable />
                </el-form-item>
                <el-form-item label="确认密码" prop="username">
                    <el-input class="input-name" v-model="form.confirm_password" clearable />
                </el-form-item>
                <el-form-item>
                    <el-button v-removeFocus type="primary" round @click="onSubmit">
                        修改密码</el-button>
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
    useSafetyProp,
    useSafetyMethod,
    useRules,
} from "@/logic/personal/userInfo/safety";
import pageHeader from "@/components/personalNavigation/pageHeader.vue";
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { onMounted } from "vue";


components: {
    pageHeader;
}
const { form ,sendVerificationCodeInfo} = useSafetyProp();
const { sendVerificationCode, onSubmit } = useSafetyMethod(form);
const { changePasswor } = useRules();

onMounted(() => {

})


</script>
<style scoped lang="scss">
@import "@/assets/styles/views/personal/userInfo/safety.scss";
@import "@/assets/styles/global/el-style.scss";
</style>
