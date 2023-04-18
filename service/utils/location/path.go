package location

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//GetCurrentAbPath 最终方案-全兼容
func GetCurrentAbPath() (dir string, err error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("获取目录失败")
	}
	dir = filepath.Dir(filename)
	rootDir := filepath.Join(dir, "..", "..")
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(err)
		return "", fmt.Errorf("获取目录失败")
	}
	return absRootDir, nil
}

// IsDir 判断是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir 创建目录
func CreateDir(dirName string) bool {
	err := os.Mkdir(dirName, 755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
