package commonality

import (
	"Go-Live/models/commonality"
	"Go-Live/utils/oss"
	"fmt"
)

func GetOssConfig() (results interface{}, err error) {
	policyToken := oss.GetPolicyToken()
	return policyToken, nil
}
func UploadingMethod(data *commonality.UploadingMethodStruct) (results interface{}, err error) {
	switch data.Method {
	case "liveCover":
		info, err := GetOssConfig()
		if err != nil {
			return nil, fmt.Errorf("获取上传方法失败")
		}
		return data.Response("aliyunOss", info), nil
	default:
		return nil, fmt.Errorf("未配置上传方法")
	}

}
