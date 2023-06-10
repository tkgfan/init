// Package file
// author gmfan
// date 2023/6/10
package file

import (
	"encoding/json"
	"os"
)

// SaveJson 保存 JSON 实例到 path 中
func SaveJson(path string, data any) (err error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return
	}

	dstFile, err := os.Create(path)
	if err != nil {
		return
	}
	defer dstFile.Close()

	_, err = dstFile.Write(bs)
	if err != nil {
		return
	}
	return
}
