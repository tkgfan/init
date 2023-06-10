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

// InitConfig 舒适化系统配置
func InitConfig() {
	// 创建 config 文件加
	_, err := file.IfNotExistCreateDir(ConfigPath())
	if err != nil {
		logs.Fatal("创建文件", ConfigPath(), "失败，", err)
	}

	// 配置文件不存在则初始化一份默认配置
	fp := ConfigPath() + string(os.PathSeparator) + configFileName
	exist, err := file.PathExist(fp)
	if err != nil {
		logs.Fatal(err)
	}
	if !exist {
		err = file.SaveJson(fp, Config)
		if err != nil {
			logs.Fatal("初始化配置文件失败", err)
		}
	}
}

func ConfigDirName() string {
	return "config"
}

// ConfigPath 配置文件路径
func ConfigPath() string {
	return BasePath() + string(os.PathSeparator) + ConfigDirName()
}
