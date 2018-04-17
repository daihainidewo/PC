// file create by daihao, time is 2018/4/16 12:52
package service

type IProjService interface {
}

var ProjService IProjService

type IWWWService interface {
	SetUserSubMsg(userid, suburl, keyword, token string, titlekeyword []string) error
}

var WWWService IWWWService
