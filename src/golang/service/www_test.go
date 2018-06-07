package service

import (
	"testing"
	"fmt"
	"golang/dao"
)

func startwww() {
	WWWService = NewWWWService()
	dao.MysqlWWWDao = dao.NewWWWMysqlClient("mysql", "root:Dai,1230@tcp(localhost:3306)/pachong?charset=utf8")
}
func TestNewWWWService_GetUserNoread(t *testing.T) {
	startwww()
	userid := "3512762103"
	ret, err := WWWService.GetUserNoread(userid)
	fmt.Println(ret, err)
}
