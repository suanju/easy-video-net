package live

import (
	"easy-video-net/global"
	"easy-video-net/utils/location"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	go func() {
		if !global.Config.ProjectConfig.ProjectStates {
			//测试环境快速开启直播服务
			if runtime.GOOS == "windows" {
				err := Start()
				if err != nil {
					global.Logger.Error("开启直播服务失败")
				}
			}
		}
	}()
}

func Start() error {
	dir, err := location.GetCurrentAbPath()
	if err != nil {
		global.Logger.Errorf("测试环境快速开启livego服务失败,获取当前可执行文件目录失败")
		return err
	}
	path := dir + `\Config\live\live-go.exe`
	if _, err := os.Stat(path); os.IsNotExist(err) {
		global.Logger.Errorf("测试环境快速开启livego服务失败 文件 %s 不存在", path)
		return err
	}
	cmd := exec.Command("cmd.exe", "/c", "start "+path)
	if stdio, err := cmd.StdoutPipe(); err != nil {
		return err
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
