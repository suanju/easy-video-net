package location

import "path"

type AppConfigStruct struct {
	ImagePath struct {
		SystemHeadPortrait string //系统头像路径
		UserHeadPortrait   string //用户头像路径
	}
}

var AppConfig AppConfigStruct

func init() {
	AppConfig = AppConfigStruct{
		ImagePath: struct {
			SystemHeadPortrait string
			UserHeadPortrait   string
		}{
			SystemHeadPortrait: path.Clean("assets/static/img/users/headPortrait/system"),
			UserHeadPortrait:   path.Clean("assets/static/img/users/headPortrait/uploaded"),
		},
	}
}
