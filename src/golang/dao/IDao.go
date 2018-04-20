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
	SetUserSubMsg(idkey string, value *[]entity.User2SubStruct) (int64, error)
	SetUserSubMsgUpdate(idkey string, value *[]entity.User2SubStruct) (int64, error)
	GetUserSubMsg(userid string) (*[]entity.User2SubStruct, error)
	SetSubUserMsg(submsg, userids string) (int64, error)
	GetSubUserMsg(submsg string) (string, error)
	Close()
}

var MysqlWWWDao IMysqlWWWDao

type IMysqlProjDao interface {
	SetPCBody(userid_timest string, value *entity.PCBreakStruct) (int64, error)
	GetPCBody(userid_timest string) (*entity.PCBreakStruct, error)
	GetUserSubMsgNoRead(userid string) (*entity.UserSubMsgStruct, error)
	GetUserSubMsgReaded(userid string) (*entity.UserSubMsgStruct, error)
	InsertUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error)
	UpdateUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error)
	InsertUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error)
	UpdateUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error)
	Close()
}

var MysqlProjDao IMysqlProjDao
