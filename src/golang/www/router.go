// file create by daihao, time is 2018/4/12 16:16
package main

import (
	"fmt"
	"golang/conf"
	"golang/utils"
	"net/http"
	"runtime/debug"
	"golang/logger"
)

func StartRouter(port int) error {
	http.HandleFunc("/", index)
	http.HandleFunc("/user/zhuce", userZhuce)
	http.HandleFunc("/user/login", userLogin)
	http.HandleFunc("/user/checkname", userCheckName)
	http.HandleFunc("/user/sub", userSub)

	http.HandleFunc("/user/delsub", userDelSub)
	http.HandleFunc("/user/getsub", userGetSub)
	http.HandleFunc("/user/readed", userReaded)
	http.HandleFunc("/user/noread", userNoread)
	http.HandleFunc("/user/readmsg", userReadMsg)
	http.HandleFunc("/user/test", userTest)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}

// 防止程序挂掉
func errorReport(action string, w http.ResponseWriter) {
	if v := recover(); v != nil {
		if conf.Conf.LogLevel == "debug" {
			return
		}
		debug.PrintStack()
		logger.Println("发生意外错误")
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	}
}
