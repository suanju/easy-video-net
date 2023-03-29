import httpRequest from "@/utils/requst"
import { loginReq , userInfoRes ,sendEmailReq,registReq ,forgetReq } from "@/types/login/login"


export const loginRequist = (params: loginReq) => {
    return httpRequest.post<userInfoRes>('/login/login', params);
}
export const regist = (params: registReq) => {
    return httpRequest.post<userInfoRes>('/login/register', params);
}
export const sendEmailVerificationCode = (params: sendEmailReq) => {
    return httpRequest.post('/login/sendEmailVerificationCode', params);
}

export const sendEmailVerificationCodeByForget = (params: sendEmailReq) => {
    return httpRequest.post('/login/sendEmailVerificationCodeByForget', params);
}

export const forgetRequist = (params: forgetReq) => {
    return httpRequest.post('/login/forget', params);
}

