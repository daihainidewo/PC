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
func (this *ProjServiceImp) startNextPC() (string, error) {
	// redis中获取排队数据
	redikey := utils.GetWaitPCQueueKey()
	pcbody, err := dao.RedisCacheDao.GetPCBodyMsg(redikey)
	if err != nil {
		return "", err
	}
	if pcbody == nil {
		return "", nil
	}
	// 从mysql获取上次中断信息
	pcbs, err := dao.MysqlProjDao.GetPCBody(utils.GetUserTimeMysqlKey(pcbody))
	if err != nil {
		return "", err
	}
	if len(pcbs.PageTitleList2Slice) == 0 {
		utils.PageTitleList = list.New()
		proj.PCService.StartPC(pcbs.URL, pcbs.Keyword, pcbs.Site, pcbs.Token, pcbody.Userid, pcbs.TitleKeyWord)
		return "", nil
	}
	// 载入内存
	for _, l := range pcbs.PageTitleList2Slice {
		utils.PageTitleList.PushBack(l)
	}
	utils.PageTitleMap = pcbs.PageTitleMap
	// 启动爬虫
	pcbs.URL = pcbs.PageTitleList2Slice[0]
	utils.PageTitleList.Remove(utils.PageTitleList.Front())
	proj.PCService.StartPC(pcbs.URL, pcbs.Keyword, pcbs.Site, pcbs.Token, pcbody.Userid, pcbs.TitleKeyWord)
	return pcbody.Userid, nil
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

func (this *ProjServiceImp) CtrlPC() {
	go func() {
		for {
			utils.UserSubUrl = make([]entity.PageTitleStruct, 0)
			// 准备下一个爬虫
			userid, err := this.startNextPC()
			if err != nil {
				if fmt.Sprintf("%s", err) != "redis: nil" {
					fmt.Println(err)
				}
				continue
			}
			this.SetUserSubMsgNoRead(userid, utils.UserSubUrl)
			time.Sleep(1 * time.Minute)
		}
	}()
}

// 获取未读消息
func (this *ProjServiceImp) GetUserSubMsgNoRead(userid string) (*entity.UserSubMsgStruct, error) {
	return dao.MysqlProjDao.GetUserSubMsgNoRead(userid)
}

// 获取已读消息
func (this *ProjServiceImp) GetUserSubMsgReaded(userid string) (*entity.UserSubMsgStruct, error) {
	return dao.MysqlProjDao.GetUserSubMsgReaded(userid)
}

// 设置未读消息
func (this *ProjServiceImp) SetUserSubMsgNoRead(userid string, val []entity.PageTitleStruct) (int64, error) {
	oldval, err := dao.MysqlProjDao.GetUserSubMsgNoRead(userid)
	if err != nil {
		return -1, err
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
	oldval, err := dao.MysqlProjDao.GetUserSubMsgReaded(userid)
	if err != nil {
		return -1, err
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
