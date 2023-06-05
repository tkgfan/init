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

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", conf.Version(), runtime.GOOS, runtime.GOARCH)
	rootCmd.AddCommand(dirCmd)
}

var rootCmd = &cobra.Command{
	Use:     "init",
	Short:   "模版初始化工具",
	Long:    `模版初始化工具，提供使用 Git 来拉取远程模版并使用该模版的功能`,
	Example: "init .",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
