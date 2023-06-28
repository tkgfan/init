// author gmfan
// date 2023/6/6

package conf

import (
	"github.com/tkgfan/init/common/logs"
	"strings"
)

// root flag

// OnlyUseLocalTemplate 只使用本地模版
var OnlyUseLocalTemplate bool

// ProjectName 项目，使用 . 代表但前目录
var ProjectName string

// TemplateName 使用的模版名称
var TemplateName string

// Placeholders 占位符参数
var Placeholders *[]string

// PlaceholderMap 解析占位符后的参数
var PlaceholderMap = make(map[string]string)

// ParsePlaceholders 解析占位符参数
func ParsePlaceholders() {
	// 处理占位符参数
	for _, p := range *Placeholders {
		p = strings.Trim(p, " ")
		idx := strings.Index(p, "=")
		if idx > 0 && idx < len(p)-1 {
			PlaceholderMap[p[0:idx]] = p[idx+1:]
		} else {
			logs.Info("占位符参数:", p, "设置不正确")
		}
	}
}

// config flag

// Repository 远程仓库地址
var Repository string
