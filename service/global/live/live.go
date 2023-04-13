package live

import (
	"easy-video-net/global"
	"easy-video-net/utils/location"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

func init() {
	go func() {
		if !global.Config.ProjectConfig.ProjectStates {
			//测试环境快速开启直播服务
			err := Start()
			if err != nil {
				global.Logger.Error("开启直播服务失败")
			}
		}
	}()
}

func Start() error {
	path := location.GetCurrentAbPath()
	path = path + `\Config\live\`
	cmd := exec.Command("cmd.exe", "/c", "start "+path+"live-go.exe")
	//获取输出对象，可以从该对象中读取输出结果
	if stdio, err := cmd.StdoutPipe(); err != nil {
		return err
		log.Fatal(err)
	} else {
		// 保证关闭输出流
		defer func(stdio io.ReadCloser) {
			err := stdio.Close()
			if err != nil {

			}
		}(stdio)
		// 运行命令
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		if _, err := ioutil.ReadAll(stdio); err != nil {
			// 读取输出结果
			log.Fatal(err)
			return err
		} else {
		}
	}
	return nil
}
