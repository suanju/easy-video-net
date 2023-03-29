import { getChatList } from "@/apis/personal";
import { useChatListStore } from "@/store/chat";

export default async () => {
    const chatListStore = useChatListStore()
    try {
        let response = await getChatList()
        if (!response.data) return false
        const chatList = response.data.filter((item) => {
            item.message_list = item.message_list.reverse()
            return item
        })
        chatListStore.chatListData = chatList
    } catch (err: any) {
        console.error(err)
    }
}