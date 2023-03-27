package socket

import (
	"time"
)

type NoticeMessageStruct struct {
	NoticeType string `json:"notice_type"`
	Unread     *int64 `json:"unread"`
}

type ChatSendTextMsgStruct struct {
	ID        uint      `json:"id"`
	Uid       uint      `json:"uid"`
	Username  string    `json:"username"`
	Photo     string    `json:"photo"`
	Tid       uint      `json:"tid"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatUnreadNoticeStruct struct {
	Uid             uint                  `json:"uid"`
	Tid             uint                  `json:"tid"`
	LastMessage     string                `json:"last_message"`
	LastMessageInfo ChatSendTextMsgStruct `json:"last_message_info"`
	Unread          int                   `json:"unread" gorm:"unread"`
}
