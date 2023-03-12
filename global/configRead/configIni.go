package configRead

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type ConfigStruct struct {
	SqlConfig     *SqlConfigStruct
	RConfig       *RConfigStruct
	ProjectConfig *ProjectConfigStruct
	LiveConfig    *LiveConfigStruct
	AliyunOss     *AliyunOss
	ProjectUrl    string
}

func init() {
	//避免全局重复导包
	ReturnsInstance()
}

var Config = new(ConfigStruct)
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
type LiveConfigStruct struct {
	IP        string `ini:"ip"`
	Agreement string `ini:"agreement"`
	RTMP      string `ini:"rtmp"`
	FLV       string `ini:"flv"`
	HLS       string `ini:"hls"`
	Api       string `ini:"api"`
}

type ProjectConfigStruct struct {
	UrlStates bool   `ini:"url_states"`
	Url       string `ini:"url"`
	UrlTest   string `ini:"url_test"`
}

type AliyunOss struct {
	AccessKeyId     string `ini:"accessKeyId"`
	AccessKeySecret string `ini:"accessKeySecret"`
	Host            string `ini:"host"`
	CallbackUrl     string `ini:"callbackUrl"`
}

func ReturnsInstance() *ConfigStruct {
	Config.SqlConfig = &SqlConfigStruct{}
	cfg, err = ini.Load("config/config.ini")
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
	if Config.ProjectConfig.UrlStates {
		Config.ProjectUrl = Config.ProjectConfig.Url
	} else {
		Config.ProjectUrl = Config.ProjectConfig.UrlTest
	}

	return Config
}
