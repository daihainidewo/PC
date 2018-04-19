// file create by daihao, time is 2018/4/19 9:59
package proj

type IPCService interface {
	StartPC(url, keyword, site, token, userid string, titleKeyword []string)
}

var PCService IPCService
