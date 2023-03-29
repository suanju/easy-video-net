
export interface MessageInfo {
    id: number;
    uid: number;
    username: string;
    photo: string;
    tid: number;
    message: string;
    type: string;
    created_at: string;
}

export interface GetChatListItem {
    to_id: number;
    username: string;
    photo: string;
    unread: number;
    last_message: string;
    message_list: Array<MessageInfo>
    last_at: string;
    updated_at: string;
}
export type GetChatListRes = Array<GetChatListItem>

export interface PersonalLetterReq {
    id: number
}

export interface DeleteChatItemReq {
    id: number
}

export interface GetChatHistoryMsgReq {
    tid: number;
    last_time: number; //最后时间
}

export type GetChatHistoryMsgRes = Array<MessageInfo>