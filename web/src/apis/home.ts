import { GetHomeInfoReq, GetHomeInfoRes } from "@/types/home/home";
import httpRequest from "@/utils/requst"

//获取主页信息
export const getHomeInfo = (params: GetHomeInfoReq) => {
    return httpRequest.post<GetHomeInfoRes>('/home/getHomeInfo',params);
}