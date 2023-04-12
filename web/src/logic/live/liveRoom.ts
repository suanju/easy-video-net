
import { getLiveRoomInfo } from "@/apis/live";
import globalScss from "@/assets/styles/global/export.module.scss";
import { webClientBarrageDeal, webClientEnterLiveRoomDeal, webClientHistoricalBarrageRes, webError } from "@/logic/live/socketFun";
import { decodeMessage } from "@/proto/pb/live";
import { useUserStore } from "@/store/main";
import { GetLiveRoomInfoReq, GetLiveRoomInfoRes } from "@/types/live/liveRoom";
import DPlayer, { DPlayerDanmaku } from "dplayer";
import flvJs from "flv.js";
import Swal from 'sweetalert2';
import { nextTick, reactive, Ref, ref, UnwrapNestedRefs } from "vue";
import { RouteLocationNormalizedLoaded, Router, useRoute, useRouter } from "vue-router";
export const useLiveRoomProp = () => {
  const route = useRoute()
  const router = useRouter()
  const videoRef = ref()
  const userStore = useUserStore()
  const roomID = ref<number>(0)
  const liveInfo = reactive(<GetLiveRoomInfoRes>{})

  return {
    route,
    router,
    videoRef,
    userStore,
    roomID,
    liveInfo
  }
}

export const useWebSocket = (dp: DPlayer, userStore: any, sideRef: Ref<any>, roomID: Ref<Number>, Router: Router) => {
  let socket: WebSocket
  const initWebSocket = (() => {
    const open = () => {
      console.log("websocket 连接成功 ")
    }
    const error = () => {
      console.error("websocket 连接失败 ")
    }
    const getMessage = async (msg: any) => {
      console.log(msg.data)
      //转义uint8ARRAY
      const reader = new FileReader();
      reader.readAsArrayBuffer(msg.data);
      reader.onload = function (e) {
        let buf = new Uint8Array(reader.result as ArrayBuffer);
        const response = decodeMessage(buf)
        switch (response.msgType) {
          case "error":
            webError(response)
            break;
          case "webClientBarrageRes":
            webClientBarrageDeal(response, dp, sideRef)
            break;
          case "webClientEnterLiveRoomRes":
            webClientEnterLiveRoomDeal(response, dp, sideRef)
            break;
          case "webClientHistoricalBarrageRes":
            webClientHistoricalBarrageRes(response, dp, sideRef)
            break;
          default: console.error("未支持的消息类型")
            break;
        }
        console.log(response)
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
      socket = new WebSocket(import.meta.env.VITE_SOCKET_URL + "/ws/liveSocket?token=" + userStore.userInfoData.token + "&liveRoom=" + roomID.value)
      // 监听socket连接
      socket.onopen = open
      // 监听socket错误信息
      socket.onerror = error
      // 监听socket消息
      socket.onmessage = getMessage
    }
  })()

  const sendMessage = (msg: string | ArrayBufferLike | Blob | ArrayBufferView) => {
    console.log("发送消息", msg)
    socket.send(msg)
  }

  return {
    sendMessage
  }

}


export const useInit = async (videoRef: Ref, route: RouteLocationNormalizedLoaded, Router: Router, roomID: Ref<Number>, liveInfo: UnwrapNestedRefs<GetLiveRoomInfoRes>) => {
  try {
    //绑定房间
    if (!route.params.id) {
      Router.back()
      Swal.fire({
        title: "访问房间失败",
        heightAuto: false,
        confirmButtonColor: globalScss.colorButtonTheme,
        icon: "error",
      })
      Router.back()
      return
    }
    roomID.value = Number(route.params.id)

    //获取直播信息

    const li = await getLiveRoomInfo(<GetLiveRoomInfoReq>{
      room_id: roomID.value
    })

    if (!li.data) {
      Swal.fire({
        title: "该用户未进行直播配置",
        heightAuto: false,
        confirmButtonColor: globalScss.colorButtonTheme,
        icon: "error",
      })
      return
    }
    liveInfo.live_title = li.data?.live_title
    liveInfo.photo = li.data?.photo
    liveInfo.username = li.data?.username
    liveInfo.flv = li.data?.flv



    //初始化播放器
    console.log(videoRef)
    const dp = new DPlayer({
      container: videoRef.value, // 容器
      loop: true, // 循环播放
      lang: "zh-cn", // 语言，可选'en', 'zh-cn', 'zh-tw',
      logo: "", // 在左上角展示一个logo
      autoplay: true,
      danmaku: true as unknown as DPlayerDanmaku, //官方文档给的就是true 但是ts中规定类型不一致取的取舍方案
      apiBackend: {
        read: function (options) {
          console.log('Pretend to connect WebSocket');
          options.success([]);
        },
        send: function (options) {
          console.log('Pretend to send danmaku via WebSocket', options.data);
          options.success();
        },
      },
      mutex: false, // 互斥，阻止多个播放器同时播放
      video: { // 视频信息
        type: "customFlv", // 视频类型 可选"auto", "hls", "flv", "dash"..
        url: liveInfo.flv, // 视频链接
        customType: {
          customFlv: (video: any, player: any) => {
            const flvPlayer = flvJs.createPlayer({
              type: "flv",
              isLive: true,
              hasAudio: true,
              url: liveInfo.flv,
            });
            nextTick(() => {
              flvPlayer.attachMediaElement(video);
              flvPlayer.load();
            })
          }
        },
      },
    });
    return dp
  } catch (err) {
    console.log(err)
  }
}
