import { useChatListStore } from '@/store/chat';
import { userInfoRes } from '@/types/login/login';
import { GetSpaceIndividualRes } from '@/types/space/space';
import { global, space, userInfo } from '@/types/store/main';
import { defineStore } from 'pinia';
import { reactive, ref } from "vue";
import { OssSTSInfo } from './../types/idnex';

export const useUserStore = defineStore("user", () => {
    let userInfoData = reactive<userInfo>({
        id: 0,
        username: "",
        photo: "",
        token: "",
        created_at: "",
        liveData: {
            address: "",
            key: ""
        },
        //消息相关
        unreadNotice: 0
    })

    const setUserInfo = (info: userInfoRes) => {
        userInfoData.username = info.username
        userInfoData.id = info.id
        userInfoData.photo = info.photo
        userInfoData.token = info.token
        userInfoData.created_at = info.created_at
    }

    const setUnreadNotice = (num: number) => {
        userInfoData.unreadNotice = num
    }

    const loginOut = () => {
        userInfoData.username = ""
        userInfoData.id = 0
        userInfoData.photo = ""
        userInfoData.token = ""
        userInfoData.created_at = ""
        userInfoData.unreadNotice = 0
        //清空消息
        let chat = useChatListStore()
        chat.chatListData = []
        chat.tid = 0
        chat.tUsername = ""
    }
    return {
        userInfoData,
        setUserInfo,
        setUnreadNotice,
        loginOut
    }

}, {
    persist: true,
})


export const useSpaceStore = defineStore("space", () => {
    const spaceInfoData = reactive<space>({
        id: 0
    })

    const spaceUserData = reactive<GetSpaceIndividualRes>({
        id: 0,
        username: '',
        photo: '',
        signature: '',
        is_attention: false,
        attention_num: 0,
        vermicelli_num: 0
    })

    const setSpaceID = (id: number) => {
        spaceInfoData.id = id
    }
    return {
        spaceInfoData,
        setSpaceID
    }
}, {
    persist: true,
})

export const useGlobalStore = defineStore("global", () => {
    const globalData = reactive<global>({
        router: {
            currentRouter: ""
        },
        loading: {
            loading: false,
            loadingText: "努力加载中!",
            loadingBackground: "rgba(122, 122, 122, 0.3)",
            loadingSvg: `
            <svg xmlns="http://www.w3.org/2000/svg" width="44" height="44" viewBox="0 0 44 44" stroke="#fff">
            <g fill="none" fill-rule="evenodd" stroke-width="2">
            <circle cx="22" cy="22" r="1">
            <animate attributeName="r" begin="0s" dur="1.8s" values="1; 20" calcMode="spline" keyTimes="0; 1" keySplines="0.165, 0.84, 0.44, 1" repeatCount="indefinite"/>
            <animate attributeName="stroke-opacity" begin="0s" dur="1.8s" values="1; 0" calcMode="spline" keyTimes="0; 1" keySplines="0.3, 0.61, 0.355, 1" repeatCount="indefinite"/>
            </circle>
            <circle cx="22" cy="22" r="1">
            <animate attributeName="r" begin="-0.9s" dur="1.8s" values="1; 20" calcMode="spline" keyTimes="0; 1" keySplines="0.165, 0.84, 0.44, 1" repeatCount="indefinite"/>
            <animate attributeName="stroke-opacity" begin="-0.9s" dur="1.8s" values="1; 0" calcMode="spline" keyTimes="0; 1" keySplines="0.3, 0.61, 0.355, 1" repeatCount="indefinite"/>
            </circle>
            </g>
            </svg>
            `,
            loadingSvgViewBox: "-10, -10, 50, 50",
        },
        scroll: {
            scrollLeft: 0,
            scrollTop: 0,
        }
    }
    )

    const ossData = ref(<OssSTSInfo>{
        region: "",
        accessKeyId: "",
        accessKeySecret: "",
        stsToken: "",
        bucket: "",
        expirationTime: 0,
    })

    const setGlobalLoading = (boll: boolean) => {
        globalData.loading.loading = boll
    }

    const setOssInfo = (data: OssSTSInfo) => {
        ossData.value = data
    }

    const setScroll = (scrollLeft: number, scrollTop: number) => {
        globalData.scroll.scrollLeft = scrollLeft
        globalData.scroll.scrollTop = scrollTop
    }
    return {
        globalData,
        ossData,
        setGlobalLoading,
        setOssInfo,
        setScroll
    }

}, {
    persist: true,
})
