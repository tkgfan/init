// author lby
// date 2023/6/6

package logs

import (
	"fmt"
	"log"
)

// Fatal 致命错误
func Fatal(args ...any) {
	log.SetFlags(0)
	log.Fatal(args)
}

func Info(args ...any) {
	fmt.Println(args)
}
