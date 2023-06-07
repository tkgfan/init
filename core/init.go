// author lby
// date 2023/6/5

package core

import (
	"fmt"
	"github.com/tkgfan/init/common/file"
	"github.com/tkgfan/init/common/logs"
	"github.com/tkgfan/init/conf"
	"io"
	"os"
	"strings"
	"sync"
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
		logs.Info("同步远程模版成功")
	} else {
		logs.Info("仅使用本地模版")
	}

	// 3 检查模版是否存在
	templatePath := conf.TemplatePath() + "/" + conf.TemplateName
	if ok, err := file.PathExist(templatePath); !ok || err != nil {
		if err != nil {
			logs.Fatal(err)
		}
		logs.Fatal("模版", conf.TemplateName, "不存在")
	}

	// 4 创建项目目录
	err := os.MkdirAll(conf.ProjectName, os.ModePerm)
	if err != nil {
		logs.Fatal(err)
	}

	// 5 拷贝模版到项目中
	t1 := time.Now().UnixMilli()
	copyTemplate(templatePath, templatePath)
	wg.Wait()
	fmt.Println("耗时", time.Now().UnixMilli()-t1)
}

var wg = &sync.WaitGroup{}

type PathDirEntry struct {
	// 文件绝对路径
	Path     string
	DirEntry os.DirEntry
}

func (p *PathDirEntry) Size() (size int64, err error) {
	info, err := p.DirEntry.Info()
	if err != nil {
		return -1, err
	}
	return info.Size(), err
}

func ToPathDirEntries(basePath string, ds []os.DirEntry) (res []PathDirEntry) {
	for _, d := range ds {
		res = append(res, PathDirEntry{
			Path:     basePath + "/" + d.Name(),
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

	pds := ToPathDirEntries(curPath, ds)
	for i := 0; i < len(pds); i++ {
		pd := pds[i]
		dstPath := getDstPath(templatePath, pd.Path)

		// 处理文件夹
		if pd.DirEntry.IsDir() {
			// 创建目录
			err = os.MkdirAll(dstPath, os.ModePerm)
			if err != nil {
				logs.Fatal(err)
			}

			// 大文件夹则开协程处理
			size, err := pd.Size()
			if err != nil {
				logs.Fatal(err)
			}
			if size > 10240 {
				wg.Add(1)
				go func() {
					copyTemplate(templatePath, pd.Path)
					wg.Done()
				}()
			} else {
				copyTemplate(templatePath, pd.Path)
			}
			continue
		}

		// 复制文件
		copyFile(pd.Path, dstPath)
	}
}

func copyFile(srcPath, dstPath string) {
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
	return conf.ProjectName + "/" + srcPath[len(templatePath)+1:]
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
