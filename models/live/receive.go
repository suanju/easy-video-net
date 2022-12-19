package live

type GetLiveRoomReceiveStruct struct {
}

type ReqGetRoom struct {
	Status int    `json:"status" binding:"required"`
	Data   string `json:"data" binding:"required"`
}
