// author gmfan
// date 2023/6/5

package core

import (
	"github.com/tkgfan/init/common/logs"
	"github.com/tkgfan/init/conf"
	"os"
	"os/exec"
)

// SyncTemplate 同步模版，同步模版不存在则下载模版，存在则同步模版。
func SyncTemplate() {
	// 检查本地是否存在模版
	info, err := os.Stat(conf.TemplatePath())
	if err != nil && !os.IsNotExist(err) {
		logs.Fatal(err)
	}

	// 模版已存在，同步模版即可
	if !os.IsNotExist(err) && info.Size() != 0 {
		git := exec.Command("git", "pull")
		git.Dir = conf.TemplatePath()

		handleExecInfo(git.Output())
		return
	}

	// 不存在则需要下载模版
	git := exec.Command("git", "clone", conf.RemoteTemplatePath(), conf.TemplateDirName())
	git.Dir = conf.BasePath()
	handleExecInfo(git.Output())
}

func handleExecInfo(bs []byte, err error) {
	if err != nil {
		logs.Info(string(bs))
		logs.Fatal(err)
	}
}
