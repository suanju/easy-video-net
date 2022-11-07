package conversion

import (
	"Go-Live/Global"
	"fmt"
)

//FormattingSrc 图片处理相关
func FormattingSrc(src string) string {
	api := Global.Config.ProjectUrl
	return fmt.Sprintf("%s%s", api, src)
}
