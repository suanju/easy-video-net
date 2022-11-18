package liveInfo

import (
	"Go-Live/utils/conversion"
	"fmt"
)

type GetLiveDataResponse struct {
	Img   string `json:"img"`
	Title string `json:"title"`
}

//GetLiveDataResponse 响应设置信息
func (li *LiveInfo) GetLiveDataResponse() (data any, err error) {
	src, errs := conversion.FormattingJsonSrc(li.Img)
	if errs != nil {
		return nil, fmt.Errorf("json format error")
	}
	return GetLiveDataResponse{
		Img:   src,
		Title: li.Title,
	}, nil
}
