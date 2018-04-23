// file create by daihao, time is 2018/4/16 11:47
package dao

import "golang/entity"

type IRedisCacheDao interface {
	GetPCBodyMsg(key string) (*entity.PCQueueStruct, error)
	SetPCBodyMsg(key string, val *entity.PCQueueStruct) error
	Close()
}

var RedisCacheDao IRedisCacheDao

type IMysqlWWWDao interface {
	InsertUserSubMsg(idkey string, value *[]entity.User2SubStruct) (int64, error)
	UpdateUserSubMsg(idkey string, value *[]entity.User2SubStruct) (int64, error)
	SelectUserSubMsg(userid string) (*[]entity.User2SubStruct, error)
	InsertSubUserMsg(submsg, userids string) (int64, error)
	SelectSubUserMsg(submsg string) (string, error)
	SelectUserSubMsgNoRead(userid string) (*entity.UserSubMsgStruct, error)
	SelectUserSubMsgReaded(userid string) (*entity.UserSubMsgStruct, error)
	Close()
}

var MysqlWWWDao IMysqlWWWDao

type IMysqlProjDao interface {
	InsertPCBody(userid_timest string, value *entity.PCBreakStruct) (int64, error)
	SelectPCBody(userid_timest string) (*entity.PCBreakStruct, error)
	InsertUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error)
	UpdateUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error)
	InsertUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error)
	UpdateUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error)
	Close()
}

var MysqlProjDao IMysqlProjDao
