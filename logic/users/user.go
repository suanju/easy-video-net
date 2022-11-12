package users

import (
	"Go-Live/global"
	"Go-Live/models/common"
	"Go-Live/models/users"
	"Go-Live/utils/conversion"
	"Go-Live/utils/location"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
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
	index := strings.LastIndex(file.Filename, ".")
	suffix := file.Filename[index:]
	switch suffix {
	case ".jpg", ".jpeg", ".png", ".ico", ".gif", ".wbmp", ".bmp", ".svg", ".webp":
	default:
		return nil, fmt.Errorf("只能上传图片格式嗷！")
	}
	dst := location.AppConfig.ImagePath.UserHeadPortrait + "/userID" + strconv.Itoa(int(userID)) + strconv.Itoa(int(time.Now().UnixNano())) + suffix

	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		global.Logger.Warn("userid %d update headPortrait err", userID)
		return nil, fmt.Errorf("更新失败")
	}
	user := &users.User{PublicModel: common.PublicModel{ID: userID}, Photo: dst}
	if user.Update() {
		return dst, nil
	}
	// 上传文件至指定的完整文件路径
	return nil, fmt.Errorf("更新失败")
}
