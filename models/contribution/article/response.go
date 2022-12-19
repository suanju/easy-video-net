package article

import (
	"Go-Live/consts"
	"Go-Live/utils/conversion"
	"fmt"
	"github.com/dlclark/regexp2"
	"time"
)

type GetArticleContributionListByUserResponse struct {
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

type GetArticleContributionByIDResponse struct {
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

type GetArticleContributionListByUserResponseList []GetArticleContributionListByUserResponse

func (l ArticlesContributionList) GetArticleContributionListByUserResponse() GetArticleContributionListByUserResponseList {
	response := make(GetArticleContributionListByUserResponseList, 0)
	for _, v := range l {
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

		response = append(response, GetArticleContributionListByUserResponse{
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
func (vc ArticlesContribution) GetArticleContributionByIDResponse() GetArticleContributionByIDResponse {
	coverSrc, _ := conversion.FormattingJsonSrc(vc.Cover)

	prefix, _ := conversion.SwitchTypeAsUrlPrefix(vc.ContentStorageType)
	//正则替换src
	reg := regexp2.MustCompile(`(?<=(img[^>]*src="))[^"]*?`+consts.UrlPrefixSubstitutionEscape, 0)
	fmt.Printf("%v", reg)
	match, _ := reg.Replace(vc.Content, prefix, -1, -1)
	vc.Content = match

	//只显示两个标签
	label := conversion.StringConversionMap(vc.Label)

	response := GetArticleContributionByIDResponse{
		Id:             vc.ID,
		Uid:            vc.Uid,
		Title:          vc.Title,
		Cover:          coverSrc,
		Label:          label,
		Content:        vc.Content,
		IsComments:     conversion.Int8TurnBool(vc.IsComments),
		Heat:           vc.Heat,
		LikesNumber:    len(vc.Likes),
		CommentsNumber: len(vc.Comments),
		CreatedAt:      vc.CreatedAt,
	}
	return response
}
