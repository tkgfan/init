// author gmfan
// date 2023/6/5

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var dirCmd = &cobra.Command{
	Use:     "dir",
	Short:   "指定初始化目录，没有指定目录则当前目录作为根目录",
	Long:    "指定初始化目录，没有指定目录则当前目录作为根目录",
	Example: "dir example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
