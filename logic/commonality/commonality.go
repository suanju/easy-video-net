package commonality

import (
	"Go-Live/models/commonality"
	"Go-Live/models/config/uploadMethod"
	"Go-Live/utils/oss"
	"fmt"
)

func GetOssConfig() (results interface{}, err error) {
	policyToken := oss.GetPolicyToken()
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
