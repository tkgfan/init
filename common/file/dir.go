// author gmfan
// date 2023/6/5

package file

import (
	"os"
)

// PathExist 判断路径是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IfNotExistCreateDir path 路径不存在则创建对应的文件夹
func IfNotExistCreateDir(path string) (exist bool, err error) {
	// 文件不存在则创建 init 文件夹
	exist, err = PathExist(path)
	if err != nil {
		return
	}
	if !exist {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return
}
