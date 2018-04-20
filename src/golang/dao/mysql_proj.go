// file create by daihao, time is 2018/4/16 11:02
package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"golang/entity"
	"encoding/json"
)

type MysqlProjClientImp struct {
	client *sql.DB
}

func NewProjMysqlClient(driverName, dataSourceName string) *MysqlProjClientImp {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	return &MysqlProjClientImp{client: db}
}

func (this *MysqlProjClientImp) doSQL(sql string, args ...interface{}) (sql.Result, error) {
	stmt, e := this.client.Prepare(sql)
	if e != nil {
		return nil, e
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (this *MysqlProjClientImp) doQuery(sql string, args ...interface{}) ([][]interface{}, error) {
	rows, err := this.client.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, nil
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	length := len(columns)
	ret := make([][]interface{}, 0)
	scanArgs := make([]interface{}, length)
	for i := range scanArgs {
		scanArgs[i] = new(interface{})
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		temp := make([]interface{}, length)

		for idx, col := range scanArgs {
			temp[idx] = *(col).(*interface{})
		}
		ret = append(ret, temp)
	}
	return ret, nil
}

// 设置爬取队列的信息
func (this *MysqlProjClientImp) SetPCBody(userid_timest string, value *entity.PCBreakStruct) (int64, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return -1, err
	}
	sql := `insert INTO pc_body_msg (pc_body_msg_user_id, pc_body_msg_body) values(?,?)`
	res, err := this.doSQL(sql, userid_timest, string(val))
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// 提取爬取队列的信息
func (this *MysqlProjClientImp) GetPCBody(userid_timest string) (*entity.PCBreakStruct, error) {
	sql := `select pc_body_msg_body from pachong.pc_body_msg where pc_body_msg_user_id=?`
	res, err := this.doQuery(sql, userid_timest)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	jsonstr := res[0][0].([]byte)

	ret := new(entity.PCBreakStruct)
	err = json.Unmarshal(jsonstr, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil

}

// 获取已读消息
func (this *MysqlProjClientImp) GetUserSubMsgNoRead(userid string) (*entity.UserSubMsgStruct, error) {
	ret := new(entity.UserSubMsgStruct)
	sql := `SELECT user_sub_msg_readed_msg FROM pachong.user_sub_msg_read where user_sub_msg_read_userid=?`
	res, err := this.doQuery(sql, userid)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return ret, json.Unmarshal(res[0][0].([]byte), ret)
}

// 获取未读消息
func (this *MysqlProjClientImp) GetUserSubMsgReaded(userid string) (*entity.UserSubMsgStruct, error) {
	ret := new(entity.UserSubMsgStruct)
	sql := `SELECT user_sub_msg_no_read_msg FROM pachong.user_sub_msg_read where user_sub_msg_read_userid=?`
	res, err := this.doQuery(sql, userid)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return ret, json.Unmarshal(res[0][0].([]byte), ret)
}

func (this *MysqlProjClientImp) InsertUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, err
	}
	empty, err := json.Marshal(new(entity.EmptyStruct))
	if err != nil {
		return -1, err
	}
	sql := `insert INTO pachong.user_sub_msg_read (user_sub_msg_read_userid,user_sub_msg_no_read_msg,user_sub_msg_readed_msg) values(?,?,?)`
	ret, err := this.doSQL(sql, userid, string(res), empty)
	if err != nil {
		return -1, err
	}
	return ret.RowsAffected()
}
func (this *MysqlProjClientImp) UpdateUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, err
	}
	sql := `update user_sub_msg_read set user_sub_msg_readed_msg=? where user_sub_msg_read_userid=?`
	ret, err := this.doSQL(sql, string(res), userid)
	if err != nil {
		return -1, err
	}
	return ret.RowsAffected()
}
func (this *MysqlProjClientImp) InsertUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, err
	}
	empty, err := json.Marshal(new(entity.EmptyStruct))
	if err != nil {
		return -1, err
	}
	sql := `insert INTO pachong.user_sub_msg_read (user_sub_msg_read_userid,user_sub_msg_no_read_msg,user_sub_msg_readed_msg) values(?,?,?)`
	ret, err := this.doSQL(sql, userid, empty, string(res))
	if err != nil {
		return -1, err
	}
	return ret.RowsAffected()
}
func (this *MysqlProjClientImp) UpdateUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, err
	}
	sql := `update user_sub_msg_read set user_sub_msg_no_read_msg=? where user_sub_msg_read_userid=?`
	ret, err := this.doSQL(sql, string(res), userid)
	if err != nil {
		return -1, err
	}
	return ret.RowsAffected()
}

// 关闭mysql
func (this *MysqlProjClientImp) Close() {
	this.client.Close()
}
