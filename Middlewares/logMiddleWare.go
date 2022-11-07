package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func LogMiddleWare() gin.HandlerFunc {
	var (
		logFilePath = "./Runtime/log" //文件存储路径
		logFileName = "system.log"
	)
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开/写入文件失败", err)
		return nil
	}
	// 实例化
	logger := logrus.New()
	// 日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置输出
	logger.Out = file
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
	logger.AddHook(lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	return func(c *gin.Context) {
		c.Next()
		//请求方式
		method := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		// 打印日志
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"client_ip":   clientIP,
			"req_method":  method,
			"req_uri":     reqUrl,
		}).Info()
	}
}
