// author lby
// date 2023/6/8

package core

import (
	"github.com/tkgfan/init/conf"
	"testing"
)

func TestHandlePlaceholderStr(t *testing.T) {
	conf.PlaceholderMap["pl1"] = "占位内容"
	conf.PlaceholderMap["pl2"] = "占位 符"
	tests := []struct {
		arg  string
		want string
	}{
		{
			arg:  "hello world{{",
			want: "hello world{{",
		},
		{
			arg:  "苹果 香蕉 {{pl1}}",
			want: "苹果 香蕉 占位内容",
		},
		{
			arg:  "苹果 香蕉 {{{{pl1}}}",
			want: "苹果 香蕉 {{占位内容}",
		},
		{
			arg:  "苹果 {{pl2}}香蕉 {{{{pl}}1}}}",
			want: "苹果 占位 符香蕉 {{{{pl}}1}}}",
		},
		{
			arg:  "{{{pl2}}苹果 香蕉 {{{{pl1}}}",
			want: "{占位 符苹果 香蕉 {{占位内容}",
		},
	}

	for _, tt := range tests {
		got := HandlePlaceholderStr(tt.arg)
		if got != tt.want {
			t.Errorf("want=%s,got=%s", tt.want, got)
		}
	}
}
