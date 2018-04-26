// file create by daihao, time is 2018/4/16 12:52
package service

import "golang/entity"

type IProjService interface {
	SetPCBody(userid string, value *entity.PCBreakStruct) error
	StartNextPC() (string, *entity.PCBreakStruct, error)
	SetUserSubMsgNoRead(userid string, val []entity.PageTitleStruct) (int64, error)
	SetUserSubMsgReaded(userid string, val []entity.PageTitleStruct) (int64, error)
	Close()
}

var ProjService IProjService

type IWWWService interface {
	SetUserSubMsg(userid, suburl, keyword, site, token string, titlekeyword []string) error
	GetUserSubMsg(userid string) (*[]entity.User2SubStruct, error)
	SetPCBody(userid, suburl, keyword, site, token string, titlekeyword []string) error
	GetUserReaded(userid string) (*entity.UserSubMsgStruct, error)
	GetUserNoread(userid string) (*entity.UserSubMsgStruct, error)
	Close()
	CheckUser(username string, passwd string) (string, error)
	SetUserMsg(username, userpasswd string) (string, error)
	ChangeUserMsg(userid, username, userpasswd string) error
	GetUserReadMsg(userid string) (*entity.UserSubMsgStruct, error)
}

var WWWService IWWWService
