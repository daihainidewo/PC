// file create by daihao, time is 2018/4/16 16:11
package service

import (
	"golang/dao"
	"golang/entity"
	"golang/utils"
)

type WWWServiceImp struct{}

func NewWWWService() *WWWServiceImp {
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

	if val == nil {
		val = new([]entity.User2SubStruct)
		*val = append(*val, temp)
		_, err = dao.MysqlWWWDao.SetUserSubMsg(userid, val)
	} else {
		// 查重
		for _, v := range *val {
			if utils.User2SubStructIsEqual(v, temp) {
				return nil
			}
		}
		*val = append(*val, temp)
		_, err = dao.MysqlWWWDao.SetUserSubMsgUpdate(userid, val)
	}
	return err
}

// 获取用户订阅的相关信息
func (this *WWWServiceImp) GetUserSubMsg(userid string) (*[]entity.User2SubStruct, error) {
	// 从mysql中查看是否有已订阅的值
	//val, err := dao.MysqlWWWDao.GetUserSubMsg(userid)
	//if err != nil {
	//	return nil, err
	//}
	return dao.MysqlWWWDao.GetUserSubMsg(userid)
}

// 通过订阅信息查询订阅用户

// 通过订阅信息设置订阅用户

// 设置爬虫信息
func (this *WWWServiceImp) SetPCBody(userid, suburl, keyword, site, token string, titlekeyword []string) error {
	val := new(entity.PCBreakStruct)
	val.TitleKeyWord = titlekeyword
	val.Token = token
	val.Site = site
	val.URL = suburl
	val.Keyword = keyword
	val.PageTitleMap = make(map[string]string)
	val.PageTitleList2Slice = make([]string, 0)
	err := ProjService.SetPCBody(userid, val)
	if err != nil {
		return err
	}
	return nil
}

func (this *WWWServiceImp) Close() {

}