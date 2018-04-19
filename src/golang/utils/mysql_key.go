// file create by daihao, time is 2018/4/18 17:57
package utils

import (
	"golang/entity"
	"fmt"
	"strings"
)

func GetUserTimeMysqlKey(pcqs *entity.PCQueueStruct) string {
	return fmt.Sprintf("%s|%s", pcqs.Userid, pcqs.Timest)
}
func ParseUserTimeMysqlKey(key string) *entity.PCQueueStruct {
	ret := new(entity.PCQueueStruct)
	sp := strings.Split(key, "|")
	ret.Userid = sp[0]
	ret.Timest = sp[1]
	return ret
}
