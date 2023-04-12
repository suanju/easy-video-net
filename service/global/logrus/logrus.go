package logrus

import (
	"encoding/json"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type JsonInfo struct {
	Time     string `json:"time"`
	Level    string `json:"level"`
	Msg      string `json:"msg"`
	File     string `json:"file,omitempty"`
	Function string `json:"function,omitempty"`
}

//JsonFormatter 自定义json 解析
type JsonFormatter struct {
	logrus.JSONFormatter
}

func (f *JsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 构造 JSON 数据
	info := &JsonInfo{
		Time:  entry.Time.Format("2006.01.02 15:04:05"),
		Level: entry.Level.String(),
		Msg:   entry.Message,
	}
	//只有等级匹配进行打印调用者信息
	if entry.Level == logrus.ErrorLevel || entry.Level == logrus.WarnLevel || entry.Level == logrus.DebugLevel || entry.Level == logrus.PanicLevel {
		info.File = entry.Caller.File
		info.Function = entry.Caller.Function
	}
	lineBreak := "\n"
	jsonData, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	formattedMsg := string(jsonData) + lineBreak
	return []byte(formattedMsg), nil
}

var (
	logFilePath = "./runtime/log" //文件存储路径
)

func ReturnsInstance() *logrus.Logger {
	Logger := logrus.New()
	// 日志级别
	Logger.SetLevel(logrus.DebugLevel)
	//打印调用者信息
	Logger.SetReportCaller(true)
	//定义到空输出
	Logger.SetOutput(ioutil.Discard)
	// 设置 rotate logs,实现文件分割
	logInfoWriter, _ := rotateLogs.New(
		logFilePath+"/%Y-%m-%d/info.log",
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(1*time.Hour),
	)
	logFataWriter, _ := rotateLogs.New(
		logFilePath+"/%Y-%m-%d/fata.log",
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(1*time.Hour),
	)
	logDebugWriter, _ := rotateLogs.New(
		logFilePath+"/%Y-%m-%d/debug.log",
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(1*time.Hour),
	)
	logWarnWriter, _ := rotateLogs.New(
		logFilePath+"/%Y-%m-%d/warn.log",
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(1*time.Hour),
	)
	logErrorWriter, _ := rotateLogs.New(
		logFilePath+"/%Y-%m-%d/error.log",
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(1*time.Hour),
	)
	logPanicWriter, _ := rotateLogs.New(
		logFilePath+"/%Y-%m-%d/panic.log",
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(1*time.Hour),
	)

	// hook机制的设置
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logInfoWriter,
		logrus.FatalLevel: logFataWriter,
		logrus.DebugLevel: logDebugWriter,
		logrus.WarnLevel:  logWarnWriter,
		logrus.ErrorLevel: logErrorWriter,
		logrus.PanicLevel: logPanicWriter,
	}
	Logger.Formatter = &JsonFormatter{}
	//给loggers添加hook
	Logger.AddHook(lfshook.NewHook(writerMap, &JsonFormatter{}))

	return Logger
}
