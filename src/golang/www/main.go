// file create by daihao, time is 2018/4/12 16:14
package main

import (
	"os"
	"fmt"
	"os/signal"
	"syscall"
	"golang/utils"
	"os/exec"
)

func main() {
	// 设置静态网页
	htmlPath, _ := exec.Command("sh", "-c", `echo $GOPATH`).Output()
	utils.HtmlPath = string(htmlPath[:len(htmlPath)-1]) + "/html/"

	// 启动cookie
	utils.Htmlcookie = utils.NewHtmlCookie()

	port := 8080
	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, os.Kill, syscall.SIGTERM)
	go StartRouter(port)
	fmt.Println("server start, port:", port)

	<-signCh
	fmt.Println("server end")
}
