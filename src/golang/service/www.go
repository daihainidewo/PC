// file create by daihao, time is 2018/4/16 16:11
package service

import (
	"golang/dao"
	"encoding/json"
	"golang/entity"
)

type WWWServiceImp struct{}

func NewWWWServiceImp() *WWWServiceImp {
	return new(WWWServiceImp)
}

// 设置用户订阅的相关信息
func (this *WWWServiceImp) SetUserSubMsg(userid, suburl, keyword, site, token string, titlekeyword []string) error {
	// 从mysql中查看是否有已订阅的值
	val, err := dao.MysqlWWWDao.GetUserSubMsg(userid)
	if err != nil {
		return err
	}
	temp := entity.User2SubStruct{
		URL:          suburl,
		Keyword:      keyword,
		Token:        token,
		TitleKeyWord: titlekeyword,
		Site:         site,
	}
	submsg := new([]entity.User2SubStruct)
	if val != "" {
		err := json.Unmarshal([]byte(val), submsg)
		if err != nil {
			return err
		}
	}
	*submsg = append(*submsg, temp)
	ret, err := json.Marshal(submsg)
	if err != nil {
		return err
	}
	_, err = dao.MysqlWWWDao.SetUserSubMsg(userid, string(ret))
	return err
}

// 获取用户订阅的相关信息
func (this *WWWServiceImp) GetUserSubMsg(userid string) (*[]entity.User2SubStruct, error) {
	// 从mysql中查看是否有已订阅的值
	val, err := dao.MysqlWWWDao.GetUserSubMsg(userid)
	if err != nil {
		return nil, err
	}
	submsg := new([]entity.User2SubStruct)
	err = json.Unmarshal([]byte(val), submsg)
	return submsg, err
}

// 通过订阅信息查询订阅用户

// 通过订阅信息设置订阅用户

func (this *WWWServiceImp) GetPCBody(userid, timesp string)  {
	
}