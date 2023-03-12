package home

import "Go-Live/models/common"

type GetHomeInfoReceiveStruct struct {
	PageInfo common.PageInfo `json:"page_info" binding:"required"`
}
