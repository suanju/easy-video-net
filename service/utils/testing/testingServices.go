package testing

import (
	"easy-video-net/global"
	"easy-video-net/global/config"
	"easy-video-net/global/live"
	"net"
	"time"
)

// LiveSeverTesting 检测直播服务端口
func LiveSeverTesting() {
	//获取直播服务端口地址
	var liveConfig = config.Config.LiveConfig

	ipPort := CheckPortsAsLocalHost(liveConfig.IP, []string{liveConfig.RTMP, liveConfig.FLV, liveConfig.Api, liveConfig.HLS})
	if len(ipPort) == 0 {
		global.Logger.Info("开启直播")
		err := live.Start()
		if err != nil {
			return
		}
	}
}

// CheckPortsAsLocalHost 检测当前主机端口
func CheckPortsAsLocalHost(ip string, Ports []string) []string {
	//未开启端口
	untenablePort := make([]string, 10)
	for _, ipPort := range Ports {
		// 检测端口
		ipPort = ip + ":" + ipPort
		conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
		if err != nil {
			untenablePort = append(untenablePort, ipPort)
			global.Logger.Warnf(ipPort, "直播服务端口:%d 未开启(fail)!", ipPort)
		} else {
			if conn != nil {
				err := conn.Close()
				if err != nil {
					return nil
				}
			} else {
				global.Logger.Warnf(ipPort, "直播服务端口:%d 未开启(fail)!", ipPort)
			}
		}
	}
	return nil
}
