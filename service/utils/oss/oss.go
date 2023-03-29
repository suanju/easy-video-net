package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"easy-video-net/global"
	"easy-video-net/models/config/uploadMethod"
	"encoding/base64"
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"hash"
	"io"
	"time"
)

// 请填写您的AccessKeyId。
var accessKeyId = global.Config.AliyunOss.AccessKeyId

// 请填写您的AccessKeySecret。
var accessKeySecret = global.Config.AliyunOss.AccessKeySecret

// host的格式为 bucket-name.endpoint ，请替换为您的真实信息。
var host = global.Config.AliyunOss.Host

// callbackUrl为 上传回调服务器的URL，请将下面的IP和Port配置为您自己的真实信息。
var callbackUrl = global.Config.AliyunOss.CallbackUrl

var roleArn = global.Config.AliyunOss.RoleArn

var roleSessionName = global.Config.AliyunOss.RoleSessionName

var durationSeconds = global.Config.AliyunOss.DurationSeconds

var endpoint = global.Config.AliyunOss.Endpoint

// 用户上传文件时指定的前缀。
//var uploadDir string = "upload/img/user/liveCover/"
var expireTime int64 = 30

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}
type PolicyToken struct {
	AccessKeyId string `json:"access_id"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

func GetPolicyToken(_interface string) (results interface{}, err error) {
	//获取当前接口对于的储存路径
	method := new(uploadMethod.UploadMethod)
	if !method.IsExistByField("interface", _interface) {
		return nil, fmt.Errorf("上传接口不存在")
	}
	if len(method.Path) == 0 {
		return nil, fmt.Errorf("请联系管理员设置接口保存路径")
	}
	uploadDir := method.Path
	now := time.Now().Unix()
	expireEnd := now + expireTime
	var tokenExpire = getGmtIso8601(expireEnd)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadDir)
	config.Conditions = append(config.Conditions, condition)

	//calculate signature
	result, err := json.Marshal(config)
	debate := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(accessKeySecret))
	_, err = io.WriteString(h, debate)
	if err != nil {
		return PolicyToken{}, nil
	}
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam CallbackParam
	callbackParam.CallbackUrl = callbackUrl
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callbackStr, err := json.Marshal(callbackParam)
	if err != nil {
		fmt.Println("callback json err:", err)
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)

	var policyToken PolicyToken
	policyToken.AccessKeyId = accessKeyId
	policyToken.Host = host
	policyToken.Expire = expireEnd
	policyToken.Signature = signedStr
	policyToken.Directory = uploadDir
	policyToken.Policy = debate
	policyToken.Callback = callbackBase64

	return policyToken, nil
}

// CreateClient
// * 使用AK&SK初始化账号Client
// * @param accessKeyId
// * @param accessKeySecret
// * @return Client
// * @throws Exception
///**
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *sts20150401.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(endpoint)
	_result = &sts20150401.Client{}
	_result, _err = sts20150401.NewClient(config)
	return _result, _err
}

func GteStsInfo() (*sts20150401.AssumeRoleResponseBodyCredentials, error) {
	client, err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		return nil, err
	}
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		RoleArn:         tea.String(roleArn),
		RoleSessionName: tea.String(roleSessionName),
		DurationSeconds: tea.Int64(int64(durationSeconds)),
	}
	runtime := &util.RuntimeOptions{}
	defer func() {
		if r := tea.Recover(recover()); r != nil {
		}
	}()
	// 复制代码运行请自行打印 API 的返回值
	res, err := client.AssumeRoleWithOptions(assumeRoleRequest, runtime)
	if err != nil {
		return nil, err
	}
	if *res.StatusCode != 200 {
		return nil, fmt.Errorf("错误的状态码: %d", res.StatusCode)
	}
	return res.Body.Credentials, nil
}

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
