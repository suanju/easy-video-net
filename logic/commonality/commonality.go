package commonality

import (
	"Go-Live/models/commonality"
	commonalityModel "Go-Live/models/commonality"
	"Go-Live/models/config/uploadMethod"
	"Go-Live/utils/oss"
	"fmt"
)

func GetOssConfig(data *commonalityModel.GetOssConfigReceiveStruct) (results interface{}, err error) {
	policyToken, err := oss.GetPolicyToken(data.Interface)
	if err != nil {
		return nil, err
	}
	return policyToken, nil
}

func UploadingMethod(data *commonality.UploadingMethodStruct) (results interface{}, err error) {
	method := new(uploadMethod.UploadMethod)
	if method.IsExistByField("interface", data.Method) {
		return data.Response(method.Method), nil
	} else {
		return nil, fmt.Errorf("未配置上传方法")
	}
}
