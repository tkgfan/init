// author lby
// date 2023/6/5

package core

import (
	"github.com/tkgfan/init/common/file"
	"github.com/tkgfan/init/common/logs"
	"github.com/tkgfan/init/conf"
	"io"
	"os"
	"strings"
	"time"
)

// ProjectInit 项目初始化
func ProjectInit() {
	// 1 项目已初始化过了
	if projectHasBeenInitialized() {
		logs.Info("项目已经初始化过了")
		return
	}

	// 2 同步远程模版
	if !conf.OnlyUseLocalTemplate {
		SyncTemplate()
	} else {
		logs.Info("仅使用本地模版")
	}

	// 3 检查模版是否存在
	templatePath := conf.TemplatePath() + string(os.PathSeparator) + conf.TemplateName
	if ok, err := file.PathExist(templatePath); !ok || err != nil {
		if err != nil {
			logs.Fatal(err)
		}
		logs.Fatal("模版", conf.TemplateName, "不存在")
	}

	logs.Info("开始初始化项目,使用", conf.TemplateName, "模版")
	// 4 创建项目目录
	initProjectStart := time.Now()
	err := os.MkdirAll(conf.ProjectName, os.ModePerm)
	if err != nil {
		logs.Fatal(err)
	}

	// 5 拷贝模版到项目中
	copyTemplate(templatePath, templatePath)
	logs.Info("初始化项目成功，耗时:", time.Since(initProjectStart))
}

// 补充 DirEntry 没有绝对路径信息
type pathDirEntry struct {
	// 文件绝对路径
	Path     string
	DirEntry os.DirEntry
}

func toPathDirEntries(basePath string, ds []os.DirEntry) (res []pathDirEntry) {
	for _, d := range ds {
		res = append(res, pathDirEntry{
			Path:     basePath + string(os.PathSeparator) + d.Name(),
			DirEntry: d,
		})
	}
	return
}

func copyTemplate(templatePath, curPath string) {
	ds, err := os.ReadDir(curPath)
	if err != nil {
		logs.Fatal(err)
	}

	pds := toPathDirEntries(curPath, ds)
	for i := 0; i < len(pds); i++ {
		pd := pds[i]
		dstPath := getDstPath(templatePath, pd.Path)
		dstPath = HandlePlaceholderStr(dstPath)

		// 处理文件夹
		if pd.DirEntry.IsDir() {
			// 创建目录
			err = os.MkdirAll(dstPath, os.ModePerm)
			if err != nil {
				logs.Fatal(err)
			}

			copyTemplate(templatePath, pd.Path)
			continue
		}

		// 复制文件
		handleFile(pd.Path, dstPath)
	}
}

func handleFile(srcPath, dstPath string) {
	// 读取模版文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		logs.Fatal(err)
	}
	defer srcFile.Close()
	bs, err := io.ReadAll(srcFile)
	if err != nil {
		logs.Fatal(err)
	}

	bs = HandlePlaceholderBytes(bs)

	// 创建目标文件
	dstFile, err := os.Create(dstPath)
	if err != nil {
		logs.Fatal(err)
	}
	defer dstFile.Close()
	// 复制 src 文件的内容到 dst 中
	_, err = dstFile.Write(bs)
	if err != nil {
		logs.Fatal(err)
	}
}

// 获取目标路径，templatePath 模版路径，srcPath 模版文件路径
func getDstPath(templatePath, srcPath string) string {
	return conf.ProjectName + string(os.PathSeparator) + srcPath[len(templatePath)+1:]
}

// 项目已经被初始化了
func projectHasBeenInitialized() bool {
	// 目录不存在
	if ok, err := file.PathExist(conf.ProjectName); !ok || err != nil {
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
