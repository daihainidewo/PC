// file create by daihao, time is 2018/4/16 11:02
package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"golang/entity"
	"golang/logger"
	"regexp"
)

type MysqlProjClientImp struct {
	client *sql.DB
}

func NewProjMysqlClient(driverName, dataSourceName string) *MysqlProjClientImp {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		logger.Println(err)
		return nil
	}
	return &MysqlProjClientImp{client: db}
}

func (this *MysqlProjClientImp) doSQL(sql string, args ...interface{}) (sql.Result, error) {
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

func (this *MysqlProjClientImp) doQuery(sql string, args ...interface{}) ([][]interface{}, error) {
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

// 设置爬取队列的信息
func (this *MysqlProjClientImp) InsertPCBody(userid_timest string, value *entity.PCBreakStruct) (int64, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:json.Marshal val=%s, error=%s", value, err)
	}
	sql := `insert INTO pc_body_msg (pc_body_msg_user_id, pc_body_msg_body) values(?,?)`
	res, err := this.doSQL(sql, userid_timest, string(val))
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:%s", err)
	}
	return res.RowsAffected()
}

// 提取爬取队列的信息
func (this *MysqlProjClientImp) SelectPCBody(userid_timest string) (*entity.PCBreakStruct, error) {
	sql := `select pc_body_msg_body from pachong.pc_body_msg where pc_body_msg_user_id=?`
	res, err := this.doQuery(sql, userid_timest)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlProjClientImp:SelectPCBody:%s", err)
	}
	if len(res) == 0 {
		return nil, nil
	}
	jsonstr := res[0][0].([]byte)

	ret := new(entity.PCBreakStruct)
	err = json.Unmarshal(jsonstr, ret)
	if err != nil {
		return nil, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:json.Unmarshal val=%s, error=%s", string(jsonstr), err)
	}
	return ret, nil

}

func (this *MysqlProjClientImp) InsertUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:json.Marshal val=%s, error=%s", value, err)
	}
	empty, err := json.Marshal(new(entity.EmptyStruct))
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:json.Marshal val=%s, error=%s", new(entity.EmptyStruct), err)
	}
	sql := `insert INTO pachong.user_sub_msg_read (user_sub_msg_read_userid,user_sub_msg_no_read_msg,user_sub_msg_readed_msg) values(?,?,?)`
	ret, err := this.doSQL(sql, userid, empty, string(res))
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertUserSubMsgReaded:%s", err)
	}
	return ret.RowsAffected()
}
func (this *MysqlProjClientImp) UpdateUserSubMsgReaded(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:json.Marshal val=%s, error=%s", value, err)
	}
	sql := `update user_sub_msg_read set user_sub_msg_readed_msg=? where user_sub_msg_read_userid=?`
	ret, err := this.doSQL(sql, string(res), userid)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:UpdateUserSubMsgReaded:%s", err)
	}
	return ret.RowsAffected()
}
func (this *MysqlProjClientImp) InsertUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:json.Marshal val=%s, error=%s", value, err)
	}
	empty, err := json.Marshal(new(entity.EmptyStruct))
	if err != nil {
		return -1, err
	}
	sql := `insert INTO pachong.user_sub_msg_read (user_sub_msg_read_userid,user_sub_msg_no_read_msg,user_sub_msg_readed_msg) values(?,?,?)`
	ret, err := this.doSQL(sql, userid, string(res), empty)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertUserSubMsgNoRead:%s", err)
	}
	return ret.RowsAffected()
}
func (this *MysqlProjClientImp) UpdateUserSubMsgNoRead(userid string, value *entity.UserSubMsgStruct) (int64, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertPCBody:json.Marshal val=%s, error=%s", value, err)
	}
	sql := `update user_sub_msg_read set user_sub_msg_no_read_msg=? where user_sub_msg_read_userid=?`
	ret, err := this.doSQL(sql, string(res), userid)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:UpdateUserSubMsgNoRead:%s", err)
	}
	return ret.RowsAffected()
}

func (this *MysqlProjClientImp) DeletePCBody(userid_timest string) (int64, error) {
	sql := `DELETE FROM pc_body_msg WHERE pc_body_msg_user_id=?`
	res, err := this.doSQL(sql, userid_timest)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:DeletePCBody:%s", err)
	}
	return res.RowsAffected()
}

// 关闭mysql
func (this *MysqlProjClientImp) Close() {
	this.client.Close()
}

// 新版本，将数据存取改成新的表
// 向表中添加获取的订阅消息
func (this *MysqlProjClientImp) InsertUserSubMsg(suburl string, value *entity.UserSubMsgStruct) (int64, error) {
	if value.SubMsg == nil || len(value.SubMsg) == 0 {
		return 0, nil
	}
	sql := `insert into user_sub_msg (user_sub_msg_user_id, user_sub_msg_user_msg, user_sub_msg_user_sub) values `
	re1, _ := regexp.Compile(`'`)
	re2, _ := regexp.Compile(`"`)
	for _, d := range value.SubMsg {
		d.Title = re1.ReplaceAllString(d.Title, `-`)
		d.Title = re2.ReplaceAllString(d.Title, `-`)
		s, _ := json.Marshal(d)
		sql = sql + `('` + value.Userid + `','` + string(s) + `','` + suburl + `'),`
	}
	sql = sql[:len(sql)-1]
	res, err := this.doSQL(sql)
	if err != nil {
		return -1, fmt.Errorf("[Dao]MysqlProjClientImp:InsertUserSubMsg:%s", err)
	}
	return res.RowsAffected()
}
