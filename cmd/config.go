// Package cmd
// author gmfan
// date 2023/6/10
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tkgfan/init/conf"
)

func init() {
	configCmd.Flags().StringVarP(&conf.Repository, "repository", "r", "", "设置自定义仓库地址")
}

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "配置",
	Long:    "自定义配置",
	Example: "init config",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
