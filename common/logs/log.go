// author lby
// date 2023/6/6

package logs

import "fmt"

// Fatal 致命错误
func Fatal(msg any) {
	panic(msg)
}

func Info(args ...any) {
	fmt.Println(args)
}
