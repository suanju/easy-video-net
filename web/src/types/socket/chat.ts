import { GetChatListItem } from "../personal/chat/chat";
import { MessageInfo } from "../store/chat";

export interface ChatSendTextMsg extends MessageInfo {

}
export interface ChatUnreadNotic {
    item: any;
    uid: number;
    tid: number;
    last_message: string;
    last_message_info: MessageInfo;
    unread: number;
}

export type ChatOnlineUnreadNotice = Array<GetChatListItem>