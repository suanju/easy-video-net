package commonality

func UploadingMethodResponse(tp string) interface{} {
	type UploadingMethodResponseStruct struct {
		Tp string `json:"type"`
	}
	return UploadingMethodResponseStruct{
		Tp: tp,
	}
}
