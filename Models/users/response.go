package users

import "Go-Live/Utils/conversion"

type UserInfoResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Photo    string `json:"photo"`
	Token    string `json:"token"`
}

//UserInfoResponse  生成返回用用户信息结构体
func (us User) UserInfoResponse(token string) UserInfoResponse {
	//判断用户是否为微信用户进行图片处理
	var photo string
	if len(us.Openid) <= 0 {
		photo = conversion.FormattingSrc(us.Photo)
	} else {
		photo = us.Photo
	}
	return UserInfoResponse{
		ID:       us.ID,
		UserName: us.Username,
		Photo:    photo,
		Token:    token,
	}
}
