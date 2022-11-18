package commonality

type UploadingMethodResponse struct {
	Tp   string      `json:"type"`
	Info interface{} `json:"info"`
}

func (s UploadingMethodStruct) Response(tp string, info interface{}) interface{} {
	return UploadingMethodResponse{
		Tp:   tp,
		Info: info,
	}
}
