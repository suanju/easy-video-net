package commonality

import (
	receive "Go-Live/interaction/receive/commonality"
	response "Go-Live/interaction/response/commonality"
	"Go-Live/models/config/uploadMethod"
	"Go-Live/utils/conversion"
	"Go-Live/utils/oss"
	"fmt"
)

func GetOssConfig(data *receive.GetOssConfigReceiveStruct) (results interface{}, err error) {
	policyToken, err := oss.GetPolicyToken(data.Interface)
	if err != nil {
		return nil, err
	}
	return policyToken, nil
}

func UploadingMethod(data *receive.UploadingMethodStruct) (results interface{}, err error) {
	method := new(uploadMethod.UploadMethod)
	if method.IsExistByField("interface", data.Method) {
		return response.UploadingMethodResponse(method.Method), nil
	} else {
		return nil, fmt.Errorf("未配置上传方法")
	}
}

func GetFullPathOfImage(data *receive.GetFullPathOfImageMethodStruct) (results interface{}, err error) {
	path, err := conversion.SwitchIngStorageFun(data.Type, data.Path)
	if err != nil {
		return nil, err
	}
	return path, nil
}
