import { danmakuApi, getVideoBarrageList, getVideoContributionByID, likeVideo, sendVideoBarrage } from "@/apis/contribution";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useGlobalStore, useUserStore } from "@/store/main";
import { GetVideoBarrageListReq, GetVideoContributionByIDReq, LikeVideoReq, SendVideoBarrageReq, VideoInfo } from "@/types/show/video/video";
import DPlayer, { DPlayerDanmakuItem, DPlayerVideoQuality } from "dplayer";
import Swal from 'sweetalert2';
import { Ref, UnwrapNestedRefs, reactive, ref } from "vue";
import { RouteLocationNormalizedLoaded, Router, useRoute, useRouter } from "vue-router";
import { numberOfViewers, responseBarrageNum } from './socketFun';


export const useVideoProp = () => {
  const route = useRoute()
  const router = useRouter()
  const global = useGlobalStore()
  const userStore = useUserStore()
  const videoRef = ref()
  const videoID = ref<number>(0)
  const videoInfo = reactive(<VideoInfo>{})
  const barrageInput = ref("")
  const barrageListShow = ref(false)
  const videoBarrage = ref(true)
  const liveNumber = ref(0)
  //回复二级评论
  const replyCommentsDialog = reactive({
    show: false,
    commentsID: 0,
  })

  return {
    route,
    router,
    userStore,
    videoRef,
    videoID,
    videoInfo,
    barrageInput,
    barrageListShow,
    liveNumber,
    replyCommentsDialog,
    videoBarrage,
    global
  }
}

export const useSendBarrage = async (text: Ref<string>, dpaler: DPlayer, userStore: any, videoInfo: UnwrapNestedRefs<VideoInfo>, socket: WebSocket) => {
  const res = await sendVideoBarrage(<SendVideoBarrageReq>{
    author: userStore.userInfoData.username,
    color: 16777215,
    id: videoInfo.videoInfo.id.toString(),
    text: text.value,
    time: dpaler.video.currentTime,
    type: 0,
    token: userStore.userInfoData.token,
  })

  console.log(userStore.userInfoData)
  if (res.code != 0) {
    Swal.fire({
      title: "弹幕服务异常",
      heightAuto: false,
      confirmButtonColor: globalScss.colorButtonTheme,
      icon: "error",
    })
    return
  }
  const danmaku = <DPlayerDanmakuItem>{
    text: text.value,
    color: '#fff',
    type: 'right',
  };

  dpaler.danmaku.draw(danmaku);

  text.value = ""

  let data = JSON.stringify({
    type: "sendBarrage",
    data: ""
  })
  socket.send(data)

}

export const useLikeVideo = async (videoInfo: UnwrapNestedRefs<VideoInfo>) => {
  try {
    await likeVideo(<LikeVideoReq>{
      video_id: videoInfo.videoInfo.id
    })
    videoInfo.videoInfo.is_like = !videoInfo.videoInfo.is_like
  } catch (err) {
    Swal.fire({
      title: "点赞失败",
      heightAuto: false,
      confirmButtonColor: globalScss.colorButtonTheme,
      icon: "error",
    })
  }
}

export const useInit = async (videoRef: Ref, route: RouteLocationNormalizedLoaded, Router: Router, videoID: Ref<Number>, videoInfo: UnwrapNestedRefs<VideoInfo>, global: any) => {
  try {
    //绑定视频id
    if (!route.params.id) {
      Router.back()
      Swal.fire({
        title: "获取视频失败",
        heightAuto: false,
        confirmButtonColor: globalScss.colorButtonTheme,
        icon: "error",
      })
      Router.back()
      return
    }
    global.globalData.loading.loading = true
    videoID.value = Number(route.params.id)
    //得到视频信息
    const vinfo = await getVideoContributionByID(<GetVideoContributionByIDReq>{
      video_id: videoID.value
    })
    if (!vinfo.data) return false
    videoInfo.videoInfo = vinfo.data.videoInfo
    videoInfo.recommendList = vinfo.data.recommendList

    //得到清晰度列表
    let quality: DPlayerVideoQuality[] = []
    if (videoInfo.videoInfo.video) {
      quality = [...quality, {
        name: "1080P超清",
        url: videoInfo.videoInfo.video
      }]
    }
    if (videoInfo.videoInfo.video_720p) {
      quality = [...quality, {
        name: "720P高清",
        url: videoInfo.videoInfo.video_720p
      }]
    }
    if (videoInfo.videoInfo.video_480p) {
      quality = [...quality, {
        name: "480P标清",
        url: videoInfo.videoInfo.video_480p
      }]
    }
    if (videoInfo.videoInfo.video_360p) {
      quality = [...quality, {
        name: "360P流畅",
        url: videoInfo.videoInfo.video_360p
      }]
    }
    //获取视频弹幕信息
    const barrageList = await getVideoBarrageList(<GetVideoBarrageListReq>{
      id: videoID.value.toString()
    })
    if (!barrageList.data) return false
    videoInfo.barrageList = barrageList.data
    //获取当前用户信息
    const userStore = useUserStore()
    //初始化播放器
    const dp = new DPlayer({
      container: videoRef.value,
      loop: true, // 循环播放
      lang: "zh-cn", // 语言
      logo: "",
      autoplay: true,
      danmaku: {
        id: videoID.value.toString(),
        api: danmakuApi,
        token: userStore.userInfoData.token
      },
      mutex: false, // 互斥，阻止多个播放器同时播放
      video: {
        quality: quality,
        defaultQuality: 0,
        url: "不填", // 视频链接
        pic: videoInfo.videoInfo.cover
      },
    });
    global.globalData.loading.loading = false
    return dp
  } catch (err) {
    global.globalData.loading.loading = false
    console.log(err)
  }
}

export const useWebSocket = (userStore: any, videoInfo: UnwrapNestedRefs<VideoInfo>, Router: Router, liveNumber: Ref<number>) => {
  let socket: WebSocket
  const open = () => {
    console.log("websocket 连接成功 ")
  }
  const error = () => {
    console.error("websocket 连接失败")
  }
  const getMessage = async (msg: any) => {
    let data = JSON.parse(msg.data)
    console.log(data)
    switch (data.type) {
      case "numberOfViewers":
        numberOfViewers(liveNumber, data.data.people)
        break;
      case "responseBarrageNum":
        responseBarrageNum(videoInfo)
        break;
    }
  }

  if (typeof (WebSocket) === "undefined") {
    Swal.fire({
      title: "您的浏览器不支持socket",
      heightAuto: false,
      confirmButtonColor: globalScss.colorButtonTheme,
      icon: "error",
    })
    Router.back()
    return
  } else {
    // 实例化socket
    socket = new WebSocket(import.meta.env.VITE_SOCKET_URL + "/ws/videoSocket?token=" + userStore.userInfoData.token + "&videoID=" + videoInfo.videoInfo.id)
    // 监听socket连接
    socket.onopen = open
    // 监听socket错误信息
    socket.onerror = error
    // 监听socket消息
    socket.onmessage = getMessage
  }

  return socket
}
