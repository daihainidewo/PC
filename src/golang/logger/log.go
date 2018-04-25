// file create by daihao, time is 2018/4/25 18:04
package logger

import (
	"fmt"
	"golang/utils"
)

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(utils.LogFile, format, a...)
}
func Println(a ...interface{}) {
	fmt.Fprintln(utils.LogFile, a...)
}
func Errorf(format string, a ...interface{}) {
	fmt.Fprintf(utils.LogFile, format, a...)
}
