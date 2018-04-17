// file create by daihao, time is 2018/4/16 11:02
package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func NewProjMysqlClient(driverName, dataSourceName string) *MysqlClientImp {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	return &MysqlClientImp{client: db}
}
