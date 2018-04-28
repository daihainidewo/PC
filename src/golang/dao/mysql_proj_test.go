// file create by daihao, time is 2018/4/18 17:08
package dao

import (
	"testing"
	"fmt"
)

func newMysqlProjTest() {
	MysqlProjDao = NewProjMysqlClient("mysql", "root:DHdh,.1234@tcp(localhost:3306)/pachong?charset=utf8")
}

func TestMysqlProjClientImp_DeletePCBody(t *testing.T) {
	newMysqlProjTest()
	fmt.Println(MysqlProjDao.DeletePCBody("1234|1524811586275469600"))
}
