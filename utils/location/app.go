package location

import "path"

type AppConfigStruct struct {
	ImagePath struct {
		SystemHeadPortrait string //系统头像路径
		UserHeadPortrait   string //用户头像路径
		LiveCover          string //直播封面
	}
}

var AppConfig *AppConfigStruct

func init() {
	AppConfig = &AppConfigStruct{
		ImagePath: struct {
			SystemHeadPortrait string
			UserHeadPortrait   string
			LiveCover          string
		}{
			SystemHeadPortrait: path.Clean("assets/static/img/users/headPortrait/system"),
			UserHeadPortrait:   path.Clean("assets/static/img/users/headPortrait/uploaded"),
			LiveCover:          path.Clean("assets/static/img/users/liveCover"),
		},
	}
}
