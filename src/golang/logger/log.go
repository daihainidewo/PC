// file create by daihao, time is 2018/4/25 18:04
package logger

import (
	"fmt"
	"golang/utils"
	"golang/conf"
	"time"
)

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(utils.LogFile, "[%s]:", conf.Conf.LogLevel)
	fmt.Fprintf(utils.LogFile, format, a...)
}
func Println(a ...interface{}) {
	fmt.Fprintf(utils.LogFile, "[%s]:", conf.Conf.LogLevel)
	fmt.Fprintln(utils.LogFile, a...)
}
func LogPrintln(a ...interface{}) {
	t := time.Now()
	fmt.Fprintf(utils.LogFile, "[%s]:[log]:[%s]", conf.Conf.LogLevel, t.Format("15:04:05"))
	fmt.Fprintln(utils.LogFile, a...)
}
func ErrPrintln(a ...interface{}) {
	t := time.Now()
	fmt.Fprintf(utils.LogFile, "[%s]:[error]:[%s]", conf.Conf.LogLevel, t.Format("15:04:05"))
	fmt.Fprintln(utils.LogFile, a...)
}
func DebugPrintln(a ...interface{}) {
	if conf.Conf.LogLevel != "debug" {
		return
	}
	t := time.Now()
	fmt.Fprintf(utils.LogFile, "[%s]:[%s]", "debug", t.Format("15:04:05"))
	fmt.Fprintln(utils.LogFile, a...)
}
func InfoPrintln(a ...interface{}) {
	if conf.Conf.LogLevel != "info" {
		return
	}
	t := time.Now()
	fmt.Fprintf(utils.LogFile, "[%s]:[%s]", "info", t.Format("15:04:05"))
	fmt.Fprintln(utils.LogFile, a...)
}
