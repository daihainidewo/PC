// file create by daihao, time is 2018/4/13 17:11
package utils

import (
	"fmt"
	"golang/entity"
	"encoding/json"
	"io/ioutil"
)

func GetUrl(path string) string {
	return fmt.Sprintf("%s%s", HtmlPath, path)
}

var defaultConf = `
{
	"redis_addr":"127.0.0.1:6379",
	"redis_passwd":"",
	"redis_db":0,

	"mysql_proj_driver_name":"mysql",
	"mysql_proj_data_source_name":"root:@tcp(localhost:3306)/pachong?charset=utf8",

	"mysql_www_driver_name":"mysql",
	"mysql_www_data_source_name":"root:DHdh,.1234@tcp(localhost:3306)/pachong?charset=utf8"
}
`

func ReadConf(path string) (*entity.ConfStruct, error) {
	ret := new(entity.ConfStruct)
	//fmt.Println(path)
	var confstring []byte
	var err error
	if path == "" {
		confstring = []byte(defaultConf)
	} else {
		confstring, err = ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
	}
	err = json.Unmarshal([]byte(confstring), ret)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(confstring))
	if ret.RedisAddr == "" {
		ret.RedisAddr = "localhost:6379"
	}
	if ret.MysqlProjDriverName == "" {
		ret.MysqlProjDriverName = "msyql"
	}
	if ret.MysqlProjDataSourceName == "" {
		return nil, fmt.Errorf("请输入mysql数据源的相关信息，建议格式（name:password@tcp(ip:port)/database?args=args）")
	}
	if ret.MysqlWWWDriverName == "" {
		ret.MysqlProjDriverName = "msyql"
	}
	if ret.MysqlWWWDataSourceName == "" {
		return nil, fmt.Errorf("请输入mysql数据源的相关信息，建议格式（name:password@tcp(ip:port)/database?args=args）")
	}
	return ret, nil
}
