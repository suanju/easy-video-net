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
	switch data.Tp {
	case "local":
		return fmt.Sprintf("%s/%s", global.Config.ProjectUrl, data.Src), nil
	case "aliyunOss":
		return fmt.Sprintf("%s/%s", "https://eraser-go-live.oss-cn-hangzhou.aliyuncs.com", data.Src), nil
	default:
		return "", fmt.Errorf("undefined format")
	}
}
