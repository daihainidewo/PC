// file create by daihao, time is 2018/4/8 10:51
package utils

import (
	"container/list"
	"golang/entity"
	"sync"
	"time"
)

// 爬虫变量

var (
	SUBSCRIBENUM       int                               // 订阅数限制，大于此则认为是用户的订阅网页
	PROJECTNUM         int                               // 创建协程数
	PATIME             int64                             // 每个协程爬取的时间
	PACOUNT            int                               // 每个协程最大爬取数
	NONEDATASLEEPTIME  time.Duration                     // 无数据消费休眠时间
	PageTitleMap       map[string]string                 // 查重map
	PageTitleChan      chan entity.PageTitleStruct       // 协程间传送数据管道
	PageTitleList      *list.List                        // 协程间传送数据链表
	PageLimitProNum    *TokenBucket                      // 限制消费者大小
	PageLimitScanNum   *TokenBucket                      // 限制生产者大小
	PageSM             *sync.Mutex                       // 读写PageTitleMap锁
	UserSubMap         map[string][]entity.UserSubStruct // 用户订阅map，key：用户id，value：用户订阅结构体，用户映射订阅
	SubUserMap         map[entity.UserSubStruct][]string // 订阅映射用户
	TempPageTitleSlice []entity.PageTitleStruct          // 临时存放序列
	UserSubUrl         []entity.PageTitleStruct          // 存放用户订阅的网页
)

// 网站变量
var (
	COOKIEEXPIRE time.Duration // cookie过期时间
	Htmlcookie   *HtmlCookie   // 网页cookie
)
