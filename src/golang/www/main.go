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
	"golang/proj"
	"golang/entity"
	"container/list"
	"sync"
	"encoding/json"
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
	// 打印配置文件信息
	confstring, err := json.Marshal(conf)
	if err != nil {
		fmt.Println("配置文件错误，请检查错误，error：", err)
		return
	}
	fmt.Println(string(confstring))
	// 启动服务
	service.WWWService = service.NewWWWService()
	service.ProjService = service.NewProjService()

	proj.PCService = proj.NewPCInit()

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
	utils.PATIME = conf.PaTime

	utils.SubUserMap = make(map[entity.UserSubStruct][]string)
	utils.UserSubMap = make(map[string][]entity.UserSubStruct)
	utils.PageTitleMap = make(map[string]string)
	utils.PageTitleList = list.New()

	utils.PageSM = new(sync.Mutex)

	// 启动爬虫服务
	go service.ProjService.CtrlPC()

	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, os.Kill, syscall.SIGTERM)
	go StartRouter(conf.StartPort)
	fmt.Println("server start, port:", conf.StartPort)

	<-signCh
	fmt.Println("server end")
}
