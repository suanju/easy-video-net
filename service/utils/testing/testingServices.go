package testing

import (
	"easy-video-net/global"
	"easy-video-net/global/config"
	_ "easy-video-net/global/live"
	"easy-video-net/utils/location"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func init() {
	LiveSeverTesting()
	AssetsSeverTesting()
	FFmpegSeverTesting()
}

// LiveSeverTesting 检测直播服务端口
func LiveSeverTesting() {
	//获取直播服务端口地址
	var liveConfig = config.Config.LiveConfig
	CheckPortsAsLocalHost(liveConfig.IP, []string{liveConfig.RTMP, liveConfig.FLV, liveConfig.Api, liveConfig.HLS})
}

//AssetsSeverTesting 运行时静态资源存储目录检测
func AssetsSeverTesting() {
	if !location.IsDir("assets") {
		if err := os.MkdirAll("assets", 0775); err != nil {
			log.Fatalf(" 初始化 assets 目录失败 , 错误原因 : %s", err.Error())
			global.Logger.Errorf(" 初始化 assets 目录失败 , 错误原因 : %s", err.Error())
		}
	}
}

//FFmpegSeverTesting ffmpeg状态检测
func FFmpegSeverTesting() {
	cmd := exec.Command("ffmpeg", "-version")
	if err := cmd.Run(); err != nil {
		log.Fatalf("ffmpeg 状态异常请检查后重试")
		global.Logger.Errorf("ffmpeg 状态异常请检查后重试")
	}
}

// CheckPortsAsLocalHost 检测当前主机端口
func CheckPortsAsLocalHost(ip string, Ports []string) []string {
	//未开启端口
	untenablePort := make([]string, 0)
	for _, ipPort := range Ports {
		// 检测端口
		ipPort = ip + ":" + ipPort
		conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
		if err != nil {
			untenablePort = append(untenablePort, ipPort)
			global.Logger.Errorf(ipPort, "直播服务端口:%d 未开启(fail)!", ipPort)
		} else {
			if conn != nil {
				err := conn.Close()
				if err != nil {
					return nil
				}
			} else {
				global.Logger.Errorf(ipPort, "直播服务端口:%d 未开启(fail)!", ipPort)
			}
		}
	}
	return nil
}
