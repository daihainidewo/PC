// file create by daihao, time is 2018/4/17 17:50
package entity

/*
用户订阅结构
例如：
url := "http://域名/文件夹/子文件夹/文件/a.html" 	// 搜索起始url
keyword := "go"                                 // 全文搜索关键字
token := "//域名/文件夹/"                     	// 控制网站搜索域
site := "http://域名"                          	// 搜索站点
titleKeyword := make([]string, 0)               // 标题关键字
*/
type User2SubStruct struct {
	URL          string   `json:"url"`          // 搜索起始url
	Keyword      string   `json:"keyword"`      // 全文搜索关键字
	Token        string   `json:"token"`        // 控制网站搜索域
	Site         string   `json:"site"`         // 搜索站点
	TitleKeyWord []string `json:"titlekeyword"` // 标题关键字
}

/*
爬虫中断结构体
PageTitleMap为查重url的map，日后会放在redis中
*/
type PCBreakStruct struct {
	User2SubStruct
	PageTitleMap        map[string]string `json:"pagetitlemap"`
	PageTitleList2Slice []string          `json:"pagetitleslice"`
}

/*
爬虫中断redis排队结构体
userid : 用户
timest : 时间戳，纳秒（键值唯一标志）
mysqlkey : 爬取过程中的去重map和待爬取网页队列的mysql键值，现改为userid|timest
*/
type PCQueueStruct struct {
	Userid string `json:"userid"`
	Timest string `json:"timest"`
	//Mysqlkey string `json:"mysqlkey"`
}

/*

*/
// 配置文件信息
type ConfStruct struct {
	RedisAddr               string `json:"redis_addr"`
	RedisPasswd             string `json:"redis_passwd"`
	RedisDB                 int    `json:"redis_db"`
	MysqlProjDriverName     string `json:"mysql_proj_driver_name"`
	MysqlProjDataSourceName string `json:"mysql_proj_data_source_name"`
	MysqlWWWDriverName      string `json:"mysql_www_driver_name"`
	MysqlWWWDataSourceName  string `json:"mysql_www_data_source_name"`
}
