// file create by daihao, time is 2018/4/12 16:14
package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"golang/conf"
	"golang/dao"
	"golang/entity"
	"golang/proj"
	"golang/service"
	"golang/utils"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"golang/logger"
)

func main() {
	// 读取配置文件
	var confPath string
	if len(os.Args) == 1 {
		fmt.Println("请输入配置文件地址，否则程序无法进行")
		return
	}
	confPath = os.Args[1]
	var err error
	conf.Conf, err = utils.ReadConf(confPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印配置文件信息
	confstring, err := json.Marshal(conf.Conf)
	if err != nil {
		fmt.Println("配置文件错误，请检查错误，error：", err)
		return
	}
	// 设置进程日志
	if conf.Conf.LogPath != "" {
		tick := time.NewTicker(1 * time.Hour)
		go func() {
			for range tick.C {
				logpath := utils.GetCurLogPath(conf.Conf.LogPath)
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
	// 启动服务
	service.WWWService = service.NewWWWService()
	service.ProjService = service.NewProjService()

	proj.PCService = proj.NewPCInit()

	dao.RedisCacheDao = dao.NewRedisCache(conf.Conf.RedisAddr, conf.Conf.RedisPasswd, conf.Conf.RedisDB)
	dao.MysqlProjDao = dao.NewProjMysqlClient(conf.Conf.MysqlProjDriverName, conf.Conf.MysqlProjDataSourceName)
	dao.MysqlWWWDao = dao.NewWWWMysqlClient(conf.Conf.MysqlWWWDriverName, conf.Conf.MysqlWWWDataSourceName)

	defer func() {
		dao.RedisCacheDao.Close()
		dao.MysqlProjDao.Close()
		dao.MysqlWWWDao.Close()
	}()

	// 启动cookie
	utils.Htmlcookie = utils.NewHtmlCookie()

	// 初始化变量
	utils.SUBSCRIBENUM = conf.Conf.SubscribeNum
	utils.PROJECTNUM = conf.Conf.ProjectNum
	utils.PACOUNT = conf.Conf.PaCount
	utils.NONEDATASLEEPTIME = time.Duration(conf.Conf.NoneDataSleepTime) * time.Millisecond
	utils.COOKIEEXPIRE = time.Duration(conf.Conf.CookieExpire) * time.Second
	utils.PATIME = conf.Conf.PaTime

	utils.SubUserMap = make(map[entity.UserSubStruct][]string)
	utils.UserSubMap = make(map[string][]entity.UserSubStruct)
	utils.PageTitleMap = make(map[string]string)
	utils.PageTitleList = list.New()

	utils.PageSM = new(sync.Mutex)

	// 启动爬虫服务
	go proj.PCService.CtrlPC()

	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, os.Kill, syscall.SIGTERM)
	go StartRouter(conf.Conf.StartPort)
	logger.Println("server start, port:", conf.Conf.StartPort)
	logger.Println(string(confstring))

	<-signCh
	logger.Println("server end")
}
