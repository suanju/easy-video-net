package live

type GetLiveRoomReceiveStruct struct {
}

type ReqGetRoom struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}
