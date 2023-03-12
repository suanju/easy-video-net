package users

import (
	"Go-Live/consts"
	"Go-Live/global"
	receive "Go-Live/interaction/receive/users"
	response "Go-Live/interaction/response/users"
	"Go-Live/models/common"
	"Go-Live/models/config/uploadMethod"
	"Go-Live/models/users"
	"Go-Live/models/users/attention"
	"Go-Live/models/users/chat/chatList"
	"Go-Live/models/users/chat/chatMsg"
	"Go-Live/models/users/collect"
	"Go-Live/models/users/favorites"
	"Go-Live/models/users/liveInfo"
	"Go-Live/models/users/notice"
	"Go-Live/models/users/record"
	"Go-Live/utils/conversion"
	"Go-Live/utils/email"
	"Go-Live/utils/jwt"
	"Go-Live/utils/location"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/datatypes"
)

func GetUserInfo(uid uint) (results interface{}, err error) {
	user := new(users.User)
	user.IsExistByField("id", uid)
	res := response.UserSetInfoResponse(user)
	return res, nil
}

func SetUserInfo(data *receive.SetUserInfoReceiveStruct, uid uint) (results interface{}, err error) {
	user := &users.User{
		PublicModel: common.PublicModel{ID: uid},
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

func DetermineNameExists(data *receive.DetermineNameExistsStruct, uid uint) (results interface{}, err error) {
	user := new(users.User)
	is := user.IsExistByField("username", data.Username)
	//判断是否未更改
	if user.ID == uid {
		return false, nil
	} else if is {
		return true, nil
	} else {
		return false, nil
	}
}

func Upload(file *multipart.FileHeader, uid uint, ctx *gin.Context) (results interface{}, err error) {
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
		global.Logger.Warn("userid %d update headPortrait err", uid)
		return nil, fmt.Errorf("上传失败")
	} else {
		return dst, nil
	}
	// 上传文件至指定的完整文件路径

}

func UpdateAvatar(data *receive.UpdateAvatarStruct, uid uint) (results interface{}, err error) {
	photo, _ := json.Marshal(common.Img{
		Src: data.ImgUrl,
		Tp:  data.Tp,
	})
	user := &users.User{PublicModel: common.PublicModel{ID: uid}, Photo: photo}
	if user.Update() {
		return conversion.FormattingSrc(data.ImgUrl), nil
	} else {
		return nil, fmt.Errorf("更新失败")
	}
}

func GetLiveData(uid uint) (results interface{}, err error) {
	info := new(liveInfo.LiveInfo)
	if info.IsExistByField("uid", uid) {
		results, err = response.GetLiveDataResponse(info)
		if err != nil {
			return nil, fmt.Errorf("获取失败")
		}
		return results, nil
	}
	return common.Img{}, nil
}

func SaveLiveData(data *receive.SaveLiveDataReceiveStruct, uid uint) (results interface{}, err error) {
	img, _ := json.Marshal(common.Img{
		Src: data.ImgUrl,
		Tp:  data.Tp,
	})
	info := &liveInfo.LiveInfo{
		Uid:   uid,
		Title: data.Title,
		Img:   datatypes.JSON(img),
	}
	if info.UpdateInfo() {
		return "修改成功", nil
	} else {
		return nil, fmt.Errorf("修改失败")
	}

}

func SendEmailVerificationCodeByChangePassword(uid uint) (results interface{}, err error) {
	user := new(users.User)
	user.Find(uid)
	//发送方
	mailTo := []string{user.Email}
	// 邮件主题
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000))
	subject := "验证码"
	// 邮件正文
	body := fmt.Sprintf("您正在修改密码,您的验证码为:%s,5分钟有效,请勿转发他人", code)
	err = email.SendMail(mailTo, subject, body)
	if err != nil {
		return nil, err
	}
	err = global.RedisDb.Set(fmt.Sprintf("%s%s", consts.EmailVerificationCodeByChangePassword, user.Email), code, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return "发送成功", nil

}

func ChangePassword(data *receive.ChangePasswordReceiveStruct, uid uint) (results interface{}, err error) {
	user := new(users.User)
	user.Find(uid)

	if data.Password != data.ConfirmPassword {
		return nil, fmt.Errorf("两次密码不一致！")
	}

	//判断验证码是否正确
	verCode, err := global.RedisDb.Get(fmt.Sprintf("%s%s", consts.EmailVerificationCodeByChangePassword, user.Email)).Result()
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

	user.Salt = string(salt)
	user.Password = passwordMd5

	registerRes := user.Update()
	if !registerRes {
		return nil, fmt.Errorf("修改失败")
	}
	return "修改成功", nil
}

func Attention(data *receive.AttentionReceiveStruct, uid uint) (results interface{}, err error) {
	at := new(attention.Attention)
	if at.Attention(uid, data.Uid) {
		if data.Uid == uid {
			return nil, fmt.Errorf("操作失败")
		}
		return "操作成功", nil
	}
	return nil, fmt.Errorf("操作失败")
}

func CreateFavorites(data *receive.CreateFavoritesReceiveStruct, uid uint) (results interface{}, err error) {
	if data.ID == 0 {
		//插入模式
		if len(data.Title) == 0 {
			return nil, fmt.Errorf("标题为空")
		}
		//判断是否只有标题
		if data.ID <= 0 && len(data.Tp) == 0 && len(data.Content) == 0 && len(data.Cover) == 0 {
			//单标题创建
			fs := &favorites.Favorites{Uid: uid, Title: data.Title, Max: 1000}
			if !fs.Create() {
				return nil, fmt.Errorf("创建失败")
			}
			return fmt.Errorf("创建成功"), nil
		} else {
			//资料齐全创建
			cover, _ := json.Marshal(common.Img{
				Src: data.Cover,
				Tp:  data.Tp,
			})
			fs := &favorites.Favorites{
				Uid:     uid,
				Title:   data.Title,
				Content: data.Content,
				Cover:   cover,
				Max:     1000,
			}
			if !fs.Create() {
				return nil, fmt.Errorf("创建失败")
			}
			return fmt.Errorf("创建成功"), nil
		}
	} else {
		//进行更新
		fs := new(favorites.Favorites)
		if !fs.Find(data.ID) {
			return nil, fmt.Errorf("查询失败")
		}
		if fs.Uid != uid {
			return nil, fmt.Errorf("查询非法操作")
		}
		cover, _ := json.Marshal(common.Img{
			Src: data.Cover,
			Tp:  data.Tp,
		})
		fs.Title = data.Title
		fs.Content = data.Content
		fs.Cover = cover
		if !fs.Update() {
			return nil, fmt.Errorf("更新失败")
		}
		return "更新成功", nil
	}
}

func GetFavoritesList(uid uint) (results interface{}, err error) {
	fl := new(favorites.FavoriteList)
	err = fl.GetFavoritesList(uid)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	res, err := response.GetFavoritesListResponse(fl)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteFavorites(data *receive.DeleteFavoritesReceiveStruct, uid uint) (results interface{}, err error) {
	fs := new(favorites.Favorites)
	err = fs.Delete(data.ID, uid)
	if err != nil {
		return nil, err
	}
	return "删除成功", nil
}

func FavoriteVideo(data *receive.FavoriteVideoReceiveStruct, uid uint) (results interface{}, err error) {
	for _, k := range data.IDs {
		fs := new(favorites.Favorites)
		fs.Find(k)
		if fs.Uid != uid {
			return nil, fmt.Errorf("非法操作")
		}
		if len(fs.CollectList)+1 > fs.Max {
			return nil, fmt.Errorf("收藏夹已满")
		}

		cl := &collect.Collect{
			Uid:         uid,
			FavoritesID: k,
			VideoID:     data.VideoID,
		}
		if !cl.Create() {
			return nil, fmt.Errorf("收藏失败")
		}
	}
	return "操作成功", nil
}

func GetFavoritesListByFavoriteVideo(data *receive.GetFavoritesListByFavoriteVideoReceiveStruct, uid uint) (results interface{}, err error) {
	//获取收藏夹列表
	fl := new(favorites.FavoriteList)
	err = fl.GetFavoritesList(uid)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	//查询该视频在那些收藏夹内已收藏
	cl := new(collect.CollectsList)
	err = cl.FindVideoExistWhere(data.VideoID)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	ids := make([]uint, 0)
	for _, v := range *cl {
		ids = append(ids, v.FavoritesID)
	}

	res, err := response.GetFavoritesListByFavoriteVideoResponse(fl, ids)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetFavoriteVideoList(data *receive.GetFavoriteVideoListReceiveStruct) (results interface{}, err error) {
	cl := new(collect.CollectsList)
	err = cl.GetVideoInfo(data.FavoriteID)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	res, err := response.GetFavoriteVideoListResponse(cl)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetRecordList(data *receive.GetRecordListReceiveStruct, uid uint) (results interface{}, err error) {
	rl := new(record.RecordsList)
	err = rl.GetRecordListByUid(uid, data.PageInfo)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	res, err := response.GetRecordListResponse(rl)
	if err != nil {
		return nil, fmt.Errorf("响应失败")
	}
	return res, nil
}

func ClearRecord(uid uint) (results interface{}, err error) {
	rl := new(record.Record)
	err = rl.ClearRecord(uid)
	if err != nil {
		return nil, fmt.Errorf("清空失败")
	}
	return "清空完成", nil
}

func DeleteRecordByID(data *receive.DeleteRecordByIDReceiveStruct, uid uint) (results interface{}, err error) {
	rl := new(record.Record)
	err = rl.DeleteRecordByID(data.ID, uid)
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return "删除成功", nil
}

func GetNoticeList(data *receive.GetNoticeListReceiveStruct, uid uint) (results interface{}, err error) {
	//获取用户通知
	messageType := make([]string, 0)
	nl := new(notice.NoticesList)
	switch data.Type {
	case "comment":
		messageType = append(messageType, notice.VideoComment, notice.ArticleComment)
		break
	case "like":
		messageType = append(messageType, notice.VideoLike, notice.ArticleLike)
	}

	err = nl.GetNoticeList(data.PageInfo, messageType, uid)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	//记录全部已读
	n := new(notice.Notice)
	err = n.ReadAll(uid)
	if err != nil {
		return nil, fmt.Errorf("已读消息失败")
	}
	res, err := response.GetNoticeListResponse(nl)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetChatList(uid uint) (results interface{}, err error) {
	//获取消息列表
	cList := new(chatList.ChatList)
	err = cList.GetListByIO(uid)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	ids := make([]uint, 0)
	for _, v := range *cList {
		ids = append(ids, v.Tid)
	}
	msgList := make(map[uint]*chatMsg.MsgList, 0)
	for _, v := range ids {
		ml := new(chatMsg.MsgList)
		err = ml.FindList(uid, v)
		if err != nil {
			break
		}
		msgList[v] = ml
	}
	res, err := response.GetChatListResponse(cList, msgList)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PersonalLetter(data *receive.PersonalLetterReceiveStruct, uid uint) (results interface{}, err error) {
	//获取消息列表
	cm := new(chatMsg.Msg)
	err = cm.GetLastMessage(data.ID, uid)
	if err != nil {
		return nil, fmt.Errorf("操作失败")
	}

	ci := &chatList.ChatsListInfo{
		Uid:         uid,
		Tid:         data.ID,
		LastMessage: cm.Message,
	}
	err = ci.AddChat()
	if err != nil {
		return nil, fmt.Errorf("操作失败")
	}
	return cm, nil
}

func DeleteChatItem(data *receive.DeleteChatItemReceiveStruct, uid uint) (results interface{}, err error) {
	ci := new(chatList.ChatsListInfo)
	fmt.Println("执行了")
	err = ci.DeleteChat(data.ID, uid)
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return "操作成功", nil
}
