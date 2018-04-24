// file create by daihao, time is 2018/4/17 11:23
package dao

import (
	"fmt"
	"testing"
)

func startTest() {
	MysqlWWWDao = NewWWWMysqlClient("mysql", "root:DHdh,.1234@tcp(localhost:3306)/pachong?charset=utf8")
}

//func TestMysqlClientImp_SetPCBody(t *testing.T) {
//	startTest()
//	MysqlWWWDao.SetPCBody("", "")
//}
//
//func TestMysqlClientImp_GetPCBody(t *testing.T) {
//	startTest()
//	MysqlWWWDao.GetPCBody()
//}
//
//func TestMysqlClientImp_SetUserSubMsg(t *testing.T) {
//	startTest()
//	i, err := MysqlWWWDao.SetUserSubMsg("dh", "value")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(i)
//}
//
//func TestMysqlClientImp_GetUserSubMsg(t *testing.T) {
//	startTest()
//	ret, err := MysqlWWWDao.GetUserSubMsg("3")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(ret)
//}
//
//func TestMysqlClientImp_SetSubUserMsg(t *testing.T) {
//	startTest()
//	ret, err := MysqlWWWDao.SetSubUserMsg("dhdhdh", "123")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(ret)
//}
//
//func TestMysqlClientImp_GetSubUserMsg(t *testing.T) {
//	startTest()
//	ret, err := MysqlWWWDao.GetSubUserMsg("dhdhdh")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(ret)
//	ret, err = MysqlProjDao.GetSubUserMsg("dh")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(ret)
//}

func TestMysqlWWWClientImp_InsertUserMsg(t *testing.T) {
	startTest()
	fmt.Println(MysqlWWWDao.InsertUserMsg("1234", "dhdh", "dhdhdhdh"))
}
