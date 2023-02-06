package response

import (
	"Go-Live/consts"
	"github.com/gorilla/websocket"
)

type DataWs struct {
	Code    MyCode      `json:"code"`
	Type    string      `json:"type,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func NotLoginWs(ws *websocket.Conn, msg string) {
	rd := &DataWs{
		Code:    CodeNotLogin,
		Message: msg,
		Data:    nil,
	}
	err := ws.WriteJSON(rd)
	if err != nil {
		return
	}
}

func SuccessWs(ws *websocket.Conn, tp string, data interface{}) {
	rd := &DataWs{
		Code:    CodeSuccess,
		Type:    tp,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	err := ws.WriteJSON(rd)
	if err != nil {
		return
	}
}

func ErrorWs(ws *websocket.Conn, msg string) {
	rd := &DataWs{
		Code:    CodeServerBusy,
		Type:    consts.VideoSocketTypeError,
		Message: msg,
		Data:    nil,
	}
	err := ws.WriteJSON(rd)
	if err != nil {
		return
	}
}
