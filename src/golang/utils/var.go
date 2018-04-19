// file create by daihao, time is 2018/4/8 10:51
package utils

import (
	"sync"
	"container/list"
	"time"
	"golang/entity"
)

const (
	SUBSCRIBENUM      = 50                    // 订阅数限制，大于此则认为是用户的订阅网页
	PROJECTNUM        = 1000                  // 创建协程数
	PACOUNT           = 20                    // 每个协程最大爬取数
	NONEDATASLEEPTIME = 10 * time.Microsecond // 无数据消费休眠时间
	//TOKEN        = "xingximing"               // 主站token
	//SITE         = "http://www.xingximing.cn" // 主站网址
)

var PageTitleMap map[string]string // 查重map

var PageTitleChan chan entity.PageTitleStruct // 协程间传送数据管道

var PageTitleList *list.List // 协程间传送数据链表

var PageLimitProNum *TokenBucket // 限制消费者大小

var PageLimitScanNum *TokenBucket // 限制生产者大小

var PageSM *sync.Mutex // 读写PageTitleMap锁

var UserSubMap map[string][]entity.UserSubStruct // 用户订阅map，key：用户id，value：用户订阅结构体，用户映射订阅

var SubUserMap map[entity.UserSubStruct][]string // 订阅映射用户

var TempPageTitleSlice []entity.PageTitleStruct // 临时存放钱
