// file create by daihao, time is 2018/4/16 16:17
package utils

import "encoding/json"

const (
	SUCCESS        = 0
	SYSTEM_ERROR   = 1001
	INVALID_PARAMS = 7000
)

var RespMsg = map[int]string{
	0:    "success",
	1001: "system error",
	2002: "no data",
	7000: "invalid params",
}

type Result struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func RespJson(errno int, errmsg string, data interface{}) []byte {
	var result = new(Result)
	result.Errno = errno
	result.Errmsg = errmsg
	result.Data = data
	res, _ := json.Marshal(result)
	return res
}

// 处理jsonp
func RespFormat(errno int, errmsg string, data interface{}, callback string) []byte {
	res := RespJson(errno, errmsg, data)
	if callback != "" {
		res = []byte(callback + "(" + string(res) + ")")
	}
	return res
}
