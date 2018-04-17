// file create by daihao, time is 2018/4/16 11:02
package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

type MysqlClientImp struct {
	client *sql.DB
}

func NewWWWMysqlClient(driverName, dataSourceName string) *MysqlClientImp {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil
	}
	return &MysqlClientImp{client: db}
}

func (this *MysqlClientImp) doSQL(sql string, args ...interface{}) (sql.Result, error) {
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

func (this *MysqlClientImp) doQuery(sql string, args ...interface{}) ([][]interface{}, error) {
	//stmt, err := this.client.Prepare(sql)
	//if err != nil {
	//	return nil, err
	//}
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
func (this *MysqlClientImp) SetPCBody(userid, value string) (int64, error) {
	sql := `INSERT INTO pc_body_msg (pc_body_msg_user_id, pc_body_msg_body) VALUES(?,?)`
	res, err := this.doSQL(sql, userid, value)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// 提取爬取队列的信息
func (this *MysqlClientImp) GetPCBody() ([][]string, error) {
	ret := make([][]string, 0)
	temp := make([]string, 2)
	sql := `SELECT pc_body_msg_user_id,pc_body_msg_body FROM pachong.pc_body_msg`
	res, err := this.doQuery(sql)
	if err != nil {
		return nil, err
	}
	for _, r := range res {
		temp[0], temp[1] = string(r[0].([]byte)), string(r[1].([]byte))
		ret = append(ret, temp)
	}

	return ret, nil
}

// 设置用户订阅
func (this *MysqlClientImp) SetUserSubMsg(userid, value string) (int64, error) {
	sql := `insert into user_sub (user_sub_user_id, user_sub_sub_msg) value (?, ?)`
	res, err := this.doSQL(sql, userid, value)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// 获取用户订阅
func (this *MysqlClientImp) GetUserSubMsg(userid string) (string, error) {
	sql := fmt.Sprintf(`SELECT user_sub_sub_msg FROM pachong.user_sub where user_sub_user_id=%s`, userid)
	res, err := this.doQuery(sql)
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		return "", nil
	}
	ret := string(res[0][0].([]byte))

	return ret, nil
}

// 设置该订阅有哪些用户
func (this *MysqlClientImp) SetSubUserMsg(submsg, userids string) (int64, error) {
	sql := `insert into pc_sub_user (pc_sub_user_sub, pc_sub_user_ids) values (?,?)`
	res, err := this.doSQL(sql, submsg, userids)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

// 获取该订阅有哪些用些
func (this *MysqlClientImp) GetSubUserMsg(submsg string) (string, error) {
	sql := `select pc_sub_user_ids from pachong.pc_sub_user where pc_sub_user_sub=?`
	res, err := this.doQuery(sql, submsg)
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		return "", nil
	}
	ret := string(res[0][0].([]byte))

	return ret, nil
}

// 关闭mysql
func (this *MysqlClientImp) Close() {
	this.client.Close()
}
