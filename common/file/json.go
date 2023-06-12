// Package file
// author gmfan
// date 2023/6/10
package file

import (
	"encoding/json"
	"io"
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

// LoadJson 加载 JSON 文件并反序列化到 p 中
func LoadJson(path string, p any) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		return
	}
	return json.Unmarshal(bs, p)
}
