// Package conf
// author gmfan
// date 2023/6/10
package conf

import (
	"github.com/tkgfan/init/common/file"
	"github.com/tkgfan/init/common/logs"
	"os"
)

type config struct {
	Repository string `json:"repository"`
}

// Config 系统配置
var Config = &config{
	Repository: "git@github.com:tkgfan/template.git",
}

const configFileName = "config.json"

// InitConfig 初始化系统配置
func InitConfig() {
	// 创建 config 文件加
	_, err := file.IfNotExistCreateDir(ConfigBasePath())
	if err != nil {
		logs.Fatal("创建文件", ConfigBasePath(), "失败，", err)
	}

	// 配置文件不存在则初始化一份默认配置
	exist, err := file.PathExist(ConfigFilePath())
	if err != nil {
		logs.Fatal(err)
	}
	if exist {
		err = file.LoadJson(ConfigFilePath(), Config)
		if err != nil {
			logs.Fatal("加载配置文件失败", err)
		}
	} else {
		err = file.SaveJson(ConfigFilePath(), Config)
		if err != nil {
			logs.Fatal("初始化配置文件失败", err)
		}
	}
}

// FlushConfig 刷新配置
func FlushConfig() (err error) {
	return file.SaveJson(ConfigFilePath(), Config)
}

func ConfigBasePath() string {
	return BasePath() + string(os.PathSeparator) + "config"
}

// ConfigFilePath 配置文件路径
func ConfigFilePath() string {
	return ConfigBasePath() + string(os.PathSeparator) + configFileName
}
