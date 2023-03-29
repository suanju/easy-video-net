
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

export interface ChatInfo {
    to_id: number;
    username: string;
    photo: string;
    unread: number;
    last_message: string;
    message_list: Array<MessageInfo>;
    last_at: string;
    updated_at: string;
}

export type ChatList = Array<ChatInfo>
