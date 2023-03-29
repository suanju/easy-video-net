export interface UserInfoRes {
    username : string
    gender : number
    birth_date : string
    is_visible : boolean
    signature : string
}

export interface DetermineNameExistsReq {
    username : string
}

export type DetermineNameExistsRes = boolean

export type SetUserInfoRes = boolean 

export interface UpdateAvatarReq {
    imgUrl :string
    type : string
}

export interface AttentionReq {
    uid :number
}