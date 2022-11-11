package location

type AppConfigStruct struct {
	ImagePath struct {
		SystemHeadPortrait string //系统头像路径
	}
}

var AppConfig AppConfigStruct

func init() {
	AppConfig = AppConfigStruct{
		ImagePath: struct {
			SystemHeadPortrait string
		}{
			SystemHeadPortrait: "/assets/static/img/users/headPortrait/system",
		},
	}
}
