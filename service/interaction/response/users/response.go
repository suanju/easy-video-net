package response

import (
	"easy-video-net/models/common"
	"easy-video-net/models/contribution/article"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/users"
	"easy-video-net/models/users/attention"
	"easy-video-net/models/users/chat/chatList"
	"easy-video-net/models/users/chat/chatMsg"
	"easy-video-net/models/users/collect"
	"easy-video-net/models/users/favorites"
	"easy-video-net/models/users/liveInfo"
	"easy-video-net/models/users/notice"
	"easy-video-net/models/users/record"
	"easy-video-net/utils/conversion"
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"time"
)

type UserInfoResponseStruct struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Photo     string    `json:"photo"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

//UserInfoResponse  生成返回用用户信息结构体
func UserInfoResponse(us *users.User, token string) UserInfoResponseStruct {
	//判断用户是否为微信用户进行图片处理
	photo, _ := conversion.FormattingJsonSrc(us.Photo)
	return UserInfoResponseStruct{
		ID:        us.ID,
		UserName:  us.Username,
		Photo:     photo,
		Token:     token,
		CreatedAt: us.CreatedAt,
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
		Gender:    us.Gender,
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
		if len(label) >= 2 {
			label = label[:1]
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

type GetFavoritesInfo struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Cover    string `json:"cover"`
	Tp       string `json:"type"`
	Src      string `json:"src"`
	Max      int    `json:"max"`
	UsesInfo struct {
		Username string `json:"username"`
	} `json:"userInfo"`
}

type GetFavoritesInfoList []GetFavoritesInfo

func GetFavoritesListResponse(al *favorites.FavoriteList) (data interface{}, err error) {
	list := make(GetFavoritesInfoList, 0)
	for _, v := range *al {
		coverInfo := new(common.Img)
		err = json.Unmarshal(v.Cover, coverInfo)

		cover, _ := conversion.FormattingJsonSrc(v.Cover)
		list = append(list, GetFavoritesInfo{
			ID:      v.ID,
			Title:   v.Title,
			Content: v.Content,
			Cover:   cover,
			Tp:      coverInfo.Tp,
			Src:     coverInfo.Src,
			Max:     v.Max,
			UsesInfo: struct {
				Username string `json:"username"`
			}{
				Username: v.UserInfo.Username,
			},
		})
	}
	return list, nil
}

type GetFavoritesListByFavoriteVideoInfo struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Cover    string `json:"cover"`
	Tp       string `json:"type"`
	Src      string `json:"src"`
	Max      int    `json:"max"`
	Selected bool   `json:"selected"`
	Present  int    `json:"present"`
	UsesInfo struct {
		Username string `json:"username"`
	} `json:"userInfo"`
}

type GetFavoritesListByFavoriteVideoInfoList []GetFavoritesListByFavoriteVideoInfo

func GetFavoritesListByFavoriteVideoResponse(al *favorites.FavoriteList, ids []uint) (data interface{}, err error) {
	list := make(GetFavoritesListByFavoriteVideoInfoList, 0)
	for _, v := range *al {
		coverInfo := new(common.Img)
		err = json.Unmarshal(v.Cover, coverInfo)
		cover, _ := conversion.FormattingJsonSrc(v.Cover)
		//判断是否已选
		selected := false
		for _, vv := range ids {
			if vv == v.ID {
				selected = true
			}
		}

		list = append(list, GetFavoritesListByFavoriteVideoInfo{
			ID:       v.ID,
			Title:    v.Title,
			Content:  v.Content,
			Cover:    cover,
			Tp:       coverInfo.Tp,
			Src:      coverInfo.Src,
			Max:      v.Max,
			Selected: selected,
			Present:  len(v.CollectList),
			UsesInfo: struct {
				Username string `json:"username"`
			}{
				Username: v.UserInfo.Username,
			},
		})
	}
	return list, nil
}

type GetFavoriteVideoListItem struct {
	ID            uint      `json:"id"`
	Uid           uint      `json:"uid"`
	Title         string    `json:"title"`
	Video         string    `json:"video"`
	Cover         string    `json:"cover"`
	VideoDuration int64     `json:"video_duration"`
	CreatedAt     time.Time `json:"created_at"`
}

type GetFavoriteVideoList []GetFavoriteVideoListItem

type GetFavoriteVideoListResponseStruct struct {
	VideoList GetFavoriteVideoList `json:"videoList"`
}

func GetFavoriteVideoListResponse(cl *collect.CollectsList) (data interface{}, err error) {
	//处理视频
	vl := make(GetFavoriteVideoList, 0)
	for _, ck := range *cl {
		lk := ck.VideoInfo
		cover, _ := conversion.FormattingJsonSrc(lk.Cover)
		videoSrc, _ := conversion.FormattingJsonSrc(lk.Video)
		vl = append(vl, GetFavoriteVideoListItem{
			ID:            lk.ID,
			Uid:           lk.Uid,
			Title:         lk.Title,
			VideoDuration: lk.VideoDuration,
			Video:         videoSrc,
			Cover:         cover,
			CreatedAt:     ck.CreatedAt,
		})
	}
	return GetFavoriteVideoListResponseStruct{VideoList: vl}, nil
}

type GetRecordListItem struct {
	ID        uint      `json:"id"`
	ToID      uint      `json:"to_id"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Username  string    `json:"username"`
	Photo     string    `json:"photo"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetRecordListItemList []GetRecordListItem

func GetRecordListResponse(rl *record.RecordsList) (data interface{}, err error) {
	list := make(GetRecordListItemList, 0)
	for _, v := range *rl {
		var cover string
		var photo string
		var title string
		var username string
		var tp string
		if v.Type == "video" {
			cover, _ = conversion.FormattingJsonSrc(v.VideoInfo.Cover)
			photo, _ = conversion.FormattingJsonSrc(v.VideoInfo.UserInfo.Photo)
			title = v.VideoInfo.Title
			username = v.VideoInfo.UserInfo.Username
			tp = "视频"
		} else if v.Type == "article" {
			cover, _ = conversion.FormattingJsonSrc(v.ArticleInfo.Cover)
			photo, _ = conversion.FormattingJsonSrc(v.ArticleInfo.UserInfo.Photo)
			title = v.ArticleInfo.Title
			username = v.ArticleInfo.UserInfo.Username
			tp = "专栏"
		} else {
			cover, _ = conversion.FormattingJsonSrc(v.Userinfo.LiveInfo.Img)
			photo, _ = conversion.FormattingJsonSrc(v.Userinfo.Photo)
			title = v.Userinfo.LiveInfo.Title
			username = v.Userinfo.Username
			tp = "直播"
		}
		list = append(list, GetRecordListItem{
			ID:        v.ID,
			ToID:      v.ToId,
			Title:     title,
			Cover:     cover,
			Username:  username,
			Photo:     photo,
			Type:      tp,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return list, nil
}

type GetNoticeListItem struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Type      string    `json:"type"`
	ToID      uint      `json:"to_id"`
	Photo     string    `json:"photo"`
	Comment   string    `json:"comment"`
	Cover     string    `json:"cover"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type GetNoticeListStruct []GetNoticeListItem

func GetNoticeListResponse(nl *notice.NoticesList) (data interface{}, err error) {
	list := make(GetNoticeListStruct, 0)
	for _, v := range *nl {
		photo, _ := conversion.FormattingJsonSrc(v.UserInfo.Photo)
		var cover string
		var title string

		//判断类型确定标题和封面
		switch v.Type {
		case notice.VideoComment:
			cover, _ = conversion.FormattingJsonSrc(v.VideoInfo.Cover)
			title = v.VideoInfo.Title
			break
		case notice.VideoLike:
			cover, _ = conversion.FormattingJsonSrc(v.VideoInfo.Cover)
			title = v.VideoInfo.Title
			break
		case notice.ArticleComment:
			cover, _ = conversion.FormattingJsonSrc(v.ArticleInfo.Cover)
			title = v.ArticleInfo.Title
			break
		case notice.ArticleLike:
			cover, _ = conversion.FormattingJsonSrc(v.ArticleInfo.Cover)
			title = v.ArticleInfo.Title
			break
		}

		list = append(list, GetNoticeListItem{
			ID:        v.ID,
			Type:      v.Type,
			ToID:      v.ToID,
			Username:  v.UserInfo.Username,
			Photo:     photo,
			Comment:   v.Content,
			Cover:     cover,
			Title:     title,
			CreatedAt: v.CreatedAt,
		})

	}

	return list, nil
}

type ChatMessageInfo struct {
	ID        uint      `json:"id"`
	Uid       uint      `json:"uid"`
	Username  string    `json:"username"`
	Photo     string    `json:"photo"`
	Tid       uint      `json:"tid"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type GetChatListItem struct {
	ToID            uint              `json:"to_id"`
	Username        string            `json:"username"`
	Photo           string            `json:"photo"`
	Unread          int               `json:"unread" gorm:"unread"`
	LastMessage     string            `json:"last_message"`
	LastMessagePage int               `json:"last_message_page"`
	MessageList     []ChatMessageInfo `json:"message_list"`
	LastAt          time.Time         `json:"last_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

type GetChatListResponseStruct []GetChatListItem

func GetChatListResponse(chatList *chatList.ChatList, msgList map[uint]*chatMsg.MsgList) (data interface{}, err error) {
	list := make(GetChatListResponseStruct, 0)
	for _, v := range *chatList {
		photo, _ := conversion.FormattingJsonSrc(v.ToUserInfo.Photo)
		messageList := make([]ChatMessageInfo, 0)
		for _, vv := range *msgList[v.Tid] {
			uPhoto, _ := conversion.FormattingJsonSrc(vv.UInfo.Photo)
			messageList = append(messageList, ChatMessageInfo{
				ID:        vv.ID,
				Uid:       vv.Uid,
				Username:  vv.UInfo.Username,
				Photo:     uPhoto,
				Tid:       vv.Tid,
				Message:   vv.Message,
				Type:      vv.Type,
				CreatedAt: vv.CreatedAt,
			})
		}
		list = append(list, GetChatListItem{
			ToID:        v.Tid,
			Username:    v.ToUserInfo.Username,
			Photo:       photo,
			Unread:      v.Unread,
			LastMessage: v.LastMessage,
			MessageList: messageList,
			LastAt:      v.LastAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return list, nil
}

func GetChatHistoryMsgResponse(list *chatMsg.MsgList) (data interface{}, err error) {
	messageList := make([]ChatMessageInfo, 0)
	for _, v := range *list {
		photo, _ := conversion.FormattingJsonSrc(v.UInfo.Photo)
		messageList = append(messageList, ChatMessageInfo{
			ID:        v.ID,
			Uid:       v.Uid,
			Username:  v.UInfo.Username,
			Photo:     photo,
			Tid:       v.Tid,
			Message:   v.Message,
			Type:      v.Type,
			CreatedAt: v.CreatedAt,
		})
	}
	return messageList, nil
}
