// file create by daihao, time is 2018/4/16 11:14
package utils

import "fmt"

// 通过获取待爬队列的信息
// reids为list
// 键为时间字符串
func GetWaitPCQueueKey() string {
	return fmt.Sprintf("waitPCQueue")
}

