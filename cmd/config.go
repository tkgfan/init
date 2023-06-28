// Package cmd
// author gmfan
// date 2023/6/10
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tkgfan/init/common/logs"
	"github.com/tkgfan/init/conf"
)

func init() {
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configSetCmd)

	configSetCmd.Flags().StringVarP(&conf.Repository, "repository", "r", "", "设置自定义仓库地址")
}

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "配置",
	Long:    "自定义配置",
	Example: "init config",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var configListCmd = &cobra.Command{
	Use:     "list",
	Short:   "查看配置列表",
	Long:    "查看配置列表",
	Example: "init config list",
	Run: func(cmd *cobra.Command, args []string) {
		logs.Info("远程仓库地址:", conf.Config.Repository)
	},
}

var configSetCmd = &cobra.Command{
	Use:     "set",
	Short:   "设置配置",
	Long:    "设置配置",
	Example: "init config set -r git@github.com:tkgfan/template.git",
	Run: func(cmd *cobra.Command, args []string) {
		if conf.Repository != "" {
			conf.Config.Repository = conf.Repository
		}

		err := conf.FlushConfig()
		if err != nil {
			logs.Fatal("保存配置失败", err)
		}
	},
}
