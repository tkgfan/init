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
	// 注册命令
	registerCmd()

	rootCmd.Version = fmt.Sprintf("%s %s/%s", conf.Version(), runtime.GOOS, runtime.GOARCH)

	rootCmd.Flags().BoolVarP(&conf.OnlyUseLocalTemplate, "onlyUseLocalTemplate", "o", false, "只使用本地模版")
	rootCmd.Flags().StringVarP(&conf.ProjectName, "name", "n", ".", "项目名称，默认为当前目录")
	rootCmd.Flags().StringVarP(&conf.TemplateName, "template", "t", "default", "指定模版")
	conf.Placeholders = rootCmd.Flags().StringSliceP("placeholders", "p", []string{}, "占位符参数使用方法 -p placeholder1=arg1,placeholder2=arg2")
}

func registerCmd() {
	rootCmd.AddCommand(configCmd)
}

var rootCmd = &cobra.Command{
	Use:     "init",
	Short:   "模版初始化工具",
	Long:    `模版初始化工具，提供使用 Git 来拉取远程模版并使用该模版的功能`,
	Example: "init	// 默认初始化当前目录，使用默认 default 模版",
	Run: func(cmd *cobra.Command, args []string) {
		conf.ParsePlaceholders()
		core.ProjectInit()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
