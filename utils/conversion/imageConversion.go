package conversion

import (
	"Go-Live/global"
	"fmt"
)

//FormattingSrc 图片处理相关
func FormattingSrc(src string) string {
	api := global.Config.ProjectUrl
	return fmt.Sprintf("%s/%s", api, src)
}
