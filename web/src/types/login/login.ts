export interface loginReq  {
    username :string,
    password : string
}

export interface userInfoRes {
    id : number
    photo : string
    token : string
    username : string
    created_at : string
}

export interface registReq {
    username :string
    password :string
    verificationCode :string
    email :string
}

export interface forgetReq {
    password :string
    verificationCode :string
    email :string
}

export interface sendEmailReq{
    email : string
}

export type sendEmailType = "regist" |"modify" 


export interface sendEmailInfo{
    btnText :string
    isPleaseClick : Boolean
    
}