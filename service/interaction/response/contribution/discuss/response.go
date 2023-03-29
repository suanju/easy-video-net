package discuss

import (
	articleComments "easy-video-net/models/contribution/article/comments"
	"easy-video-net/models/contribution/video/barrage"
	videoComments "easy-video-net/models/contribution/video/comments"
	"easy-video-net/utils/conversion"
	"time"
)

type GetDiscussVideoListItem struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Photo     string    `json:"photo"`
	Comment   string    `json:"comment"`
	Cover     string    `json:"cover"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type GetDiscussVideoListStruct []GetDiscussVideoListItem

func GetDiscussVideoListResponse(cml *videoComments.CommentList) interface{} {
	list := make(GetDiscussVideoListStruct, 0)
	for _, v := range *cml {
		photo, _ := conversion.FormattingJsonSrc(v.UserInfo.Photo)
		cover, _ := conversion.FormattingJsonSrc(v.VideoInfo.Cover)
		list = append(list, GetDiscussVideoListItem{
			ID:        v.ID,
			Username:  v.UserInfo.Username,
			Photo:     photo,
			Comment:   v.Context,
			Cover:     cover,
			Title:     v.VideoInfo.Title,
			CreatedAt: v.CreatedAt,
		})
	}
	return list
}

type GetDiscussArticleListItem struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Photo     string    `json:"photo"`
	Comment   string    `json:"comment"`
	Cover     string    `json:"cover"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type GetDiscussArticleListStruct []GetDiscussArticleListItem

func GetDiscussArticleListResponse(cml *articleComments.CommentList) interface{} {
	list := make(GetDiscussArticleListStruct, 0)
	for _, v := range *cml {
		photo, _ := conversion.FormattingJsonSrc(v.UserInfo.Photo)
		cover, _ := conversion.FormattingJsonSrc(v.ArticleInfo.Cover)
		list = append(list, GetDiscussArticleListItem{
			ID:        v.ID,
			Username:  v.UserInfo.Username,
			Photo:     photo,
			Comment:   v.Context,
			Cover:     cover,
			Title:     v.ArticleInfo.Title,
			CreatedAt: v.CreatedAt,
		})
	}
	return list
}

type GetDiscussBarrageListItem struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Photo     string    `json:"photo"`
	Comment   string    `json:"comment"`
	Cover     string    `json:"cover"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type GetDiscussBarrageListStruct []GetDiscussBarrageListItem

func GetDiscussBarrageListResponse(cml *barrage.BarragesList) interface{} {
	list := make(GetDiscussBarrageListStruct, 0)
	for _, v := range *cml {
		photo, _ := conversion.FormattingJsonSrc(v.UserInfo.Photo)
		cover, _ := conversion.FormattingJsonSrc(v.VideoInfo.Cover)
		list = append(list, GetDiscussBarrageListItem{
			ID:        v.ID,
			Username:  v.UserInfo.Username,
			Photo:     photo,
			Comment:   v.Text,
			Cover:     cover,
			Title:     v.VideoInfo.Title,
			CreatedAt: v.CreatedAt,
		})
	}
	return list
}
