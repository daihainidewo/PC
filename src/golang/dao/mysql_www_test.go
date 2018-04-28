// file create by daihao, time is 2018/4/17 11:23
package dao

func startTest() {
	MysqlWWWDao = NewWWWMysqlClient("mysql", "root:DHdh,.1234@tcp(localhost:3306)/pachong?charset=utf8")
}
