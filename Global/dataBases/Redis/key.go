package Redis

type keyNameStruct struct {
	//邮箱验证码
	RegEmailVerCode         string
	RegEmailVerCodeByForget string
}

func initRedisKey() {
	KeyName = keyNameStruct{
		RegEmailVerCode:         "RegEmailVerCode_",         //注册验证码
		RegEmailVerCodeByForget: "RegEmailVerCodeByForget_", //找回密码验证码
	}
}
