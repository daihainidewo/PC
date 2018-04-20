// file create by daihao, time is 2018/4/19 18:45
package conf

/*
服务启动配置文件结构体
*/
// 配置文件信息
type ConfStruct struct {
	StartPort               int    `json:"start_port"`
	RedisAddr               string `json:"redis_addr"`
	RedisPasswd             string `json:"redis_passwd"`
	RedisDB                 int    `json:"redis_db"`
	MysqlProjDriverName     string `json:"mysql_proj_driver_name"`
	MysqlProjDataSourceName string `json:"mysql_proj_data_source_name"`
	MysqlWWWDriverName      string `json:"mysql_www_driver_name"`
	MysqlWWWDataSourceName  string `json:"mysql_www_data_source_name"`

	SubscribeNum      int   `json:"subscribe_num"`        // 最低订阅数
	ProjectNum        int   `json:"project_num"`          // 创建协程数
	PaTime            int64 `json:"pa_time"`              // 协程存活时间
	PaCount           int   `json:"pa_count"`             // 每个协程爬取数量
	NoneDataSleepTime int   `json:"none_data_sleep_time"` // 单位是微秒
	CookieExpire      int   `json:"cookie_expire"`        // 单位是秒
}
