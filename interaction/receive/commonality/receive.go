package commonality

type UploadingMethodStruct struct {
	Method string `json:"method"  binding:"required"`
}

type GetOssConfigReceiveStruct struct {
	Interface string `json:"interface"  binding:"required"`
}
type GetFullPathOfImageMethodStruct struct {
	Path string `json:"path"  binding:"required"`
	Type string `json:"type"  binding:"required"`
}
