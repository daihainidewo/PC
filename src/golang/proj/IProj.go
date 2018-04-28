// file create by daihao, time is 2018/4/19 9:59
package proj

type IPCService interface {
	CtrlPC()
	Close()
}

var PCService IPCService
