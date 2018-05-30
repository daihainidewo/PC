// file create by daihao, time is 2018/4/13 17:11
package utils

import (
	"encoding/json"
	"fmt"
	"golang/conf"
	"golang/entity"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
	"os"
	"crypto/md5"
	"encoding/hex"
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
	if ret.LogPath != "" {
		f, err := os.Stat(ret.LogPath)
		if err != nil || !f.IsDir() {
			return nil, fmt.Errorf("请输入正确的日志文件夹路径")
		}
	}
	ret.LogLevel = strings.ToLower(ret.LogLevel)
	if ret.LogLevel == "" {
		ret.LogLevel = "debug"
	}
	if ret.LogLevel != "debug" && ret.LogLevel != "info" {
		return nil, fmt.Errorf("日志等级出错")
	}
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
	if ret.PaTime == 0 {
		ret.PaTime = 30
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

func GetCurLogPath(baseLogPath string) string {
	t := time.Now()
	dirpath := fmt.Sprintf("%s%s", baseLogPath, t.Format("/2006/01/02"))
	logpath := fmt.Sprintf("%s%s.logger", baseLogPath, t.Format("/2006/01/02/15"))
	if _, err := os.Stat(logpath); err == nil {
		return logpath
	} else {
		if _, err := os.Stat(dirpath); err != nil {
			err := os.MkdirAll(dirpath, 0755)
			if err != nil {
				fmt.Println("创建日志文件夹失败", err)
				return ""
			}
		}
	}
	return logpath
}
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}