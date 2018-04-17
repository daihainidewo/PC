// file create by daihao, time is 2018/4/12 19:40
package utils

import "sync"

var Count int           // 数量统计
var Repacecount int64   // url重复的次数
var Runcount int        //当前运行的协程
var Countsm *sync.Mutex // 数字锁
