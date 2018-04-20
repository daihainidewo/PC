// file create by daihao, time is 2018/4/12 16:14
package main

import (
	"os"
	"fmt"
	"os/signal"
	"syscall"
	"golang/utils"
	"golang/service"
	"golang/dao"
	"time"
)

func main() {
	// 读取配置文件
	var confPath string
	if len(os.Args) == 1 {
		fmt.Println("请输入配置文件地址，否则程序无法进行")
		return
	}
	confPath = os.Args[1]
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

	defer func() {
		dao.RedisCacheDao.Close()
		dao.MysqlProjDao.Close()
		dao.MysqlWWWDao.Close()
	}()

	// 启动cookie
	utils.Htmlcookie = utils.NewHtmlCookie()

	// 初始化变量
	utils.SUBSCRIBENUM = conf.SubscribeNum
	utils.PROJECTNUM = conf.ProjectNum
	utils.PACOUNT = conf.PaCount
	utils.NONEDATASLEEPTIME = time.Duration(conf.NoneDataSleepTime) * time.Millisecond
	utils.COOKIEEXPIRE = time.Duration(conf.CookieExpire) * time.Second

	// 启动爬虫服务
	service.ProjService.CtrlPC()

	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, os.Kill, syscall.SIGTERM)
	go StartRouter(conf.StartPort)
	fmt.Println("server start, port:", conf.StartPort)

	<-signCh
	fmt.Println("server end")
}
