// author gmfan
// date 2023/6/5

package conf

import (
	"github.com/tkgfan/init/common/file"
	"github.com/tkgfan/init/common/logs"
	"os"
	"os/user"
)

var basePath string

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	basePath = u.HomeDir + string(os.PathSeparator) + ".init"

	// 文件不存在则创建 init 文件夹
	_, err = file.IfNotExistCreateDir(basePath)
	if err != nil {
		logs.Fatal("创建文件", basePath, "失败，", err)
	}

	// 初始化配置文件
	InitConfig()
}

// BasePath BathPath 默认为用户主目录
func BasePath() string {
	return basePath
}

func TemplatePath() string {
	return BasePath() + string(os.PathSeparator) + TemplateDirName()
}

func TemplateDirName() string {
	return "template"
}

// RemoteTemplatePath 远程模版路径
func RemoteTemplatePath() string {
	return Config.Repository
}
