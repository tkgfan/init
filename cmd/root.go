// author gmfan
// date 2023/5/9

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tkgfan/init/conf"
	"os"
	"runtime"
)

var rootCmd = &cobra.Command{
	Use:   "init",
	Short: "模版初始化工具",
	Long:  `模版初始化工具，提供从 Git 上拉取模版来初始化项目功能`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetArgs([]string{"-v"})
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", conf.Version(), runtime.GOOS, runtime.GOARCH)
}
