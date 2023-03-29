interface UserInfo {
    username : string
}

export interface GetFavoritesListItem {
    id : number
    title :string
    content :string
    cover :string
    type : string
    src : string
    userInfo : UserInfo
}

export type GetFavoritesListRes = Array<GetFavoritesListItem>


export interface DeleteFavoritesReq {
    id :number
}

export interface FavoriteVideoReq {
    ids :Array<Number>
    video_id : number
}

export interface GetFavoritesListByFavoriteVideoReq {
    video_id : number
}

export interface GetFavoritesListByFavoriteVideoItem {
    id : number
    title :string
    content :string
    cover :string
    type : string
    src : string
    userInfo : UserInfo
    present :number
    max :number
    selected : boolean //是否已经收藏
    choose : boolean | undefined //用于收藏视频时选择器
}


export type getFavoritesListByFavoriteVideoRes = Array<GetFavoritesListByFavoriteVideoItem>
