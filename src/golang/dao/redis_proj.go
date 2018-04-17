// file create by daihao, time is 2018/4/16 10:57
package dao

import (
	"github.com/go-redis/redis"
	"fmt"
)

type RedisCacheImp struct {
	client *redis.Client
}

func NewProjRedis(opt string) *RedisCacheImp {
	client := new(RedisCacheImp)
	client.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func (this *RedisCacheImp) GetPCBody(key string) string {
	ret, err := this.client.LPop(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func (this *RedisCacheImp) SetPCBody(key, val string) error {
	return this.client.RPush(key, val).Err()
}
