package calculate

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func ComputationalPages(size int, pageSize int) int {
	if size < pageSize {
		return 1
	}
	if size%pageSize == 0 {
		return size % pageSize
	} else {
		return size%pageSize + 1
	}
}

func ArrayIsContain[T comparable](items []T, item T) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//GetVideoResolution 计算视频的宽高
func GetVideoResolution(filePath string) (int, int, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=p=0", filePath)
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	res := strings.Split(string(output), ",")
	if len(res) != 2 {
		return 0, 0, fmt.Errorf("获取视频分辨率失败")
	}
	width, err := strconv.Atoi(strings.TrimSpace(res[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("获取视频分辨率失败")
	}
	height, err := strconv.Atoi(strings.TrimSpace(res[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("获取视频分辨率失败")
	}
	return width, height, nil
}
