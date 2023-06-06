// author gmfan
// date 2023/5/9

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tkgfan/init/conf"
	"github.com/tkgfan/init/core"
	"os"
	"runtime"
)

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", conf.Version(), runtime.GOOS, runtime.GOARCH)

	rootCmd.Flags().BoolVarP(&conf.OnlyUseLocalTemplate, "onlyUseLocalTemplate", "o", false, "只使用本地模版")
	rootCmd.Flags().StringVarP(&conf.ProjectName, "name", "n", ".", "项目名称，默认为当前目录")
	rootCmd.Flags().StringVarP(&conf.TemplateName, "template", "t", "default", "指定模版")
}

var rootCmd = &cobra.Command{
	Use:     "init",
	Short:   "模版初始化工具",
	Long:    `模版初始化工具，提供使用 Git 来拉取远程模版并使用该模版的功能`,
	Example: "init .",
	Run: func(cmd *cobra.Command, args []string) {
		core.ProjectInit()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
