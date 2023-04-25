package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Info struct {
	SqlConfig     *SqlConfigStruct
	RConfig       *RConfigStruct
	EmailConfig   *EmailConfigStruct
	ProjectConfig *ProjectConfigStruct
	LiveConfig    *LiveConfigStruct
	AliyunOss     *AliyunOss
	ProjectUrl    string
}

func init() {
	//避免全局重复导包
	ReturnsInstance()
}

var Config = new(Info)
var cfg *ini.File
var err error

type SqlConfigStruct struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Host     int    `ini:"host"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RConfigStruct struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
}

type EmailConfigStruct struct {
	User string `ini:"user"`
	Pass string `ini:"pass"`
	Host string `ini:"host"`
	Port string `ini:"port"`
}

type LiveConfigStruct struct {
	IP        string `ini:"ip"`
	Agreement string `ini:"agreement"`
	RTMP      string `ini:"rtmp"`
	FLV       string `ini:"flv"`
	HLS       string `ini:"hls"`
	Api       string `ini:"api"`
}

type ProjectConfigStruct struct {
	ProjectStates bool   `ini:"project_states"`
	Url           string `ini:"url"`
	UrlTest       string `ini:"url_test"`
}

type AliyunOss struct {
	Region                   string `ini:"region"`
	Bucket                   string `ini:"bucket"`
	AccessKeyId              string `ini:"accessKeyId"`
	AccessKeySecret          string `ini:"accessKeySecret"`
	Host                     string `ini:"host"`
	Endpoint                 string `ini:"endpoint"`
	RoleArn                  string `ini:"roleArn"`
	RoleSessionName          string `ini:"roleSessionName"`
	DurationSeconds          int    `ini:"durationSeconds"`
	IsOpenTranscoding        bool   `ini:"isOpenTranscoding"`
	TranscodingTemplate360p  string `ini:"transcodingTemplate360p"`
	TranscodingTemplate480p  string `ini:"transcodingTemplate480p"`
	TranscodingTemplate720p  string `ini:"transcodingTemplate720p"`
	TranscodingTemplate1080p string `ini:"transcodingTemplate1080p"`
}

func ReturnsInstance() *Info {
	Config.SqlConfig = &SqlConfigStruct{}
	cfg, err = ini.Load(filepath.ToSlash("/service/config/config.ini"))
	if err != nil {
		fmt.Printf("配置文件不存在,请检查环境: %v \n", err)
		os.Exit(1)
	}

	err = cfg.Section("mysql").MapTo(Config.SqlConfig)
	if err != nil {
		fmt.Printf("Mysql读取配置文件错误: %v \n", err)
		os.Exit(1)
	}
	Config.RConfig = &RConfigStruct{}
	err = cfg.Section("redis").MapTo(Config.RConfig)
	if err != nil {
		fmt.Printf("Redis读取配置文件错误: %v \n", err)
		os.Exit(1)
	}
	Config.EmailConfig = &EmailConfigStruct{}
	err = cfg.Section("email").MapTo(Config.EmailConfig)
	if err != nil {
		fmt.Printf("email读取配置文件错误: %v \n", err)
		os.Exit(1)
	}
	Config.ProjectConfig = &ProjectConfigStruct{}
	err = cfg.Section("project").MapTo(Config.ProjectConfig)
	if err != nil {
		fmt.Printf("Project读取配置文件错误: %v \n", err)
		os.Exit(1)
	}

	Config.LiveConfig = &LiveConfigStruct{}
	err = cfg.Section("live").MapTo(Config.LiveConfig)
	if err != nil {
		fmt.Printf("Live读取配置文件错误: %v \n", err)
		os.Exit(1)
	}

	Config.AliyunOss = &AliyunOss{}
	err = cfg.Section("aliyunOss").MapTo(Config.AliyunOss)
	if err != nil {
		fmt.Printf("Live读取配置文件错误: %v \n", err)
		os.Exit(1)
	}

	//判断是否为正式环境
	if Config.ProjectConfig.ProjectStates {
		Config.ProjectUrl = Config.ProjectConfig.Url
	} else {
		Config.ProjectUrl = Config.ProjectConfig.UrlTest
	}

	return Config
}
