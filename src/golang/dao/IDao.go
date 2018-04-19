// file create by daihao, time is 2018/4/16 11:47
package dao

import "golang/entity"

type IRedisCacheDao interface {
	GetPCBodyMsg(key string) (*entity.PCQueueStruct, error)
	SetPCBodyMsg(key string, val *entity.PCQueueStruct) error
}

var RedisCacheDao IRedisCacheDao

type IMysqlWWWDao interface {
	SetUserSubMsg(idkey, value string) (int64, error)
	GetUserSubMsg(userid string) (string, error)
	SetSubUserMsg(submsg, userids string) (int64, error)
	GetSubUserMsg(submsg string) (string, error)
	Close()
}

var MysqlWWWDao IMysqlWWWDao

type IMysqlProjDao interface {
	SetPCBody(userid_timest string, value *entity.PCBreakStruct) (int64, error)
	GetPCBody(userid_timest string) (*entity.PCBreakStruct, error)
}

var MysqlProjDao IMysqlProjDao
