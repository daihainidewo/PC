// file create by daihao, time is 2018/4/16 11:02
package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"golang/entity"
	"encoding/json"
)

type MysqlWWWClientImp struct {
	client *sql.DB
}

func NewWWWMysqlClient(driverName, dataSourceName string) *MysqlWWWClientImp {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil
	}
	return &MysqlWWWClientImp{client: db}
}

func (this *MysqlWWWClientImp) doSQL(sql string, args ...interface{}) (sql.Result, error) {
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

func (this *MysqlWWWClientImp) doQuery(sql string, args ...interface{}) ([][]interface{}, error) {
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

// 设置用户订阅
func (this *MysqlWWWClientImp) InsertUserSubMsg(idkey string, value *[]entity.User2SubStruct) (int64, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return -1, nil
	}
	sql := `insert into user_sub (user_sub_user_id, user_sub_sub_msg) value (?, ?)`
	res, err := this.doSQL(sql, idkey, string(val))
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}
func (this *MysqlWWWClientImp) UpdateUserSubMsg(idkey string, value *[]entity.User2SubStruct) (int64, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return -1, nil
	}
	sql := `update user_sub set user_sub_sub_msg=? where user_sub_user_id=?`
	res, err := this.doSQL(sql, string(val), idkey)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// 获取用户订阅, 通过用户id去获取主播曾经订阅的内容
func (this *MysqlWWWClientImp) SelectUserSubMsg(userid string) (*[]entity.User2SubStruct, error) {
	sql := fmt.Sprintf(`SELECT user_sub_sub_msg FROM pachong.user_sub where user_sub_user_id=%s`, userid)
	res, err := this.doQuery(sql)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	ret := new([]entity.User2SubStruct)
	err = json.Unmarshal(res[0][0].([]byte), ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// 设置该订阅有哪些用户
func (this *MysqlWWWClientImp) InsertSubUserMsg(submsg, userids string) (int64, error) {
	sql := `insert into pc_sub_user (pc_sub_user_sub, pc_sub_user_ids) values (?,?)`
	res, err := this.doSQL(sql, submsg, userids)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// 获取该订阅有哪些用些
func (this *MysqlWWWClientImp) SelectSubUserMsg(submsg string) (string, error) {
	sql := `select pc_sub_user_ids from pachong.pc_sub_user where pc_sub_user_sub=?`
	res, err := this.doQuery(sql, submsg)
	if err != nil || len(res) == 0 {
		return "", err
	}
	ret := string(res[0][0].([]byte))

	return ret, nil
}

// 获取已读消息
func (this *MysqlWWWClientImp) SelectUserSubMsgReaded(userid string) (*entity.UserSubMsgStruct, error) {
	ret := new(entity.UserSubMsgStruct)
	sql := `SELECT user_sub_msg_readed_msg FROM pachong.user_sub_msg_read where user_sub_msg_read_userid=?`
	res, err := this.doQuery(sql, userid)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return ret, json.Unmarshal(res[0][0].([]byte), ret)
}

// 获取未读消息
func (this *MysqlWWWClientImp) SelectUserSubMsgNoRead(userid string) (*entity.UserSubMsgStruct, error) {
	ret := new(entity.UserSubMsgStruct)
	sql := `SELECT user_sub_msg_no_read_msg FROM pachong.user_sub_msg_read where user_sub_msg_read_userid=?`
	res, err := this.doQuery(sql, userid)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return ret, json.Unmarshal(res[0][0].([]byte), ret)
}

// 关闭mysql
func (this *MysqlWWWClientImp) Close() {
	this.client.Close()
}
