package response

type MyCode int64

const (
	CodeDefault MyCode = 0

	CodeSuccess       MyCode = 200
	CodeInvalidParams MyCode = 201
	CodeNoData        MyCode = 202
	CodeServerBusy    MyCode = 500

	CodeInvalidToken      MyCode = 301
	CodeInvalidAuthFormat MyCode = 302
	CodeNotLogin          MyCode = 303

	CodeTypeError MyCode = 415

	CodePasswordMistake MyCode = 1001
)

var msgFlags = map[MyCode]string{
	CodeDefault:       "请求成功",
	CodeSuccess:       "success",
	CodeInvalidParams: "请求参数错误",
	CodeServerBusy:    "服务繁忙",
	CodeNoData:        "没有数据",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",

	CodeTypeError: "类型错误",

	CodePasswordMistake: "密码错误",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
