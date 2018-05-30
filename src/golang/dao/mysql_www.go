// file create by daihao, time is 2018/4/16 11:02
package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"golang/entity"
	"strings"
	"golang/logger"
)

type MysqlWWWClientImp struct {
	client *sql.DB
}

func NewWWWMysqlClient(driverName, dataSourceName string) *MysqlWWWClientImp {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		logger.Println(err)
		return nil
	}
	return &MysqlWWWClientImp{client: db}
}

func (this *MysqlWWWClientImp) doSQL(sql string, args ...interface{}) (sql.Result, error) {
	stmt, e := this.client.Prepare(sql)
	if e != nil {
		return nil, fmt.Errorf("doSQL:client.Prepare sql=%s, error=%s", sql, e)
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("doSQL:client.Prepare sql=%s, args=%s, error=%s", sql, args, err)
	}
	return res, nil
}

func (this *MysqlWWWClientImp) doQuery(sql string, args ...interface{}) ([][]interface{}, error) {
	rows, err := this.client.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("doQuery:client.Query sql=%s, args=%s, error=%s", sql, args, err)
	}
	if rows == nil {
		return nil, nil
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("doQuery:rows.Columns error=%s", err)
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
			return nil, fmt.Errorf("doQuery:rows.Scan args=%s, error=%s", scanArgs, err)
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
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:InsertUserSubMsg:%s", err)
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
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:UpdateUserSubMsg:%s", err)
	}
	return res.RowsAffected()
}

// 获取用户订阅, 通过用户id去获取主播曾经订阅的内容
func (this *MysqlWWWClientImp) SelectUserSubMsg(userid string) (*[]entity.User2SubStruct, error) {
	sql := fmt.Sprintf(`SELECT user_sub_sub_msg FROM pachong.user_sub where user_sub_user_id=%s`, userid)
	res, err := this.doQuery(sql)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserSubMsg:%s", err)
	}
	if len(res) == 0 {
		return nil, nil
	}
	ret := new([]entity.User2SubStruct)
	err = json.Unmarshal(res[0][0].([]byte), ret)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserSubMsg:json.Unmarshal val=%s, error=%s", string(res[0][0].([]byte)), err)
	}
	return ret, nil
}

// 设置该订阅有哪些用户
func (this *MysqlWWWClientImp) InsertSubUserMsg(submsg, userids string) (int64, error) {
	sql := `insert into pc_sub_user (pc_sub_user_sub, pc_sub_user_ids) values (?,?)`
	res, err := this.doSQL(sql, submsg, userids)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:InsertSubUserMsg:%s", err)
	}
	return res.RowsAffected()
}

// 获取该订阅有哪些用些
func (this *MysqlWWWClientImp) SelectSubUserMsg(submsg string) (string, error) {
	sql := `select pc_sub_user_ids from pachong.pc_sub_user where pc_sub_user_sub=?`
	res, err := this.doQuery(sql, submsg)
	if err != nil {
		return "", fmt.Errorf("[Dao]MysqlWWWClientImp:SelectSubUserMsg:%s", err)
	}
	if len(res) == 0 {
		return "", nil
	}
	ret := string(res[0][0].([]byte))

	return ret, nil
}

// 获取已读消息
func (this *MysqlWWWClientImp) SelectUserSubMsgReaded(userid string) (*entity.UserSubMsgStruct, error) {
	ret := new(entity.UserSubMsgStruct)
	sql := `SELECT user_sub_msg_readed_msg FROM pachong.user_sub_msg_read where user_sub_msg_read_userid=?`
	res, err := this.doQuery(sql, userid)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserSubMsgReaded:%s", err)
	}
	if len(res) == 0 {
		return nil, nil
	}
	err = json.Unmarshal(res[0][0].([]byte), ret)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserSubMsgReaded:json.Unmarshal val=%s, error=%s", string(res[0][0].([]byte)), err)
	}
	return ret, nil
}

// 获取未读消息
func (this *MysqlWWWClientImp) SelectUserSubMsgNoRead(userid string) (*entity.UserSubMsgStruct, error) {
	ret := new(entity.UserSubMsgStruct)
	sql := `SELECT user_sub_msg_no_read_msg FROM pachong.user_sub_msg_read where user_sub_msg_read_userid=?`
	res, err := this.doQuery(sql, userid)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserSubMsgNoRead:%s", err)
	}
	if len(res) == 0 {
		return nil, nil
	}
	err = json.Unmarshal(res[0][0].([]byte), ret)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserSubMsgNoRead:json.Unmarshal val=%s, error=%s", string(res[0][0].([]byte)), err)
	}
	return ret, nil
}

// 插入用户信息
func (this *MysqlWWWClientImp) InsertUserMsg(userid, username, passwd string) (int64, error) {
	sql := `insert into user (userid, username, userpasswd) values(?,?,?)`
	res, err := this.doSQL(sql, userid, username, passwd)

	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate") {
			return 0, nil
		}
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:InsertUserMsg:%s", err)
	}
	return res.RowsAffected()
}

// 更改用户信息
func (this *MysqlWWWClientImp) UpdateUserMsgUsername(userid, username string) (int64, error) {
	sql := `update user set username=? where userid=?`
	res, err := this.doSQL(sql, username, userid)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:UpdateUserMsgUsername:%s", err)
	}
	return res.RowsAffected()
}
func (this *MysqlWWWClientImp) UpdateUserMsgUserpasswd(userid, userpasswd string) (int64, error) {
	sql := `update user set userpasswd=? where userid=?`
	res, err := this.doSQL(sql, userpasswd, userid)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:UpdateUserMsgUserpasswd:%s", err)
	}
	return res.RowsAffected()
}

// 获取用户信息
func (this *MysqlWWWClientImp) SelectUserMsg(userid string) (string, string, error) {
	sql := `SELECT username,userpasswd FROM pachong.user where userid=?`
	res, err := this.doQuery(sql, userid)
	if err != nil {
		return "", "", fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserMsg:%s", err)
	}
	if res == nil {
		return "", "", nil
	}
	username := string(res[0][0].([]byte))
	userpasswd := string(res[0][1].([]byte))
	return username, userpasswd, nil
}

// 查重用户名
func (this *MysqlWWWClientImp) SelectUserSameName(username string) (bool, error) {
	sql := `select userid from user where username=?`
	res, err := this.doQuery(sql, username)
	if err != nil {
		return false, fmt.Errorf("[Dao]MysqlWWWClientImp:SelectUserSameName:%s", err)
	}
	if len(res) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// 检测用户名和密码是否匹配
func (this *MysqlWWWClientImp) CheckUserNamePasswd(username, userpasswd string) (string, error) {
	sql := `select userid from user where username=? and userpasswd=?`
	res, err := this.doQuery(sql, username, userpasswd)
	if err != nil {
		return "", fmt.Errorf("[Dao]MysqlWWWClientImp:CheckUserNamePasswd:%s", err)
	}
	if res == nil {
		return "", nil
	}
	return string(res[0][0].([]byte)), nil
}

// 清空用户未读订阅消息
func (this *MysqlWWWClientImp) DelectUserSubMsgNoRead(userid string) (int64, error) {
	sql := `update user_sub_msg_read set user_sub_msg_no_read_msg='{}' where user_sub_msg_read_userid=?`
	res, err := this.doSQL(sql, userid)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:CheckUserNamePasswd:%s", err)
	}
	return res.RowsAffected()
}

// 更新用户已读消息
func (this *MysqlWWWClientImp) UpdateUserSubMsgReaded(val *entity.UserSubMsgStruct) (int64, error) {
	ret, err := json.Marshal(val)
	if err != nil {
		return -1, err
	}
	sql := `update user_sub_msg_read set user_sub_msg_readed_msg=? where user_sub_msg_read_userid=?`
	res, err := this.doSQL(sql, string(ret), val.Userid)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlWWWClientImp:UpdateUserSubMsgReaded:%s", err)
	}
	return res.RowsAffected()
}

// 关闭mysql
func (this *MysqlWWWClientImp) Close() {
	this.client.Close()
}
