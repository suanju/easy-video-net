package conversion

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"encoding/json"
	"fmt"
)

//FormattingSrc 图片处理相关
func FormattingSrc(src string) string {
	api := global.Config.ProjectUrl
	return fmt.Sprintf("%s/%s", api, src)
}

func FormattingJsonSrc(str []byte) (url string, err error) {
	data := new(common.Img)
	err = json.Unmarshal(str, data)
	if err != nil {
		return "", fmt.Errorf("json format error")
	}
	path, err := SwitchIngStorageFun(data.Tp, data.Src)
	if err != nil {
		return "", err
	}
	return path, nil
}

//SwitchIngStorageFun 根据类型拼接路径
func SwitchIngStorageFun(tp string, path string) (url string, err error) {
	prefix, err := SwitchTypeAsUrlPrefix(tp)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", prefix, path), nil
}

//SwitchTypeAsUrlPrefix 取url前缀
func SwitchTypeAsUrlPrefix(tp string) (url string, err error) {
	switch tp {
	case "local":
		return global.Config.ProjectUrl, nil
	case "aliyunOss":
		return global.Config.AliyunOss.Host, nil
	case "wx":
		return "", nil
	default:
		return "", fmt.Errorf("undefined format")
	}
}
