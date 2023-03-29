import httpRequest from "@/utils/requst"

import { GetLiveRoomRes } from "@/types/home/home"
import { GetLiveRoomInfoReq, GetLiveRoomInfoRes } from "@/types/live/liveRoom";
import { GetBeLiveListRes } from "@/types/home/live";


export const getLiveRoom = () => {
    return httpRequest.post<GetLiveRoomRes>('/live/getLiveRoom');
}

//获取直播信息
export const getLiveRoomInfo = (params: GetLiveRoomInfoReq) => {
    return httpRequest.post<GetLiveRoomInfoRes>('/live/getLiveRoomInfo', params);
}

//获取正在直播的列表
export const getBeLiveList = () => {
    return httpRequest.post<GetBeLiveListRes>('/live/getBeLiveList');
}
