package oss

import (
	"Go-Live/global"
	"Go-Live/models/config/uploadMethod"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
