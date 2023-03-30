package contribution

import (
	"easy-video-net/consts"
	receive "easy-video-net/interaction/receive/contribution/article"
	response "easy-video-net/interaction/response/contribution/article"
	"easy-video-net/logic/users/notice"
	"easy-video-net/models/common"
	"easy-video-net/models/contribution/article"
	"easy-video-net/models/contribution/article/classification"
	"easy-video-net/models/contribution/article/comments"
	noticeModel "easy-video-net/models/users/notice"
	"easy-video-net/models/users/record"
	"easy-video-net/utils/conversion"
	"encoding/json"
	"fmt"

	"github.com/dlclark/regexp2"
)

func CreateArticleContribution(data *receive.CreateArticleContributionReceiveStruct, uid uint) (results interface{}, err error) {
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
		Uid:                uid,
		ClassificationID:   data.ClassificationID,
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

func UpdateArticleContribution(data *receive.UpdateArticleContributionReceiveStruct, uid uint) (results interface{}, err error) {
	//更新专栏
	articleInfo := new(article.ArticlesContribution)
	if !articleInfo.GetInfoByID(data.ID) {
		return nil, fmt.Errorf("操作视频不存在")
	}
	if articleInfo.Uid != uid {
		return nil, fmt.Errorf("非法操作")
	}
	coverImg, _ := json.Marshal(common.Img{
		Src: data.Cover,
		Tp:  data.CoverUploadType,
	})
	updateList := map[string]interface{}{
		"cover":             coverImg,
		"title":             data.Title,
		"label":             conversion.MapConversionString(data.Label),
		"content":           data.Content,
		"is_comments":       data.Comments,
		"classification_id": data.ClassificationID,
	}
	//进行视频资料更新
	if !articleInfo.Update(updateList) {
		return nil, fmt.Errorf("更新数据失败")
	}
	return "更新成功", nil
}

func DeleteArticleByID(data *receive.DeleteArticleByIDReceiveStruct, uid uint) (results interface{}, err error) {
	al := new(article.ArticlesContribution)
	if !al.Delete(data.ID, uid) {
		return nil, fmt.Errorf("删除失败")
	}
	return "删除成功", nil
}

func GetArticleContributionList(data *receive.GetArticleContributionListReceiveStruct) (results interface{}, err error) {
	articlesContribution := new(article.ArticlesContributionList)
	if !articlesContribution.GetList(data.PageInfo) {
		return nil, fmt.Errorf("查询失败")
	}
	return response.GetArticleContributionListResponse(articlesContribution), nil
}

func GetArticleContributionListByUser(data *receive.GetArticleContributionListByUserReceiveStruct) (results interface{}, err error) {
	articlesContribution := new(article.ArticlesContributionList)
	if !articlesContribution.GetListByUid(data.UserID) {
		return nil, fmt.Errorf("查询失败")
	}
	return response.GetArticleContributionListByUserResponse(articlesContribution), nil
}

func GetArticleContributionByID(data *receive.GetArticleContributionByIDReceiveStruct, uid uint) (results interface{}, err error) {
	articlesContribution := new(article.ArticlesContribution)
	if !articlesContribution.GetInfoByID(data.ArticleID) {
		return nil, fmt.Errorf("查询失败")
	}
	if uid > 0 {
		//添加历史记录
		rd := new(record.Record)
		err = rd.AddArticleRecord(uid, data.ArticleID)
		if err != nil {
			return nil, fmt.Errorf("添加历史记录失败")
		}
	}
	return response.GetArticleContributionByIDResponse(articlesContribution), nil
}

func ArticlePostComment(data *receive.ArticlesPostCommentReceiveStruct, uid uint) (results interface{}, err error) {
	articleInfo := new(article.ArticlesContribution)
	if !articleInfo.GetInfoByID(data.ArticleID) {
		return nil, fmt.Errorf("评论文章不存在")
	}
	ct := comments.Comment{
		PublicModel: common.PublicModel{ID: data.ContentID},
	}
	CommentFirstID := ct.GetCommentFirstID()

	ctu := comments.Comment{
		PublicModel: common.PublicModel{ID: data.ContentID},
	}
	CommentUserID := ctu.GetCommentUserID()
	comment := comments.Comment{
		Uid:            uid,
		ArticleID:      data.ArticleID,
		Context:        data.Content,
		CommentID:      data.ContentID,
		CommentUserID:  CommentUserID,
		CommentFirstID: CommentFirstID,
	}
	if !comment.Create() {
		return nil, fmt.Errorf("发布失败")
	}

	//socket推送(在线的情况下)
	if _, ok := notice.Severe.UserMapChannel[articleInfo.UserInfo.ID]; ok {
		userChannel := notice.Severe.UserMapChannel[articleInfo.UserInfo.ID]
		userChannel.NoticeMessage(noticeModel.ArticleComment)
	}

	return "发布成功", nil
}

func GetArticleComment(data *receive.GetArticleCommentReceiveStruct) (results interface{}, err error) {
	articlesContribution := new(article.ArticlesContribution)
	if !articlesContribution.GetArticleComments(data.ArticleID, data.PageInfo) {
		return nil, fmt.Errorf("查询失败")
	}
	return response.GetArticleContributionCommentsResponse(articlesContribution), nil
}

func GetArticleClassificationList() (results interface{}, err error) {
	cn := new(classification.ClassificationsList)
	err = cn.FindAll()
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	return response.GetArticleClassificationListResponse(cn), nil
}

func GetArticleTotalInfo() (results interface{}, err error) {
	//查询文章数量
	articleNm := new(int64)
	al := new(article.ArticlesContributionList)
	al.GetAllCount(articleNm)
	//查询文章分类信息
	cn := make(classification.ClassificationsList, 0)
	err = cn.FindAll()
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	cnNum := int64(len(cn))
	return response.GetArticleTotalInfoResponse(&cn, articleNm, cnNum), nil
}

func GetArticleManagementList(data *receive.GetArticleManagementListReceiveStruct, uid uint) (results interface{}, err error) {
	//获取个人发布专栏信息
	list := new(article.ArticlesContributionList)
	err = list.GetArticleManagementList(data.PageInfo, uid)
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	res, err := response.GetArticleManagementListResponse(list)
	if err != nil {
		return nil, fmt.Errorf("响应失败")
	}
	return res, nil
}
