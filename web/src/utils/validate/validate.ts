import {determineNameExistsRequist} from "@/apis/personal"
import {DetermineNameExistsReq} from "@/types/personal/userInfo/userInfo"
//验证验证码
export const validateUsername = (rule: any, value: any, callback: any) => {
    value as string;
    console.log(value)
    if (value === '') {
      callback(new Error('请输入用户名'));
    } else if (value.length < 2 || value.length > 8) {
      callback(new Error('昵称长度需要在2-8位之间'));
    } else{
      callback();
    }
  };
  //验证密码
  export const validatePassword = (rule: any, value: any, callback: any) => {
    const reg = new RegExp(/[a-zA-Z0-9!?.]+/);
    if (value === '') {
      callback(new Error('请输入密码'));
    } else if (value.length < 6 || value.length > 16) {
      callback(new Error('密码长度需要在6-16位之间'));
    } else if (!reg.test(value)) {
      callback(new Error('密码中不能包含除了!?.之外的特殊符号'));
    } else {
      callback();
    }
  }; 
  //验证验证码
  export const validateVarCode = (rule: any, value: any, callback: any) => {
    value as string;
    if (value === '') {
      callback(new Error('请输入验证码'));
    } else if (value.length != 6) {
      callback(new Error('验证码为6位数嗷'));
    } else {
      callback();
    }
  };
  //验证邮箱
  export const validateEmail = (rule: any, value: any, callback: any) => {
    const reg = new RegExp(/[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+/);
    if (value.length === 0) {
      callback(new Error('可以告诉我你的邮箱吗~'));
    } else if (value.length < 5 || value.length > 64) {
      callback(new Error('你的邮箱长度有点问题~'));
    } else if (!reg.test(value)) {
      callback(new Error('可以输入正确的邮箱格式吗~'));
    } else {
      callback();
    }
  };


   //验证重复昵称
   export const validateRepeatName = async (rule: any, value: any, callback: any) => {
    value as string;
    let requist = <DetermineNameExistsReq>{
      username : value
    }
    console.log(value)
    if (value === '') {
      callback(new Error('请输入用户名'));
    } else if (value.length < 2 || value.length > 8) {
      callback(new Error('昵称长度需要在2-16位之间'));
    } else if ((await determineNameExistsRequist(requist)).data) {
      callback(new Error('昵称有人用过了嗷'));
    } else{
      callback();
    }
};


   //验证直播标题
   export const validateLiveTitle = async (rule: any, value: any, callback: any) => {
    value as string;
    if (value === '') {
      callback(new Error('请输入直播标题'));
    } else if (value.length < 8 || value.length > 20) {
      callback(new Error('标题长度需要在8-20位之间'));
    } else{
      callback();
    }
};

  //验证视频标题
  export const validateVideoTitle = async (rule: any, value: any, callback: any) => {
    value as string;
    if (value === '') {
      callback(new Error('请输入直播标题'));
    } else if (value.length < 4 || value.length > 30) {
      callback(new Error('标题长度需要在4-30位之间'));
    } else{
      callback();
    }
};

  //验证视频介绍
  export const validateVideoIntroduce = async (rule: any, value: any, callback: any) => {
    value as string;
    if (value === '') {
      callback(new Error('请输入直播介绍'));
    }else{
      callback();
    }
};

  //验证视频标题
  export const validateArticleTitle = async (rule: any, value: any, callback: any) => {
    value as string;
    if (value === '') {
      callback(new Error('请输入文章标题'));
    } else if (value.length < 8 || value.length > 20) {
      callback(new Error('标题长度需要在8-20位之间'));
    } else{
      callback();
    }
};



   //验证收藏夹标题
   export const validateCollectTitle = async (rule: any, value: any, callback: any) => {
    value as string;
    if (value === '') {
      callback(new Error('请输入标题'));
    } else if (value.length < 1 || value.length > 20) {
      callback(new Error('标题长度需要在1-20位之间'));
    } else{
      callback();
    }
};