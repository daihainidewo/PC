// file create by daihao, time is 2018/4/13 17:04
package utils

import "time"

const (
	COOKIEEXPIRE = 30 * time.Second // cookie过期时间
)

var HtmlPath string        // 静态网页代码存放地
var Htmlcookie *HtmlCookie // 网页cookie
