package commonality

type UploadingMethodResponse struct {
	Tp string `json:"type"`
}

func (s UploadingMethodStruct) Response(tp string) interface{} {
	return UploadingMethodResponse{
		Tp: tp,
	}
}
