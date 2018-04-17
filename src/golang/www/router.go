// file create by daihao, time is 2018/4/12 16:16
package main

import (
	"net/http"
	"fmt"
)

func StartRouter(port int) error {
	http.HandleFunc("/index", index)
	http.HandleFunc("/user/zhuce", userZhuce)
	http.HandleFunc("/user/login", userLogin)
	http.HandleFunc("/user/sub", userSub)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
