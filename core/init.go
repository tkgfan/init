// author lby
// date 2023/6/5

package core

import (
	"github.com/tkgfan/init/common/dir"
	"github.com/tkgfan/init/common/logs"
	"github.com/tkgfan/init/conf"
	"os"
	"strings"
)

// ProjectInit 项目初始化
func ProjectInit() {
	// 1 项目已初始化
	if projectHasBeenInitialized() {
		logs.Info("项目已经初始化过了")
		return
	}

	// 2 同步远程模版
	if !conf.OnlyUseLocalTemplate {
		SyncTemplate()
		logs.Info("同步远程模版成功")
	} else {
		logs.Info("仅使用本地模版")
	}

	// 3 执行初始化
	ExecuteInit()
}

func ExecuteInit() {
	// 1 创建项目
	createProject(conf.ProjectName)

	// 2 加载模版数据
	loadTemplate(conf.TemplateName)
}

func loadTemplate(template string) {
	tp := conf.TemplatePath() + "/" + template
	os.Open(tp)
}

func createProject(name string) {
	err := os.MkdirAll(name, os.ModePerm)
	if err != nil {
		logs.Fatal(err)
	}
}

// 项目已经被初始化了
func projectHasBeenInitialized() bool {
	// 目录不存在
	if ok, err := dir.PathExist(conf.ProjectName); !ok || err != nil {
		if err != nil {
			logs.Fatal(err)
		}
		return false
	}

	ds, err := os.ReadDir(conf.ProjectName)
	if err != nil {
		logs.Fatal(err)
	}
	for _, d := range ds {
		if strings.Index(d.Name(), ".") != 0 {
			return true
		}
	}
	return false
}
