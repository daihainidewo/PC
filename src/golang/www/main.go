// file create by daihao, time is 2018/4/12 16:14
package main

import (
	"os"
	"fmt"
	"os/signal"
	"syscall"
	"golang/utils"
	"os/exec"
	"golang/service"
	"golang/dao"
	"flag"
)

func main() {
	confPath := *flag.String("conf", "", "conf file path")

	fmt.Println("reading...", confPath)
	conf, err := utils.ReadConf(confPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 启动服务
	service.WWWService = service.NewWWWService()
	service.ProjService = service.NewProjService()

	dao.RedisCacheDao = dao.NewRedisCache(conf.RedisAddr, conf.RedisPasswd, conf.RedisDB)
	dao.MysqlProjDao = dao.NewProjMysqlClient(conf.MysqlProjDriverName, conf.MysqlProjDataSourceName)
	dao.MysqlWWWDao = dao.NewWWWMysqlClient(conf.MysqlWWWDriverName, conf.MysqlWWWDataSourceName)

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
