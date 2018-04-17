// file create by daihao, time is 2018/4/16 16:11
package service

import (
	"golang/dao"
	"golang/utils"
	"fmt"
)

type WWWServiceImp struct{}

func NewWWWServiceImp() *WWWServiceImp {
	return new(WWWServiceImp)
}

func (this *WWWServiceImp) SetUserSubMsg(userid, suburl, keyword, token string, titlekeyword []string) error {
	// 将信息存放至redis中进行排队
	key := utils.GetWaitPCQueueKey()
	err := dao.RedisCacheDao.SetPCBody(key, userid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 将其余相关信息存放至mysql

	return nil
}
