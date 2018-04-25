// file create by daihao, time is 2018/4/18 10:38
package dao

import (
	"fmt"
	"golang/entity"
	"testing"
	"time"
)

func start() {
	RedisCacheDao = NewRedisCache("localhost:6379", "", 0)

}

func TestRedisCacheImp_GetPCBody(t *testing.T) {
	start()
	val := new(entity.PCQueueStruct)
	val.Timest = fmt.Sprintf("%d", time.Now().UnixNano())
	val.Userid = "123421"
	err := RedisCacheDao.SetPCBody("key", val)
	if err != nil {
		fmt.Println(err)
	}

	val, err1 := RedisCacheDao.GetPCBody("key")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(*val)
}
