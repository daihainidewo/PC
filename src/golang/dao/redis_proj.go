// file create by daihao, time is 2018/4/16 10:57
package dao

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"golang/entity"
)

type RedisCacheImp struct {
	client *redis.Client
}

func NewRedisCache(addr, passwd string, db int) *RedisCacheImp {
	client := new(RedisCacheImp)
	client.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd, // no password set
		DB:       db,     // use default DB
	})
	return client
}

/*
存放获取相关数据至redis进行爬虫爬取排队：
key : 需要操作的键（redis队列的键）
val : 存放排队信息
*/
// 获取爬虫的主体信息
func (this *RedisCacheImp) GetPCBodyMsg(key string) (*entity.PCQueueStruct, error) {
	ret := new(entity.PCQueueStruct)
	val, err := this.client.LPop(key).Result() // 只返回
	if err != nil || val == "" {
		return nil, fmt.Errorf("[Dao]RedisCacheImp:GetPCBodyMsg:client.LPop key=%s, error=%s", key, err)
	}
	err = json.Unmarshal([]byte(val), ret)
	if err != nil {
		return nil, fmt.Errorf("[Dao]RedisCacheImp:GetPCBodyMsg:json.Unmarshal data=%s, error=%s", val, err)
	}
	return ret, nil
}

// 设置爬虫的主体信息
func (this *RedisCacheImp) SetPCBodyMsg(key string, val *entity.PCQueueStruct) error {
	ret, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("[Dao]RedisCacheImp:SetPCBodyMsg:json.Marshal data=%s, error=%s", val, err)
	}
	return fmt.Errorf("[Dao]RedisCacheImp:SetPCBodyMsg:client.RPush key=%s, val=%s, error=%s", key, string(ret), this.client.RPush(key, string(ret)).Err())
}

// 关闭redis
func (this *RedisCacheImp) Close() {
	this.client.Close()
}
