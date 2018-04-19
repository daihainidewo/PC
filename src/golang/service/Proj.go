// file create by daihao, time is 2018/4/16 12:51
package service

import (
	"golang/dao"
	"golang/utils"
	"golang/proj"
	"container/list"
	"time"
	"fmt"
	"golang/entity"
)

type ProjServiceImp struct {
}

func NewProjService() *ProjServiceImp {
	return new(ProjServiceImp)
}

// 通过获取redis和mysql的数据启动爬虫程序
func (this *ProjServiceImp) StartNextPC() error {
	// redis中获取排队数据
	redikey := utils.GetWaitPCQueueKey()
	pcbody, err := dao.RedisCacheDao.GetPCBodyMsg(redikey)
	if err != nil {
		return err
	}
	// 从mysql获取上次中断信息
	pcbs, err := dao.MysqlProjDao.GetPCBody(utils.GetUserTimeMysqlKey(pcbody))
	if err != nil {
		return err
	}
	if len(pcbs.PageTitleList2Slice) == 0 {
		utils.PageTitleList = list.New()
		proj.StartPC(pcbs.URL, pcbs.Keyword, pcbs.Site, pcbs.Token, pcbody.Userid, pcbs.TitleKeyWord)
		return err
	}
	// 载入内存
	for _, l := range pcbs.PageTitleList2Slice {
		utils.PageTitleList.PushBack(l)
	}
	utils.PageTitleMap = pcbs.PageTitleMap
	// 启动爬虫
	pcbs.URL = pcbs.PageTitleList2Slice[0]
	utils.PageTitleList.Remove(utils.PageTitleList.Front())
	proj.StartPC(pcbs.URL, pcbs.Keyword, pcbs.Site, pcbs.Token, pcbody.Userid, pcbs.TitleKeyWord)
	return nil
}

// 将爬虫中断程序存放至redis和mysql中进行排队
func (this *ProjServiceImp) SetPCBody(userid string, value *entity.PCBreakStruct) error {
	timest := time.Now().UnixNano()
	pcbody := new(entity.PCQueueStruct)
	pcbody.Userid = userid
	pcbody.Timest = fmt.Sprintf("%d", timest)
	// 存放入mysql
	_, err := dao.MysqlProjDao.SetPCBody(utils.GetUserTimeMysqlKey(pcbody), value)
	if err != nil {
		return err
	}
	// 存放入redis进行排队
	redikey := utils.GetWaitPCQueueKey()
	err = dao.RedisCacheDao.SetPCBodyMsg(redikey, pcbody)
	return err
}
