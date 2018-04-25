// file create by daihao, time is 2018/4/16 12:51
package service

import (
	"fmt"
	"golang/dao"
	"golang/entity"
	"golang/proj"
	"golang/utils"
	"strings"
	"time"
	"golang/logger"
)

type ProjServiceImp struct {
}

func NewProjService() *ProjServiceImp {
	return new(ProjServiceImp)
}

// 通过获取redis和mysql的数据启动爬虫程序
func (this *ProjServiceImp) startNextPC() (string, *entity.PCBreakStruct, error) {
	// redis中获取排队数据
	redikey := utils.GetWaitPCQueueKey()
	pcbody, err := dao.RedisCacheDao.GetPCBodyMsg(redikey)
	if err != nil {
		return "", nil, fmt.Errorf("startNextPC:%s", err)
	}
	if pcbody == nil {
		return "", nil, nil
	}
	logger.Println("get new")
	logger.Println(pcbody)
	// 从mysql获取上次中断信息
	pcbs, err := dao.MysqlProjDao.SelectPCBody(utils.GetUserTimeMysqlKey(pcbody))
	if err != nil {
		return "", nil, fmt.Errorf("startNextPC:%s", err)
	}
	if pcbs == nil {
		return "", nil, nil
	}
	// 清除本次信息
	_, err = dao.MysqlProjDao.DeletePCBody(utils.GetUserTimeMysqlKey(pcbody))
	if err != nil {
		return "", nil, fmt.Errorf("startNextPC:%s", err)
	}
	// 载入内存
	for _, l := range pcbs.PageTitleList2Slice {
		utils.PageTitleList.PushBack(l)
	}
	utils.PageTitleMap = pcbs.PageTitleMap

	// 设置本次开始的url
	if len(pcbs.PageTitleList2Slice) != 0 {
		pcbs.URL = pcbs.PageTitleList2Slice[0]
		utils.PageTitleList.Remove(utils.PageTitleList.Front())
	}
	return pcbody.Userid, pcbs, nil
}

// 将爬虫中断程序存放至redis和mysql中进行排队
func (this *ProjServiceImp) SetPCBody(userid string, value *entity.PCBreakStruct) error {
	timest := time.Now().UnixNano()
	pcbody := new(entity.PCQueueStruct)
	pcbody.Userid = userid
	pcbody.Timest = fmt.Sprintf("%d", timest)
	// 存放入mysql
	_, err := dao.MysqlProjDao.InsertPCBody(utils.GetUserTimeMysqlKey(pcbody), value)
	if err != nil {
		return fmt.Errorf("[Service]ProjServiceImp:SetPCBody:%s", err)
	}
	// 存放入redis进行排队
	redikey := utils.GetWaitPCQueueKey()
	err = dao.RedisCacheDao.SetPCBodyMsg(redikey, pcbody)
	if err != nil {
		return fmt.Errorf("[Service]ProjServiceImp:SetPCBody:%s", err)
	}
	return nil
}

func (this *ProjServiceImp) CtrlPC() {
	ch := make(chan int, 1)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			logger.Println("the next pc")
			utils.UserSubUrl = make([]entity.PageTitleStruct, 0)
			// 准备下一个爬虫
			userid, pcbs, err := this.startNextPC()
			if err != nil {
				if !strings.Contains(err.Error(), "redis: nil") {
					logger.Println(err)
				}
				continue
			}
			logger.Println("PC ing ...")
			proj.PCService.StartPC(pcbs.URL, pcbs.Keyword, pcbs.Site, pcbs.Token, userid, pcbs.TitleKeyWord)
			if userid == "" {
				logger.Println("userid is nil")
				continue
			}
			_, err = this.SetUserSubMsgNoRead(userid, utils.UserSubUrl)
			if err != nil {
				logger.Println(err)
				continue
			}
			logger.Println("set one")
			time.Sleep(5 * time.Second)
			// 将爬虫存放进爬取队列
			err = this.SetPCBody(userid, pcbs)
			if err != nil {
				logger.Println(err)
				continue
			}
		}
	}()
	<-ch
}

// 设置未读消息
func (this *ProjServiceImp) SetUserSubMsgNoRead(userid string, val []entity.PageTitleStruct) (int64, error) {
	oldval, err := dao.MysqlWWWDao.SelectUserSubMsgNoRead(userid)
	if err != nil {
		return -1, fmt.Errorf("[Service]ProjServiceImp:SetUserSubMsgNoRead:%s", err)
	}
	if oldval == nil {
		res := new(entity.UserSubMsgStruct)
		res.SubMsg = val
		res.Userid = userid
		return dao.MysqlProjDao.InsertUserSubMsgNoRead(userid, res)
	}
	for _, v := range val {
		oldval.SubMsg = append(oldval.SubMsg, v)
	}
	return dao.MysqlProjDao.UpdateUserSubMsgNoRead(userid, oldval)
}

// 设置已读消息
func (this *ProjServiceImp) SetUserSubMsgReaded(userid string, val []entity.PageTitleStruct) (int64, error) {
	oldval, err := dao.MysqlWWWDao.SelectUserSubMsgReaded(userid)
	if err != nil {
		return -1, fmt.Errorf("[Service]ProjServiceImp:SetUserSubMsgReaded:%s", err)
	}
	if oldval == nil {
		res := new(entity.UserSubMsgStruct)
		res.SubMsg = val
		res.Userid = userid
		return dao.MysqlProjDao.InsertUserSubMsgReaded(userid, res)
	}
	for _, v := range val {
		oldval.SubMsg = append(oldval.SubMsg, v)
	}
	return dao.MysqlProjDao.UpdateUserSubMsgReaded(userid, oldval)
}

func (this *ProjServiceImp) Close() {

}
