// file create by daihao, time is 2018/4/16 11:47
package dao

type IRedisCacheDao interface {
	GetPCBody(key string) string
	SetPCBody(key, val string) error
}

var RedisCacheDao IRedisCacheDao

type IMysqlWWWDao interface {
	SetPCBody(userid, value string) (int64, error)
	GetPCBody() ([][]string, error)
	SetUserSubMsg(userid, value string) (int64, error)
	GetUserSubMsg(userid string) (string, error)
	SetSubUserMsg(submsg, userids string) (int64, error)
	GetSubUserMsg(submsg string) (string, error)
	Close()
}

var MysqlWWWDao IMysqlWWWDao

type IMysqlProjDao interface {
}

var MysqlProjDao IMysqlProjDao
