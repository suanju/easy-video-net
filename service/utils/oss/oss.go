package oss

import (
	"easy-video-net/global"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ice20201109 "github.com/alibabacloud-go/ice-20201109/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

var accessKeyId = global.Config.AliyunOss.AccessKeyId

var accessKeySecret = global.Config.AliyunOss.AccessKeySecret

var roleArn = global.Config.AliyunOss.RoleArn

var roleSessionName = global.Config.AliyunOss.RoleSessionName

var durationSeconds = global.Config.AliyunOss.DurationSeconds

var endpoint = global.Config.AliyunOss.Endpoint

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

func CreateStsClient(accessKeyId *string, accessKeySecret *string) (_result *sts20150401.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(endpoint)
	_result = &sts20150401.Client{}
	_result, _err = sts20150401.NewClient(config)
	return _result, _err
}

func GteStsInfo() (*sts20150401.AssumeRoleResponseBodyCredentials, error) {
	client, err := CreateStsClient(tea.String(accessKeyId), tea.String(accessKeySecret))
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
	res, err := client.AssumeRoleWithOptions(assumeRoleRequest, runtime)
	if err != nil {
		return nil, err
	}
	if *res.StatusCode != 200 {
		return nil, fmt.Errorf("错误的状态码: %d", res.StatusCode)
	}
	return res.Body.Credentials, nil
}

func CreateIceClient(accessKeyId *string, accessKeySecret *string) (_result *ice20201109.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ice.cn-hangzhou.aliyuncs.com")
	result := &ice20201109.Client{}
	result, err := ice20201109.NewClient(config)
	return result, err
}
func RegisterMediaInfo(inputUrl, mediaType, Title string) (body *ice20201109.RegisterMediaInfoResponseBody, err error) {
	client, err := CreateIceClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		global.Logger.Errorf("初始化cilent失败 err : %s", err.Error())
	}
	registerMediaInfoRequest := &ice20201109.RegisterMediaInfoRequest{
		Overwrite: tea.Bool(true),
		InputURL:  tea.String(inputUrl),
		MediaType: tea.String(mediaType),
		Title:     tea.String(Title),
	}
	runtime := &util.RuntimeOptions{}
	defer func() {
		if r := tea.Recover(recover()); r != nil {
		}
	}()
	result, _err := client.RegisterMediaInfoWithOptions(registerMediaInfoRequest, runtime)
	if _err != nil {
		global.Logger.Errorf("注册媒资失败 err %s ", err.Error())
		fmt.Println(_err)
	}
	if *result.StatusCode != 200 {
		return nil, err
	}
	return result.Body, nil
}

func GetMediaInfo(mediaID *string) (body *ice20201109.GetMediaInfoResponse, err error) {
	client, err := CreateIceClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		global.Logger.Errorf("初始化cilent失败 err : %s", err.Error())
	}

	getMediaInfoRequest := &ice20201109.GetMediaInfoRequest{
		MediaId: mediaID,
	}
	runtime := &util.RuntimeOptions{}
	if r := tea.Recover(recover()); r != nil {
	}
	result, err := client.GetMediaInfoWithOptions(getMediaInfoRequest, runtime)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SubmitTranscodeJob(taskName, mediaID, outputUrl, template string) (body *ice20201109.SubmitTranscodeJobResponseBody, err error) {
	client, err := CreateIceClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		global.Logger.Errorf("初始化cilent失败 err : %s", err.Error())
	}
	inputGroup0 := &ice20201109.SubmitTranscodeJobRequestInputGroup{
		Type:  tea.String("Media"),
		Media: tea.String(mediaID),
	}
	outputGroup0Output := &ice20201109.SubmitTranscodeJobRequestOutputGroupOutput{
		Type:  tea.String("OSS"),
		Media: tea.String(outputUrl),
	}
	outputGroup0ProcessConfigTranscode := &ice20201109.SubmitTranscodeJobRequestOutputGroupProcessConfigTranscode{
		TemplateId: tea.String(template),
	}
	outputGroup0ProcessConfig := &ice20201109.SubmitTranscodeJobRequestOutputGroupProcessConfig{
		Transcode: outputGroup0ProcessConfigTranscode,
	}
	outputGroup0 := &ice20201109.SubmitTranscodeJobRequestOutputGroup{
		ProcessConfig: outputGroup0ProcessConfig,
		Output:        outputGroup0Output,
	}
	submitTranscodeJobRequest := &ice20201109.SubmitTranscodeJobRequest{
		OutputGroup: []*ice20201109.SubmitTranscodeJobRequestOutputGroup{outputGroup0},
		Name:        tea.String(taskName),
		InputGroup:  []*ice20201109.SubmitTranscodeJobRequestInputGroup{inputGroup0},
	}
	runtime := &util.RuntimeOptions{}
	defer func() {
		if r := tea.Recover(recover()); r != nil {
		}
	}()
	result, err := client.SubmitTranscodeJobWithOptions(submitTranscodeJobRequest, runtime)
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}

func GetTranscodeJob(taskID string) (body *ice20201109.GetTranscodeJobResponseBody, err error) {
	client, err := CreateIceClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		global.Logger.Errorf("初始化cilent失败 err : %s", err.Error())
	}
	getTranscodeJobRequest := &ice20201109.GetTranscodeJobRequest{
		ParentJobId: tea.String("9ce776d01f034d23b31bc68ffbb2e276"),
	}
	runtime := &util.RuntimeOptions{}
	result, err := client.GetTranscodeJobWithOptions(getTranscodeJobRequest, runtime)
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}
