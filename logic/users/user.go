package users

import (
	"Go-Live/models/common"
	"Go-Live/models/users"
	"Go-Live/utils/conversion"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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
	update := map[string]interface{}{
		"Username":  data.Username,
		"Gender":    0,
		"BirthDate": data.BirthDate,
		"IsVisible": conversion.BoolTurnInt8(*data.IsVisible),
		"Signature": data.Signature,
	}

	return user.UpdatePureZero(update), nil
}

func DetermineNameExists(data *users.DetermineNameExistsStruct, userID uint) (results interface{}, err error) {
	user := new(users.User)
	is := user.IsExistByField("username", data.Username)
	//判断是否未更改
	if user.ID == userID {
		return false, nil
	} else if is {
		return true, nil
	} else {
		return false, nil
	}
}

func Upload(file *multipart.FileHeader, userID uint, ctx *gin.Context) (results interface{}, err error) {
	dst := "./assets/static/img/users/uploaded/" + file.Filename
	// 上传文件至指定的完整文件路径
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
