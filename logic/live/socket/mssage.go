package socket

type Response struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	MsgType string `json:"msgType"`
	Sender  uint   `json:"sender"`
	Data    string `json:"data"`
}
type SingleMsg struct {
	Tid     int    `json:"tid"`
	Message string `json:"message"`
}
