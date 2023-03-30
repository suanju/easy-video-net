package logrus

import (
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func init() {

}

var (
	logFilePath = "./runtime/log" //文件存储路径
	logFileName = "system.log"
)

func ReturnsInstance() *logrus.Logger {
	Logger := logrus.New()
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开/写入文件失败", err)
	}
	// 日志级别
	Logger.SetLevel(logrus.DebugLevel)
	// 设置输出
	Logger.Out = file
	// 设置 rotate logs,实现文件分割
	logWriter, err := rotateLogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotateLogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotateLogs.WithMaxAge(7*24*time.Hour), //以hour为单位的整数
		// 设置日志切割时间间隔(1天)
		rotateLogs.WithRotationTime(1*time.Hour),
	)
	// hook机制的设置
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	//给loggers添加hook
	Logger.AddHook(lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return Logger
}
