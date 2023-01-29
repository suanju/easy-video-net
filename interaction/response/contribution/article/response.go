package response

import (
	"Go-Live/consts"
	"Go-Live/models/contribution/article"
	"Go-Live/models/users"
	"Go-Live/utils/conversion"
	"github.com/dlclark/regexp2"
	"time"
)

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

type GetArticleContributionListByUserResponseStruct struct {
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
	CreatedAt      time.Time `json:"created_at"`
}

type GetArticleContributionByIDResponseStruct struct {
	Id             uint             `json:"id"`
	Uid            uint             `json:"uid" `
	Title          string           `json:"title" `
	Cover          string           `json:"cover" `
	Label          []string         `json:"label" `
	Content        string           `json:"content"`
	IsComments     bool             `json:"is_comments"`
	Heat           int              `json:"heat"`
	LikesNumber    int              `json:"likes_number"`
	Comments       commentsInfoList `json:"comments"`
	CommentsNumber int              `json:"comments_number"`
	CreatedAt      time.Time        `json:"created_at"`
}

type GetArticleContributionListByUserResponseList []GetArticleContributionListByUserResponseStruct

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

func GetArticleContributionListByUserResponse(l *article.ArticlesContributionList) GetArticleContributionListByUserResponseList {
	response := make(GetArticleContributionListByUserResponseList, 0)
	for _, v := range *l {
		coverSrc, _ := conversion.FormattingJsonSrc(v.Cover)

		//正则替换首文内容
		reg := regexp2.MustCompile(`<(\S*?)[^>]*>.*?|<.*? />`, 0)
		match, _ := reg.Replace(v.Content, "", -1, -1)
		matchRune := []rune(match)
		if len(matchRune) > 70 {
			v.Content = string(matchRune[:70]) + "..."
		} else {
			v.Content = match
		}

		//只显示两个标签
		label := conversion.StringConversionMap(v.Label)
		if len(label) >= 3 {
			label = label[:2]
		}

		response = append(response, GetArticleContributionListByUserResponseStruct{
			Id:             v.ID,
			Uid:            v.Uid,
			Title:          v.Title,
			Cover:          coverSrc,
			Label:          label,
			Content:        v.Content,
			IsComments:     conversion.Int8TurnBool(v.IsComments),
			Heat:           v.Heat,
			LikesNumber:    len(v.Likes),
			CommentsNumber: len(v.Comments),
			CreatedAt:      v.CreatedAt,
		})
	}
	return response
}
func GetArticleContributionByIDResponse(vc *article.ArticlesContribution) GetArticleContributionByIDResponseStruct {
	coverSrc, _ := conversion.FormattingJsonSrc(vc.Cover)

	prefix, _ := conversion.SwitchTypeAsUrlPrefix(vc.ContentStorageType)
	//正则替换src
	reg := regexp2.MustCompile(`(?<=(img[^>]*src="))[^"]*?`+consts.UrlPrefixSubstitutionEscape, 0)
	match, _ := reg.Replace(vc.Content, prefix, -1, -1)
	vc.Content = match

	//只显示两个标签
	label := conversion.StringConversionMap(vc.Label)

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
	response := GetArticleContributionByIDResponseStruct{
		Id:             vc.ID,
		Uid:            vc.Uid,
		Title:          vc.Title,
		Cover:          coverSrc,
		Label:          label,
		Content:        vc.Content,
		IsComments:     conversion.Int8TurnBool(vc.IsComments),
		Heat:           vc.Heat,
		LikesNumber:    len(vc.Likes),
		Comments:       commentsList,
		CommentsNumber: len(vc.Comments),
		CreatedAt:      vc.CreatedAt,
	}
	return response
}
func GetArticleContributionCommentsResponse(vc *article.ArticlesContribution) GetArticleContributionCommentsResponseStruct {
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
