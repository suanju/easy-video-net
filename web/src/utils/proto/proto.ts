import {Message,encodeMessage} from "@/proto/pb/live" 


export const encodeProtoFormat = (msgType :string ,data : Uint8Array) => {
    let encodeData = encodeMessage(<Message>{msgType : msgType ,data :data});
    return encodeData;
}