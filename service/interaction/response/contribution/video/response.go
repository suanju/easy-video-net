package response

import (
	"easy-video-net/models/common"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/contribution/video/barrage"
	"easy-video-net/models/users"
	"easy-video-net/utils/conversion"
	"encoding/json"
	"time"
)

//Info 视频信息
type Info struct {
	ID             uint             `json:"id"`
	Uid            uint             `json:"uid" `
	Title          string           `json:"title" `
	Video          string           `json:"video"`
	Cover          string           `json:"cover" `
	VideoDuration  int64            `json:"video_duration"`
	Label          []string         `json:"label"`
	Introduce      string           `json:"introduce"`
	Heat           int              `json:"heat"`
	BarrageNumber  int              `json:"barrageNumber"`
	Comments       commentsInfoList `json:"comments"`
	IsLike         bool             `json:"is_like"`
	IsCollect      bool             `json:"is_collect"`
	CommentsNumber int              `json:"comments_number"`
	CreatorInfo    creatorInfo      `json:"creatorInfo"`
	CreatedAt      time.Time        `json:"created_at"`
}

//创作者信息
type creatorInfo struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	Signature   string `json:"signature"`
	IsAttention bool   `json:"is_attention"`
}

//推荐视频信息
type recommendVideo struct {
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
type RecommendList []recommendVideo

type Response struct {
	VideoInfo     Info          `json:"videoInfo"`
	RecommendList RecommendList `json:"recommendList"`
}

func GetVideoContributionByIDResponse(vc *video.VideosContribution, recommendVideoList *video.VideosContributionList, isAttention bool, isLike bool, isCollect bool) Response {
	//处理视频主要信息
	creatorAvatar, _ := conversion.FormattingJsonSrc(vc.UserInfo.Photo)
	cover, _ := conversion.FormattingJsonSrc(vc.Cover)
	videoSrc, _ := conversion.FormattingJsonSrc(vc.Video)

	//评论
	comments := commentsInfoList{}
	for _, v := range vc.Comments {
		photo, _ := conversion.FormattingJsonSrc(v.UserInfo.Photo)
		commentUser := users.User{}
		commentUser.Find(v.CommentUserID)
		comments = append(comments, &commentsInfo{
			ID:              v.ID,
			CommentID:       v.CommentID,
			CommentFirstID:  v.CommentFirstID,
			CommentUserID:   v.CommentUserID,
			CommentUserName: commentUser.Username,
			CreatedAt:       v.CreatedAt,
			Context:         v.Context,
			Uid:             v.UserInfo.ID,
			Username:        v.UserInfo.Username,
			Photo:           photo,
		})
	}

	commentsList := comments.getChildComment()

	response := Response{
		VideoInfo: Info{
			ID:             vc.ID,
			Uid:            vc.Uid,
			Title:          vc.Title,
			Video:          videoSrc,
			Cover:          cover,
			VideoDuration:  vc.VideoDuration,
			Label:          conversion.StringConversionMap(vc.Label),
			Introduce:      vc.Introduce,
			Heat:           vc.Heat,
			BarrageNumber:  len(vc.Barrage),
			Comments:       commentsList,
			CommentsNumber: len(commentsList),
			IsLike:         isLike,
			IsCollect:      isCollect,
			CreatorInfo: creatorInfo{
				ID:          vc.UserInfo.ID,
				Username:    vc.UserInfo.Username,
				Avatar:      creatorAvatar,
				Signature:   vc.UserInfo.Signature,
				IsAttention: isAttention,
			},
			CreatedAt: vc.CreatedAt,
		},
	}
	//处理推荐视频
	rl := make(RecommendList, 0)
	for _, lk := range *recommendVideoList {
		cover, _ := conversion.FormattingJsonSrc(lk.Cover)
		videoSrc, _ := conversion.FormattingJsonSrc(lk.Video)
		info := recommendVideo{
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
		}
		rl = append(rl, info)
	}
	response.RecommendList = rl
	return response
}

func GetVideoBarrageResponse(list *barrage.BarragesList) interface{} {
	barrageInfoList := make([][]interface{}, 0)
	for _, v := range *list {
		info := make([]interface{}, 0)
		info = append(info, v.Time)
		info = append(info, v.Type)
		info = append(info, v.Color)
		info = append(info, v.Author)
		info = append(info, v.Text)
		barrageInfoList = append(barrageInfoList, info)
	}
	return barrageInfoList
}

//获取视频弹幕响应
type barrageInfo struct {
	Time     int       `json:"time"`
	Text     string    `json:"text"`
	SendTime time.Time `json:"sendTime"`
}

type barrageInfoList []barrageInfo

func GetVideoBarrageListResponse(list *barrage.BarragesList) interface{} {
	barrageList := make(barrageInfoList, 0)
	for _, v := range *list {
		info := barrageInfo{
			Time:     int(v.Time),
			Text:     v.Text,
			SendTime: v.PublicModel.CreatedAt,
		}
		barrageList = append(barrageList, info)
	}
	return barrageList
}

//评论信息
type commentsInfo struct {
	ID              uint             `json:"id"`
	CommentID       uint             `json:"comment_id"`
	CommentFirstID  uint             `json:"comment_first_id"`
	CreatedAt       time.Time        `json:"created_at"`
	Context         string           `json:"context"`
	Uid             uint             `json:"uid"`
	Username        string           `json:"username"`
	Photo           string           `json:"photo"`
	CommentUserID   uint             `json:"comment_user_id"`
	CommentUserName string           `json:"comment_user_name"`
	LowerComments   commentsInfoList `json:"lowerComments"`
}

type commentsInfoList []*commentsInfo

type GetArticleContributionCommentsResponseStruct struct {
	Id             uint             `json:"id"`
	Comments       commentsInfoList `json:"comments"`
	CommentsNumber int              `json:"comments_number"`
}

//得到分级结构
func (l commentsInfoList) getChildComment() commentsInfoList {
	topList := commentsInfoList{}
	for _, v := range l {
		if v.CommentID == 0 {
			//顶层
			topList = append(topList, v)
		}
	}
	return commentsInfoListSecondTree(topList, l)
}

//生成树结构
func commentsInfoListTree(menus commentsInfoList, allData commentsInfoList) commentsInfoList {
	//循环所有一级菜单
	for k, v := range menus {
		//查询所有该菜单下的所有子菜单
		var nodes commentsInfoList //定义子节点目录
		for _, av := range allData {
			if av.CommentID == v.ID {
				nodes = append(nodes, av)
			}
		}
		for kk, _ := range nodes {
			menus[k].LowerComments = append(menus[k].LowerComments, nodes[kk])
		}
		//将刚刚查询出来的子菜单进行递归,查询出三级菜单和四级菜单
		commentsInfoListTree(nodes, allData)
	}
	return menus
}

func commentsInfoListSecondTree(menus commentsInfoList, allData commentsInfoList) commentsInfoList {
	//循环所有一级菜单
	for k, v := range menus {
		//查询所有该菜单下的所有子菜单
		var nodes commentsInfoList //定义子节点目录
		for _, av := range allData {
			if av.CommentFirstID == v.ID {
				nodes = append(nodes, av)
			}
		}
		for kk, _ := range nodes {
			menus[k].LowerComments = append(menus[k].LowerComments, nodes[kk])
		}
		//将刚刚查询出来的子菜单进行递归,查询出三级菜单和四级菜单
		commentsInfoListTree(nodes, allData)
	}
	return menus
}

func GetVideoContributionCommentsResponse(vc *video.VideosContribution) GetArticleContributionCommentsResponseStruct {
	//评论
	comments := commentsInfoList{}
	for _, v := range vc.Comments {
		photo, _ := conversion.FormattingJsonSrc(v.UserInfo.Photo)
		commentUser := users.User{}
		commentUser.Find(v.CommentUserID)
		comments = append(comments, &commentsInfo{
			ID:              v.ID,
			CommentID:       v.CommentID,
			CommentFirstID:  v.CommentFirstID,
			CommentUserID:   v.CommentUserID,
			CommentUserName: commentUser.Username,
			CreatedAt:       v.CreatedAt,
			Context:         v.Context,
			Uid:             v.UserInfo.ID,
			Username:        v.UserInfo.Username,
			Photo:           photo,
		})
	}
	commentsList := comments.getChildComment()
	//输出
	response := GetArticleContributionCommentsResponseStruct{
		Id:             vc.ID,
		Comments:       commentsList,
		CommentsNumber: len(vc.Comments),
	}
	return response
}

type GetVideoManagementListItem struct {
	ID              uint      `json:"id"`
	Uid             uint      `json:"uid" `
	Title           string    `json:"title" `
	Video           string    `json:"video"`
	Cover           string    `json:"cover" `
	Reprinted       bool      `json:"reprinted"`
	CoverUrl        string    `json:"cover_url"`
	CoverUploadType string    `json:"cover_upload_type"`
	VideoDuration   int64     `json:"video_duration"`
	Label           []string  `json:"label"`
	Introduce       string    `json:"introduce"`
	Heat            int       `json:"heat"`
	BarrageNumber   int       `json:"barrageNumber"`
	CommentsNumber  int       `json:"comments_number"`
	CreatedAt       time.Time `json:"created_at"`
}

type GetVideoManagementList []GetVideoManagementListItem

func GetVideoManagementListResponse(vc *video.VideosContributionList) (interface{}, error) {
	list := make(GetVideoManagementList, 0)
	for _, v := range *vc {
		coverJson := new(common.Img)
		_ = json.Unmarshal(v.Cover, coverJson)
		cover, _ := conversion.FormattingJsonSrc(v.Cover)
		videoSrc, _ := conversion.FormattingJsonSrc(v.Video)
		info := GetVideoManagementListItem{
			ID:              v.ID,
			Uid:             v.Uid,
			Title:           v.Title,
			Video:           videoSrc,
			Cover:           cover,
			Reprinted:       conversion.Int8TurnBool(v.Reprinted),
			CoverUploadType: coverJson.Tp,
			CoverUrl:        coverJson.Src,
			VideoDuration:   v.VideoDuration,
			Label:           conversion.StringConversionMap(v.Label),
			Introduce:       v.Introduce,
			Heat:            v.Heat,
			BarrageNumber:   len(v.Barrage),
			CommentsNumber:  len(v.Comments),
			CreatedAt:       v.CreatedAt,
		}
		list = append(list, info)
	}
	return list, nil
}
