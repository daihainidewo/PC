// file create by daihao, time is 2018/4/13 17:11
package utils

import (
	"fmt"
	"golang/entity"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"golang/conf"
)

func ParseURL(u ...string) ([]string, error) {
	ret := make([]string, 0)
	for _, l := range u {
		ur, err := url.Parse(l)
		if err != nil {
			return nil, err
		}
		ret = append(ret, ur.Path)
	}
	return ret, nil
}

func ReadConf(path string) (*conf.ConfStruct, error) {
	ret := new(conf.ConfStruct)

	confstring, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
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
	if ret.SubscribeNum == 0 {
		ret.SubscribeNum = 50
	}
	if ret.ProjectNum == 0 {
		ret.ProjectNum = 1000
	}
	if ret.PaCount == 0 {
		ret.PaCount = 20
	}
	if ret.CookieExpire == 0 {
		ret.CookieExpire = 5 * 60
	}
	if ret.NoneDataSleepTime == 0 {
		ret.NoneDataSleepTime = 30
	}
	return ret, nil
}

func User2SubStructIsEqual(a, b entity.User2SubStruct) bool {
	if a.Keyword == b.Keyword && a.URL == b.URL && a.Site == b.Site && a.Token == b.Token {
		if len(a.TitleKeyWord) != len(b.TitleKeyWord) {
			return false
		}

		if (a.TitleKeyWord == nil) != (b.TitleKeyWord == nil) {
			return false
		}

		for i, v := range a.TitleKeyWord {
			if v != b.TitleKeyWord[i] {
				return false
			}
		}
		return true
	}
	return false
}
