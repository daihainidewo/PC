// file create by daihao, time is 2018/4/18 17:08
package dao

import (
	"testing"
	"fmt"
	"golang/entity"
)

func newMysqlProjTest() {
	MysqlProjDao = NewProjMysqlClient("mysql", "root:Dai,1230@tcp(localhost:3306)/pachong?charset=utf8")
}

func TestMysqlProjClientImp_DeletePCBody(t *testing.T) {
	newMysqlProjTest()
	fmt.Println(MysqlProjDao.DeletePCBody("1234|1524811586275469600"))
}
func TestMysqlProjClientImp_InsertUserSubMsg(t *testing.T) {
	newMysqlProjTest()
	value := new(entity.UserSubMsgStruct)
	value.Userid = "1234"
	value.SubMsg = make([]entity.PageTitleStruct, 0)
	//value.SubMsg = append(value.SubMsg, entity.PageTitleStruct{"biaoti", "http://www.baidu.com"})
	//value.SubMsg = append(value.SubMsg, entity.PageTitleStruct{"biaoti", "http://www.baidu.com"})
	//value.SubMsg = append(value.SubMsg, entity.PageTitleStruct{"biaoti", "http://www.baidu.com"})
	row, err := MysqlProjDao.InsertUserSubMsg("用户订阅url", value)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)
}
