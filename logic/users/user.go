package users

import (
	"Go-Live/global"
	receive "Go-Live/interaction/receive/users"
	response "Go-Live/interaction/response/users"
	"Go-Live/models/common"
	"Go-Live/models/config/uploadMethod"
	"Go-Live/models/users"
	"Go-Live/models/users/liveInfo"
	"Go-Live/utils/conversion"
	"Go-Live/utils/location"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"mime/multipart"
	"os"
	"strings"
)

func GetUserInfo(userID uint) (results interface{}, err error) {
	user := new(users.User)
	user.IsExistByField("id", userID)
	res := response.UserSetInfoResponse(user)
	return res, nil
}

func SetUserInfo(data *receive.SetUserInfoStruct, userID uint) (results interface{}, err error) {
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

func DetermineNameExists(data *receive.DetermineNameExistsStruct, userID uint) (results interface{}, err error) {
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
	//如果文件大小超过maxMemory,则使用临时文件来存储multipart/form中文件数据
	err = ctx.Request.ParseMultipartForm(128)
	if err != nil {
		return
	}
	mForm := ctx.Request.MultipartForm
	//上传文件明
	fileNameSlice := mForm.Value["name"]
	var fileName string
	fileName = strings.Join(fileNameSlice, fileName)

	fileInterfaceSlice := mForm.Value["interface"]
	var fileInterface string
	fileInterface = strings.Join(fileInterfaceSlice, fileInterface)

	method := new(uploadMethod.UploadMethod)
	if !method.IsExistByField("interface", fileInterface) {
		return nil, fmt.Errorf("上传接口不存在")
	}
	if len(method.Path) == 0 {
		return nil, fmt.Errorf("请联系管理员设置接口保存路径")
	}
	//取出文件
	_, fileHeader, err := ctx.Request.FormFile("file")
	index := strings.LastIndex(fileHeader.Filename, ".")
	suffix := fileHeader.Filename[index:]
	switch suffix {
	case ".jpg", ".jpeg", ".png", ".ico", ".gif", ".wbmp", ".bmp", ".svg", ".webp", ".mp4":
	default:
		return nil, fmt.Errorf("非法后缀！")
	}
	if !location.IsDir(method.Path) {
		if err = os.MkdirAll(method.Path, 077); err != nil {
			return nil, fmt.Errorf("创建保存路径失败")
		}
	}
	dst := method.Path + "/" + fileName
	err = ctx.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		global.Logger.Warn("userid %d update headPortrait err", userID)
		return nil, fmt.Errorf("上传失败")
	} else {
		return dst, nil
	}
	// 上传文件至指定的完整文件路径
	return nil, fmt.Errorf("上传失败")

}

func UpdateAvatar(data *receive.UpdateAvatarStruct, userID uint) (results interface{}, err error) {
	method := new(uploadMethod.UploadMethod)
	if !method.IsExistByField("interface", data.Interface) {
		return nil, fmt.Errorf("上传接口不存在")
	}
	photo, _ := json.Marshal(common.Img{
		Src: data.ImgUrl,
		Tp:  method.Method,
	})
	user := &users.User{PublicModel: common.PublicModel{ID: userID}, Photo: photo}
	if user.Update() {
		return conversion.FormattingSrc(data.ImgUrl), nil
	} else {
		return nil, fmt.Errorf("更新失败")
	}
}

func GetLiveData(userID uint) (results interface{}, err error) {
	info := new(liveInfo.LiveInfo)
	if info.IsExistByField("uid", userID) {
		results, err = response.GetLiveDataResponse(info)
		if err != nil {
			return nil, fmt.Errorf("获取失败")
		}
		return results, nil
	}
	return common.Img{}, nil
}

func SaveLiveData(data *receive.SaveLiveDataStruct, userID uint) (results interface{}, err error) {
	img, _ := json.Marshal(common.Img{
		Src: data.ImgUrl,
		Tp:  data.Tp,
	})
	info := &liveInfo.LiveInfo{
		Uid:   userID,
		Title: data.Title,
		Img:   datatypes.JSON(img),
	}
	if info.UpdateInfo() {
		return "修改成功", nil
	} else {
		return nil, fmt.Errorf("修改失败")
	}

}
