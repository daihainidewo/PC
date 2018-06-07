// file create by daihao, time is 2018/4/17 11:23
package dao

import (
	"testing"
	"golang/entity"
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
)

func startTest() {
	MysqlWWWDao = NewWWWMysqlClient("mysql", "root:Dai,1230@tcp(localhost:3306)/pachong?charset=utf8")
}

func TestMysqlWWWClientImp_InsertUserSubMsg(t *testing.T) {
	startTest()
	userid := fmt.Sprintf("%d", time.Now().UnixNano())
	fmt.Println(userid)
	fmt.Println(userid[:7])
	fmt.Println(userid[7:17])
	data := make([]entity.User2SubStruct, 0)
	MysqlWWWDao.InsertUserSubMsg("1234", &data)
	MysqlWWWDao.Close()
}
func TestMysqlWWWClientImp_CheckUserNamePasswd(t *testing.T) {
	startTest()
	username := "daihao"
	passwd := "1234"
	id, err := MysqlWWWDao.CheckUserNamePasswd(username, passwd)
	fmt.Println(id, err)
}

func TestMysqlWWWClientImp_SelectUserSubMsgReaded(t *testing.T) {
	startTest()
	userid := "3512762103"
	us, err := MysqlWWWDao.SelectUserSubMsgReaded(userid)
	fmt.Println(us, err)
}
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
func TestMysqlWWWClientImp_Close(t *testing.T) {
	a := make([]string, 0)
	fmt.Println(a == nil, a)
}

func TestMysqlWWWClientImp_SelectUserSubMsgNoread(t *testing.T) {
	startTest()
	userid := "3512762103"
	ret, err := MysqlWWWDao.SelectUserSubMsgNoread(userid)
	fmt.Println(ret, err)
}

func TestMysqlWWWClientImp_UpdateUserSubMsgRead(t *testing.T) {
	startTest()
	userid := "1234"
	submsg := new(entity.PageTitleStruct)
	submsg.Title = "biaoti"
	submsg.URL = "http://www.baidu.com"
	res, err := MysqlWWWDao.UpdateUserSubMsgRead(userid, submsg)
	fmt.Println(res, err)
}
