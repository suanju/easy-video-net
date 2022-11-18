package liveInfo

//SaveLiveDataStruct 设置直播信息
type SaveLiveDataStruct struct {
	Tp     string `json:"type" binding:"required"`
	ImgUrl string `json:"imgUrl" binding:"required"`
	Title  string `json:"title" binding:"required"`
}
