package users

import (
	"crypto/md5"
	"easy-video-net/consts"
	"easy-video-net/global"
	receive "easy-video-net/interaction/receive/users"
	response "easy-video-net/interaction/response/users"
	"easy-video-net/models/common"
	userModel "easy-video-net/models/users"
	"easy-video-net/utils/conversion"
	"easy-video-net/utils/email"
	"easy-video-net/utils/jwt"
	"easy-video-net/utils/location"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func WxAuthorization(data *receive.WxAuthorizationReceiveStruct) (results interface{}, err error) {
	/*微信小程序登录 返回值*/
	type WXLoginResp struct {
		OpenID     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionID    string `json:"unionId"`
		ErrCode    int    `json:"errCode"`
		ErrMsg     string `json:"errMsg"`
		Token      string `json:"token"`
		Auth       uint32 `json:"auth"`
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Nickname   string `json:"nickname"`  //微信昵称
		HeadImage  string `json:"headImage"` //微信头像
	}
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, "wxfbd9d7966fc9796c", "92fbe8e2921e00fc3ba68e34d5d0b986", data.Code)
	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		//response.ResponseError(ctx, err.Error())
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		//response.ResponseError(ctx, err.Error())
		return nil, err
	}
	//得到openid进行处理
	users := new(userModel.User)
	if !users.IsExistByField("openid", wxResp.OpenID) {
		//当这个微信没有注册
		photo, _ := json.Marshal(common.Img{
			Src: data.AvatarUrl,
			Tp:  "wx",
		})
		users := userModel.User{
			Username: data.NickName,
			Openid:   wxResp.OpenID,
			Photo:    photo,
		}
		registerRes := users.Create()
		if !registerRes {
			return nil, fmt.Errorf("注册失败")
		}
		//注册token
		tokenString := jwt.NextToken(users.ID)
		src, _ := conversion.FormattingJsonSrc(users.Photo)
		userInfo := response.UserInfoResponseStruct{
			ID:       users.ID,
			UserName: users.Username,
			Photo:    src,
			Token:    tokenString,
		}
		return userInfo, nil
	}
	//已经注册的话直接返回token
	fmt.Printf("查询到的用户id是：%v", users.ID)
	src, _ := conversion.FormattingJsonSrc(users.Photo)
	tokenString := jwt.NextToken(users.ID)
	userInfo := response.UserInfoResponseStruct{
		ID:       users.ID,
		UserName: users.Username,
		Photo:    src,
		Token:    tokenString,
	}
	return userInfo, nil
}

func Register(data *receive.RegisterReceiveStruct) (results interface{}, err error) {
	//判断邮箱是否唯一
	users := new(userModel.User)
	if users.IsExistByField("email", data.Email) {
		return nil, fmt.Errorf("邮箱已被注册")
	}
	//判断验证码是否正确
	verCode, err := global.RedisDb.Get(fmt.Sprintf("%s%s", consts.RegEmailVerCode, data.Email)).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("验证码过期！")
	}

	if verCode != data.VerificationCode {
		return nil, fmt.Errorf("验证码错误")
	}
	//生成密码盐 8 位
	salt := make([]byte, 6)
	for i := range salt {
		salt[i] = jwt.SaltStr[rand.Int63()%int64(len(jwt.SaltStr))]
	}
	password := []byte(fmt.Sprintf("%s%s%s", salt, data.Password, salt))
	passwordMd5 := fmt.Sprintf("%x", md5.Sum(password))
	photo, _ := json.Marshal(common.Img{
		Src: fmt.Sprintf("%s%s%d%s", location.AppConfig.ImagePath.SystemHeadPortrait, "/auto", rand.Intn(10), ".png"),
		Tp:  "local",
	})
	registerData := &userModel.User{
		Email:     data.Email,
		Username:  data.UserName,
		Salt:      string(salt),
		Password:  passwordMd5,
		Photo:     photo,
		BirthDate: time.Now(),
	}
	registerRes := registerData.Create()
	if !registerRes {
		return nil, fmt.Errorf("注册失败")
	}
	//注册token
	tokenString := jwt.NextToken(registerData.ID)
	results = response.UserInfoResponse(registerData, tokenString)

	return results, nil

}

func Login(data *receive.LoginReceiveStruct) (results interface{}, err error) {
	users := new(userModel.User)
	if !users.IsExistByField("username", data.Username) {
		return nil, fmt.Errorf("账号不存在")
	}
	if !users.IfPasswordCorrect(data.Password) {
		return nil, fmt.Errorf("密码错误")
	}
	//注册token
	tokenString := jwt.NextToken(users.ID)
	userInfo := response.UserInfoResponse(users, tokenString)
	return userInfo, nil
}

func SendEmailVerCode(data *receive.SendEmailVerCodeReceiveStruct) (results interface{}, err error) {
	users := new(userModel.User)
	if users.IsExistByField("email", data.Email) {
		return nil, fmt.Errorf("邮箱已被注册")
	}
	//发送方
	mailTo := []string{data.Email}
	// 邮件主题
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000))
	subject := "验证码"
	// 邮件正文
	body := fmt.Sprintf("您正在注册验证码为:%s,5分钟有效,请勿转发他人", code)
	err = email.SendMail(mailTo, subject, body)
	if err != nil {
		global.Logger.Errorf("发送至%s邮箱验证码失败", data.Email)
		return nil, fmt.Errorf("发送失败")
	}
	err = global.RedisDb.Set(fmt.Sprintf("%s%s", consts.RegEmailVerCode, data.Email), code, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return "发送成功", nil
}

func SendEmailVerCodeByForget(data *receive.SendEmailVerCodeReceiveStruct) (results interface{}, err error) {
	//判断用户是否存在
	users := new(userModel.User)
	if !users.IsExistByField("email", data.Email) {
		return nil, fmt.Errorf("该邮箱未注册")
	}
	//发送方
	mailTo := []string{data.Email}
	// 邮件主题
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000))
	subject := "验证码"
	// 邮件正文
	body := fmt.Sprintf("您正在找回密码您的验证码为:%s,5分钟有效,请勿转发他人", code)
	err = email.SendMail(mailTo, subject, body)
	if err != nil {
		return nil, err
	}
	err = global.RedisDb.Set(fmt.Sprintf("%s%s", consts.RegEmailVerCodeByForget, data.Email), code, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return "发送成功", nil
}

func Forget(data *receive.ForgetReceiveStruct) (results interface{}, err error) {
	//判断邮箱是否唯一
	users := new(userModel.User)
	if !users.IsExistByField("email", data.Email) {
		return nil, fmt.Errorf("该账号不存在")
	}
	//判断验证码是否正确
	verCode, err := global.RedisDb.Get(fmt.Sprintf("%s%s", consts.RegEmailVerCodeByForget, data.Email)).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("验证码过期！")
	}

	if verCode != data.VerificationCode {
		return nil, fmt.Errorf("验证码错误")
	}
	//生成密码盐 8 位
	salt := make([]byte, 6)
	for i := range salt {
		salt[i] = jwt.SaltStr[rand.Int63()%int64(len(jwt.SaltStr))]
	}
	password := []byte(fmt.Sprintf("%s%s%s", salt, data.Password, salt))
	passwordMd5 := fmt.Sprintf("%x", md5.Sum(password))

	registerData := userModel.User{
		Salt:     string(salt),
		Password: passwordMd5,
	}
	registerRes := registerData.Update()
	if !registerRes {
		return nil, fmt.Errorf("修改失败")
	}
	return "修改成功", nil
}
