// file create by daihao, time is 2018/4/18 17:08
package dao

func newMysqlProjTest() {
	MysqlProjDao = NewProjMysqlClient("mysql", "root:DHdh,.1234@tcp(localhost:3306)/pachong?charset=utf8")
}
