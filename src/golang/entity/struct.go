// file create by daihao, time is 2018/4/8 10:45
package entity

type PageTitleStruct struct {
	Title string // 网页标题
	URL   string // 网页url
}

type PageSiteTokeStruct struct {
	Site  string // 主站网址
	Token string // 主站token，www.xxm.com, token xxm，（可以改进成www.xxm，需要改代码)
}

type UserSubStruct struct {
	PageSiteTokeStruct
	Keyword string // 搜索关键字
}


