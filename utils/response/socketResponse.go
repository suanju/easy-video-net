package response

import (
	"github.com/gorilla/websocket"
)

func NotLoginWs(ws *websocket.Conn, msg string) {
	rd := &Data{
		Code:    CodeNotLogin,
		Message: msg,
		Data:    nil,
	}
	err := ws.WriteJSON(rd)
	if err != nil {
		return
	}
}

func ErrorWithMsgWs(ws *websocket.Conn, code MyCode, data interface{}) {
	rd := &Data{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	err := ws.WriteJSON(rd)
	if err != nil {
		return
	}
}

func SuccessWs(ws *websocket.Conn, data interface{}) {
	rd := &Data{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	err := ws.WriteJSON(rd)
	if err != nil {
		return
	}
}

func ErrorWs(ws *websocket.Conn, msg string) {
	rd := &Data{
		Code:    CodeServerBusy,
		Message: msg,
		Data:    nil,
	}
	err := ws.WriteJSON(rd)
	if err != nil {
		return
	}
}
