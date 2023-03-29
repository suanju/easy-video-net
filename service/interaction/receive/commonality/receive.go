package commonality

import "easy-video-net/models/common"

type UploadingMethodStruct struct {
	Method string `json:"method"  binding:"required"`
}
type UploadingDirStruct struct {
	Interface string `json:"interface"  binding:"required"`
}

type GetFullPathOfImageMethodStruct struct {
	Path string `json:"path"  binding:"required"`
	Type string `json:"type"  binding:"required"`
}

type SearchStruct struct {
	PageInfo common.PageInfo `json:"page_info" binding:"required"`
	Type     string          `json:"type" binding:"required"`
}

type UploadSliceInfo struct {
	Index int    `json:"index" `
	Hash  string `json:"hash"`
}
type UploadSliceList []UploadSliceInfo

type UploadCheckStruct struct {
	FileMd5   string          `json:"file_md5" binding:"required"`
	Interface string          `json:"interface" binding:"required"`
	SliceList UploadSliceList `json:"slice_list"  binding:"required"`
}

type UploadMergeStruct struct {
	FileName  string          `json:"file_name" binding:"required"`
	Interface string          `json:"interface" binding:"required"`
	SliceList UploadSliceList `json:"slice_list"  binding:"required"`
}
