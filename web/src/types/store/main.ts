
export interface userInfo {
    id: number
    photo: string
    token: string
    username: string
    created_at: string
    liveData: {
        address: string
        key: string
    }
    unreadNotice: number
}

export interface global {
    router: {
        currentRouter: string
    }
    loading: {
        loading: boolean
        loadingText?: string
        loadingBackground?: string
        loadingSvg?: string
        loadingSvgViewBox?: string
    }
    scroll: {
        scrollTop: number
        scrollLeft: number
    }
}

export interface space {
    id: number
}