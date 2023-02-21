package response

import (
	"Go-Live/models/contribution/article"
	"Go-Live/models/contribution/video"
	"Go-Live/models/users"
	"Go-Live/models/users/attention"
	"Go-Live/models/users/liveInfo"
	"Go-Live/utils/conversion"
	"fmt"
	"github.com/dlclark/regexp2"
	"time"
)

type UserInfoResponseStruct struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Photo    string `json:"photo"`
	Token    string `json:"token"`
}

//UserInfoResponse  生成返回用用户信息结构体
func UserInfoResponse(us *users.User, token string) UserInfoResponseStruct {
	//判断用户是否为微信用户进行图片处理
	photo, _ := conversion.FormattingJsonSrc(us.Photo)
	return UserInfoResponseStruct{
		ID:       us.ID,
		UserName: us.Username,
		Photo:    photo,
		Token:    token,
	}
}

type UserSetInfoResponseStruct struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Gender    int8      `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	IsVisible bool      `json:"is_visible"`
	Signature string    `json:"signature"`
}

//UserSetInfoResponse  设置用户信息返回结果
func UserSetInfoResponse(us *users.User) UserSetInfoResponseStruct {
	return UserSetInfoResponseStruct{
		ID:        us.ID,
		UserName:  us.Username,
		Gender:    int8(us.Gender),
		BirthDate: us.BirthDate,
		IsVisible: conversion.IntTurnBool(int(us.IsVisible)),
		Signature: us.Signature,
	}
}

type GetLiveDataResponseStruct struct {
	Img   string `json:"img"`
	Title string `json:"title"`
}

//GetLiveDataResponse 响应设置信息
func GetLiveDataResponse(li *liveInfo.LiveInfo) (data any, err error) {
	src, errs := conversion.FormattingJsonSrc(li.Img)
	if errs != nil {
		return nil, fmt.Errorf("json format error")
	}
	return GetLiveDataResponseStruct{
		Img:   src,
		Title: li.Title,
	}, nil
}

type GetSpaceIndividualResponseStruct struct {
	ID            uint   `json:"id"`
	UserName      string `json:"username"`
	Photo         string `json:"photo"`
	Signature     string `json:"signature"`
	IsAttention   bool   `json:"is_attention"`
	AttentionNum  *int64 `json:"attention_num"`
	VermicelliNum *int64 `json:"vermicelli_num"`
}

//GetSpaceIndividualResponse 获取空间信息
func GetSpaceIndividualResponse(us *users.User, isAttention bool, attentionNum *int64, vermicelliNum *int64) (data any, err error) {
	photo, _ := conversion.FormattingJsonSrc(us.Photo)
	return GetSpaceIndividualResponseStruct{
		ID:            us.ID,
		UserName:      us.Username,
		Photo:         photo,
		Signature:     us.Signature,
		IsAttention:   isAttention,
		AttentionNum:  attentionNum,
		VermicelliNum: vermicelliNum,
	}, nil
}

type ReleaseInformationVideoInfo struct {
	ID            uint      `json:"id"`
	Uid           uint      `json:"uid" `
	Title         string    `json:"title" `
	Video         string    `json:"video"`
	Cover         string    `json:"cover" `
	VideoDuration int64     `json:"video_duration"`
	Label         []string  `json:"label"`
	Introduce     string    `json:"introduce"`
	Heat          int       `json:"heat"`
	BarrageNumber int       `json:"barrageNumber"`
	Username      string    `json:"username"`
	CreatedAt     time.Time `json:"created_at"`
}

type ReleaseInformationVideoInfoList []ReleaseInformationVideoInfo

type ReleaseInformationArticleInfo struct {
	Id             uint      `json:"id"`
	Uid            uint      `json:"uid" `
	Title          string    `json:"title" `
	Cover          string    `json:"cover" `
	Label          []string  `json:"label" `
	Content        string    `json:"content"`
	IsComments     bool      `json:"is_comments"`
	Heat           int       `json:"heat"`
	LikesNumber    int       `json:"likes_number"`
	CommentsNumber int       `json:"comments_number"`
	Classification string    `json:"classification"`
	CreatedAt      time.Time `json:"created_at"`
}

type ReleaseInformationArticleInfoList []ReleaseInformationArticleInfo

type GetReleaseInformationResponseStruct struct {
	VideoList   ReleaseInformationVideoInfoList   `json:"videoList"`
	ArticleList ReleaseInformationArticleInfoList `json:"articleList"`
}

//GetReleaseInformationResponse 获取视频专栏发布信息
func GetReleaseInformationResponse(videoList *video.VideosContributionList, articleList *article.ArticlesContributionList) (data interface{}, err error) {
	//处理视频
	vl := make(ReleaseInformationVideoInfoList, 0)
	for _, lk := range *videoList {
		cover, _ := conversion.FormattingJsonSrc(lk.Cover)
		videoSrc, _ := conversion.FormattingJsonSrc(lk.Video)
		vl = append(vl, ReleaseInformationVideoInfo{
			ID:            lk.ID,
			Uid:           lk.Uid,
			Title:         lk.Title,
			Video:         videoSrc,
			Cover:         cover,
			VideoDuration: lk.VideoDuration,
			Label:         conversion.StringConversionMap(lk.Label),
			Introduce:     lk.Introduce,
			Heat:          lk.Heat,
			BarrageNumber: len(lk.Barrage),
			Username:      lk.UserInfo.Username,
			CreatedAt:     lk.CreatedAt,
		})
	}

	//处理专栏
	al := make(ReleaseInformationArticleInfoList, 0)
	for _, v := range *articleList {
		coverSrc, _ := conversion.FormattingJsonSrc(v.Cover)

		//正则替换首文内容
		reg := regexp2.MustCompile(`<(\S*?)[^>]*>.*?|<.*? />`, 0)
		match, _ := reg.Replace(v.Content, "", -1, -1)
		matchRune := []rune(match)
		if len(matchRune) > 100 {
			v.Content = string(matchRune[:100]) + "..."
		} else {
			v.Content = match
		}

		//只显示两个标签
		label := conversion.StringConversionMap(v.Label)
		if len(label) >= 3 {
			label = label[:2]
		}
		al = append(al, ReleaseInformationArticleInfo{
			Id:             v.ID,
			Uid:            v.Uid,
			Title:          v.Title,
			Cover:          coverSrc,
			Label:          label,
			Content:        v.Content,
			Classification: v.Classification.Label,
			IsComments:     conversion.Int8TurnBool(v.IsComments),
			Heat:           v.Heat,
			LikesNumber:    len(v.Likes),
			CommentsNumber: len(v.Comments),
			CreatedAt:      v.CreatedAt,
		})
	}
	return GetReleaseInformationResponseStruct{
		VideoList:   vl,
		ArticleList: al,
	}, nil
}

type GetAttentionListInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Signature   string `json:"signature"`
	Photo       string `json:"photo"`
	IsAttention bool   `json:"is_attention"`
}

type GetAttentionListInfoList []GetAttentionListInfo

func GetAttentionListResponse(al *attention.AttentionsList, arr []uint) (data interface{}, err error) {
	list := make(GetAttentionListInfoList, 0)
	for _, v := range *al {
		photo, _ := conversion.FormattingJsonSrc(v.AttentionUserInfo.Photo)
		isAttention := false
		for _, ak := range arr {
			if ak == v.AttentionID {
				isAttention = true
			}
		}
		list = append(list, GetAttentionListInfo{
			ID:          v.AttentionID,
			Name:        v.AttentionUserInfo.Username,
			Signature:   v.AttentionUserInfo.Signature,
			Photo:       photo,
			IsAttention: isAttention,
		})
	}
	return list, nil
}

type GetVermicelliListInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Signature   string `json:"signature"`
	Photo       string `json:"photo"`
	IsAttention bool   `json:"is_attention"`
}

type GetVermicelliListInfoList []GetVermicelliListInfo

func GetVermicelliListResponse(al *attention.AttentionsList, arr []uint) (data interface{}, err error) {
	list := make(GetVermicelliListInfoList, 0)
	for _, v := range *al {
		photo, _ := conversion.FormattingJsonSrc(v.UserInfo.Photo)
		isAttention := false
		for _, ak := range arr {
			if ak == v.Uid {
				isAttention = true
			}
		}
		list = append(list, GetVermicelliListInfo{
			ID:          v.Uid,
			Name:        v.UserInfo.Username,
			Signature:   v.UserInfo.Signature,
			Photo:       photo,
			IsAttention: isAttention,
		})
	}
	return list, nil
}
