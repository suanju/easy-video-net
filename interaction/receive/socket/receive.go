package socket

type Receive struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
