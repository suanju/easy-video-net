package testing

import (
	"Go-Live/Global"
	"Go-Live/Global/configRead"
	"Go-Live/Global/live"
	"net"
	"time"
)

// LiveSeverTesting 检测直播服务端口
func LiveSeverTesting() {
	//获取直播服务端口地址
	var liveConfig = configRead.Config.LiveConfig

	ipPort := CheckPortsAsLocalHost(liveConfig.IP, []string{"8090", "7001"})
	if len(ipPort) == 0 {
		Global.Logger.Info("开启直播")
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
			Global.Logger.Warn(ipPort, "端口未开启(fail)!")
		} else {
			if conn != nil {
				Global.Logger.Info(ipPort, ipPort, "端口已开启(success)!")
				err := conn.Close()
				if err != nil {
					return nil
				}
			} else {
				Global.Logger.Warn(ipPort, "端口未开启(fail)!")
			}
		}
	}
	return nil
}
