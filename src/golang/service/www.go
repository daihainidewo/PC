// file create by daihao, time is 2018/4/16 16:11
package service

import (
	"fmt"
	"golang/dao"
	"golang/entity"
	"golang/utils"
	"time"
	"golang/logger"
	"encoding/json"
)

type WWWServiceImp struct{}

func NewWWWService() *WWWServiceImp {
	return new(WWWServiceImp)
}

// 设置用户订阅的相关信息
func (this *WWWServiceImp) SetUserSubMsg(userid, suburl, keyword, site, token string, titlekeyword []string) error {
	// 从mysql中查看是否有已订阅的值
	val, err := dao.MysqlWWWDao.SelectUserSubMsg(userid)
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

	// 添加订阅信息至mysql
	if val == nil {
		val = new([]entity.User2SubStruct)
		*val = append(*val, temp)
		_, err = dao.MysqlWWWDao.InsertUserSubMsg(userid, val)
		if err != nil {
			return fmt.Errorf("[Service]WWWServiceImp:SetUserSubMsg:%s", err)
		}
	} else {
		// 查重
		for _, v := range *val {
			if utils.User2SubStructIsEqual(v, temp) {
				return nil
			}
		}
		*val = append(*val, temp)
		_, err = dao.MysqlWWWDao.UpdateUserSubMsg(userid, val)
		if err != nil {
			return fmt.Errorf("[Service]WWWServiceImp:SetUserSubMsg:%s", err)
		}
	}

	// 将信息存放至mysql中
	tempsub := new(entity.PCQueueStruct)
	tempsub.Userid = userid
	tempsub.Timest = fmt.Sprintf("%d", time.Now().Unix())
	tempstruct := new(entity.PCBreakStruct)
	tempstruct.User2SubStruct = temp

	//utils.PageTitleMap[starttoken.URL] = starttoken.Title
	//utils.PageTitleList.PushBack(starttoken)
	starttoken := entity.PageTitleStruct{"", temp.URL}
	tempstruct.PageTitleList2Slice = make([]string, 1)
	t, _ := json.Marshal(starttoken)
	tempstruct.PageTitleList2Slice[0] = string(t)
	tempstruct.PageTitleMap = make(map[string]string)
	logger.Println(tempstruct)
	// 添加信息至mysql和redis排队
	err = ProjService.SetPCBody(userid, tempstruct)

	return err
}

// 获取用户订阅的相关信息
func (this *WWWServiceImp) GetUserSubMsg(userid string) (*[]entity.User2SubStruct, error) {
	return dao.MysqlWWWDao.SelectUserSubMsg(userid)
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
		return fmt.Errorf("[Service]WWWServiceImp:SetPCBody:%s", err)
	}
	return nil
}

// 查看用户已读消息
func (this *WWWServiceImp) GetUserReaded(userid string) (*entity.UserSubMsgStruct, error) {
	ret, err := dao.MysqlWWWDao.SelectUserSubMsgReaded(userid)
	if err != nil {
		return nil, fmt.Errorf("[Service]WWWServiceImp:GetUserReaded:%s", err)
	}
	if ret != nil && ret.Userid == "" {
		ret.Userid = userid
	}
	return ret, nil
}

// 查看用户未读消息
func (this *WWWServiceImp) GetUserNoread(userid string) (*entity.UserSubMsgStruct, error) {
	ret, err := dao.MysqlWWWDao.SelectUserSubMsgNoRead(userid)
	if err != nil {
		return nil, fmt.Errorf("[Service]WWWServiceImp:GetUserNoread:%s", err)
	}
	if ret != nil && ret.Userid == "" {
		ret.Userid = userid
	}
	return ret, nil
}

// 设置用户信息
func (this *WWWServiceImp) SetUserMsg(username, userpasswd string) (string, error) {
	ok, err := dao.MysqlWWWDao.SelectUserSameName(username)
	if err != nil {
		return "", fmt.Errorf("[Service]WWWServiceImp:SetUserMsg:%s", err)
	}
	// 表示有相同的用户名
	if ok {
		return "", nil
	}
	userid := fmt.Sprintf("%d", time.Now().UnixNano())[7:17]
	//fmt.Println(userid)
	_, err = dao.MysqlWWWDao.InsertUserMsg(userid, username, userpasswd)
	return userid, err
}

// 用户修改相关信息
func (this *WWWServiceImp) ChangeUserMsg(userid, username, userpasswd string) error {
	if username != "" {
		_, err := dao.MysqlWWWDao.UpdateUserMsgUsername(userid, username)
		if err != nil {
			return fmt.Errorf("[Service]WWWServiceImp:ChangeUserMsg:%s", err)
		}
	}
	if userpasswd != "" {
		_, err := dao.MysqlWWWDao.UpdateUserMsgUserpasswd(userid, userpasswd)
		if err != nil {
			return fmt.Errorf("[Service]WWWServiceImp:ChangeUserMsg:%s", err)
		}
	}
	return nil
}

// 获取用户信息
func (this *WWWServiceImp) GetUserMsg(userid string) (string, string, error) {
	return dao.MysqlWWWDao.SelectUserMsg(userid)
}

// 检测用户名与密码是否匹配
func (this *WWWServiceImp) CheckUser(username string, userpasswd string) (string, error) {
	return dao.MysqlWWWDao.CheckUserNamePasswd(username, userpasswd)
}

func (this *WWWServiceImp) GetUserReadMsg(userid string) (*entity.UserSubMsgStruct, error) {
	data, err := dao.MysqlWWWDao.SelectUserSubMsgNoRead(userid)
	if err != nil {
		return nil, fmt.Errorf("[Service]WWWServiceImp:GetUserReadMsg:%s", err)
	}
	readed, err := dao.MysqlWWWDao.SelectUserSubMsgReaded(userid)
	if err != nil {
		return nil, fmt.Errorf("[Service]WWWServiceImp:GetUserReadMsg:%s", err)
	}
	if readed != nil {
		readed = data
	} else {
		for _, d := range data.SubMsg {
			readed.SubMsg = append(readed.SubMsg, d)
		}
	}
	_, err = dao.MysqlWWWDao.UpdateUserSubMsgReaded(readed)
	if err != nil {
		return nil, fmt.Errorf("[Service]WWWServiceImp:GetUserReadMsg:%s", err)
	}
	_, err = dao.MysqlWWWDao.DelectUserSubMsgNoRead(userid)
	if err != nil {
		return nil, fmt.Errorf("[Service]WWWServiceImp:GetUserReadMsg:%s", err)
	}
	return data, nil
}

func (this *WWWServiceImp) Close() {

}
