// file create by daihao, time is 2018/4/25 18:04
package logger

import (
	"fmt"
	"golang/utils"
	"golang/conf"
	"time"
	"os"
)

func StartLog(path string){
	if path != "" {
		tick := time.NewTicker(1 * time.Hour)
		go func() {
			for range tick.C {
				logpath := utils.GetCurLogPath(path)
				logf, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
				if err != nil {
					fmt.Println("读取日志文件错误，error：", err)
				}
				utils.LogFile = logf
			}
		}()
	} else {
		utils.LogFile = os.Stdout
	}
}

// 打印普通消息函数
func Printf(format string, a ...interface{}) {
	fmt.Fprintf(utils.LogFile, "[%s]:", conf.Conf.LogLevel)
	fmt.Fprintf(utils.LogFile, format, a...)
}
func Println(a ...interface{}) {
	fmt.Fprintf(utils.LogFile, "[%s]:", conf.Conf.LogLevel)
	fmt.Fprintln(utils.LogFile, a...)
}

// 打印日志消息
func LogPrintln(a ...interface{}) {
	t := time.Now()
	fmt.Fprintf(utils.LogFile, "[%s]:[log]:[%s]", conf.Conf.LogLevel, t.Format("15:04:05"))
	fmt.Fprintln(utils.LogFile, a...)
}
// 打印错误消息
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
