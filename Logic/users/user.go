package users

import (
	"Go-Live/Models/common"
	"Go-Live/Models/users"
	"Go-Live/Utils/conversion"
)

func GetUserInfo(userID uint) (results interface{}, err error) {
	user := new(users.User)
	user.IsExistByField("id", userID)
	response := user.UserSetInfoResponse()
	return response, nil
}

func SetUserInfo(data *users.SetUserInfoStruct, userID uint) (results interface{}, err error) {
	user := &users.User{
		PublicModel: common.PublicModel{ID: userID},
	}
	update := &users.User{
		Username:  data.Username,
		Gender:    data.Gender,
		BirthDate: data.BirthDate,
		IsVisible: int(conversion.BoolTurnInt8(data.IsVisible)),
		Signature: data.Signature,
	}
	return user.Update(update), nil
}

func DetermineNameExists(data *users.DetermineNameExistsStruct) (results interface{}, err error) {
	user := new(users.User)
	return user.IsExistByField("username", data.Username), nil
}
