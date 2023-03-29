import { EnterLiveRoom, WebClientSendBarrageRes} from "@/proto/pb/live"
export  interface sideData {
    baeeage : Array<WebClientSendBarrageRes>
    online : Array<EnterLiveRoom>
}