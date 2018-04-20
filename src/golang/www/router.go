// file create by daihao, time is 2018/4/12 16:16
package main

import (
	"net/http"
	"fmt"
	"runtime/debug"
	"golang/utils"
)

func StartRouter(port int) error {
	http.HandleFunc("/", index)
	http.HandleFunc("/user/zhuce", userZhuce)
	http.HandleFunc("/user/login", userLogin)
	http.HandleFunc("/user/sub", userSub)
	http.HandleFunc("/user/getSub", userGetSub)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}

// 防止程序挂掉
func errorReport(action string, w http.ResponseWriter) {
	if v := recover(); v != nil {
		debug.PrintStack()
		fmt.Println("发生意外错误")
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	}
}
