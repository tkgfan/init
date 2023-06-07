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
	basePath = u.HomeDir + "/.init"

	// 文件不存在则创建 init 文件夹
	ok, err := file.PathExist(basePath)
	if err != nil {
		logs.Fatal(err)
	}
	if !ok {
		err = os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			logs.Fatal(err)
		}
	}
}

// BasePath BathPath 默认为用户主目录
func BasePath() string {
	return basePath
}

func TemplatePath() string {
	return BasePath() + "/" + TemplateDirName()
}

func TemplateDirName() string {
	return "template"
}

// RemoteTemplatePath 远程模版路径
func RemoteTemplatePath() string {
	return "git@github.com:tkgfan/template.git"
}
