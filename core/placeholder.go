// author gmfan
// date 2023/6/8

package core

import "github.com/tkgfan/init/conf"

// HandlePlaceholderStr 处理可能包含占位符的字符串
func HandlePlaceholderStr(str string) string {
	rs := []rune(str)
	rs = HandlePlaceholderRunes(rs)
	return string(rs)
}

func HandlePlaceholderBytes(bs []byte) []byte {
	res := HandlePlaceholderStr(string(bs))
	return []byte(res)
}

func HandlePlaceholderRunes(rs []rune) []rune {
	for i := 0; i < len(rs)-1; i++ {
		// 匹配前缀
		if rs[i] != '{' || rs[i+1] != '{' {
			continue
		}

		// 匹配后缀
		flag := false
		j := i + 2
		for ; j < len(rs)-1; j++ {
			if rs[j] != '}' || rs[j+1] != '}' {
				continue
			}
			flag = true
			break
		}

		if flag {
			key := string(rs[i+2 : j])
			// 替换
			if v, ok := conf.PlaceholderMap[key]; ok {
				pre, suf := rs[:i], rs[j+2:]
				vs := []rune(v)
				rs = append(pre, append(vs, suf...)...)
				i = len(pre) + len(vs)
			}
		}
	}
	return rs
}
