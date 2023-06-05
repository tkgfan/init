// author lby
// date 2023/6/5

package core

import (
	"github.com/tkgfan/init/common/errs"
	"github.com/tkgfan/init/conf"
	"log"
	"os"
	"os/exec"
)

// PrepareTemplate 模版不存在则下载模版，存在则同步模版。
func PrepareTemplate() {
	// 检查本地是否存在模版
	path := conf.BasePath() + "/" + conf.TemplatePath()
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		errs.Fatal(err)
	}

	// 模版已存在，同步模版即可
	if !os.IsNotExist(err) && info.Size() != 0 {
		git := exec.Command("git", "pull")
		git.Dir = path

		handleExecInfo(git.Output())
		return
	}

	// 不存在则需要下载模版
	git := exec.Command("git", "clone", conf.RemoteTemplatePath(), conf.TemplatePath())
	git.Dir = conf.BasePath()
	handleExecInfo(git.Output())
}

func handleExecInfo(bs []byte, err error) {
	if err != nil {
		errs.Fatal(err)
	}
	if len(bs) > 0 {
		log.Println(string(bs))
	}
}
