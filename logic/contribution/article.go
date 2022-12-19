package contribution

import (
	"Go-Live/consts"
	"Go-Live/models/common"
	"Go-Live/models/contribution/article"
	"Go-Live/utils/conversion"
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
)

func CreateArticleContribution(data *article.CreateArticleContributionReceiveStruct, userID uint) (results interface{}, err error) {
	//进行内容判断
	for _, v := range data.Label {
		vRune := []rune(v) //避免中文占位问题
		if len(vRune) > 7 {
			return nil, fmt.Errorf("标签长度不能大于7位")
		}
	}

	coverImg, _ := json.Marshal(common.Img{
		Src: data.Cover,
		Tp:  data.CoverUploadType,
	})

	//正则匹配替换url
	//取url前缀
	prefix, err := conversion.SwitchTypeAsUrlPrefix(data.ArticleContributionUploadType)
	if err != nil {
		return nil, fmt.Errorf("保存资源方式不存在")
	}
	//正则匹配替换
	reg := regexp2.MustCompile(`(?<=(img[^>]*src="))[^"]*?`+prefix, 0)
	match, err := reg.Replace(data.Content, consts.UrlPrefixSubstitution, -1, -1)
	data.Content = match
	//插入数据
	articlesContribution := article.ArticlesContribution{
		Uid:                userID,
		Title:              data.Title,
		Cover:              coverImg,
		Timing:             conversion.BoolTurnInt8(*data.Timing),
		TimingTime:         data.TimingTime,
		Label:              conversion.MapConversionString(data.Label),
		Content:            data.Content,
		ContentStorageType: data.ArticleContributionUploadType,
		IsComments:         conversion.BoolTurnInt8(*data.Comments),
		Heat:               0,
	}

	if *data.Timing {
		//发布视频后进行的推送相关（待开发）
	}
	if !articlesContribution.Create() {
		return nil, fmt.Errorf("保存失败")
	}
	return "保存成功", nil
}

func GetArticleContributionListByUser(data *article.GetArticleContributionListByUserReceiveStruct, userID uint) (results interface{}, err error) {
	articlesContribution := new(article.ArticlesContributionList)
	if !articlesContribution.GetListByUid(data.UserID) {
		return nil, fmt.Errorf("查询失败")
	}
	return articlesContribution.GetArticleContributionListByUserResponse(), nil
}

func GetArticleContributionByID(data *article.GetArticleContributionByIDReceiveStruct, userID uint) (results interface{}, err error) {
	articlesContribution := new(article.ArticlesContribution)
	if !articlesContribution.GetInfoByID(data.ArticleID) {
		return nil, fmt.Errorf("查询失败")
	}
	return articlesContribution.GetArticleContributionByIDResponse(), nil
}
