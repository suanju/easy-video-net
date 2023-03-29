import { Message, decodeWebClientSendBarrageRes ,decodeWebClientEnterLiveRoomRes ,decodeWebClientHistoricalBarrageRes} from "@/proto/pb/live"
import DPlayer, { DPlayerDanmakuItem } from "dplayer"
import globalScss from "@/assets/styles/global/export.module.scss"
import Swal from "sweetalert2";
import { Ref } from "vue";
import { Uint8ArrayToString } from "@/utils/conversion/unt8Array";
 import router from "@/router";

export const webError = (data: Message) => {
    Swal.fire({
        title: new TextDecoder("utf-8").decode(data.data as Uint8Array),
        heightAuto: false,
        confirmButtonColor: globalScss.colorButtonTheme,
        icon: "error",
    })
    router.back()
    return
}

export const webClientBarrageDeal = (data: Message, dp: DPlayer, sideRef: Ref<any>) => {
    //格式化消息
    if (!data.data) return;
    const msg = decodeWebClientSendBarrageRes(data.data)
    console.log(msg)
    sideRef.value.addBarrage(msg)
    dp.danmaku.draw(<DPlayerDanmakuItem>{
        text: msg.text,
        color: msg.color,
        type: msg.type,
    });
    console.log(msg)
}

export const webClientEnterLiveRoomDeal = (data: Message, dp: DPlayer, sideRef: Ref<any>) => {
    //格式化消息
    if (!data.data) return;
    const msg = decodeWebClientEnterLiveRoomRes(data.data)
    const list = msg.list
    sideRef.value.updataOnline(list)
    console.log(msg)

}

export const webClientHistoricalBarrageRes = (data: Message, dp: DPlayer, sideRef: Ref<any>) => {
    //格式化消息
    if (!data.data) return;
    const msg = decodeWebClientHistoricalBarrageRes(data.data)
    sideRef.value.addHistoryBarrage(msg)
    console.log(msg)
}