// file create by daihao, time is 2018/4/8 11:19
package proj

import (
	"testing"
	"golang/utils"
	"sync"
	"fmt"
	"time"
	"strings"
	"container/list"
	"golang/entity"
)

func startPC() {
	utils.Count = 0
	utils.Repacecount = 1
	utils.Runcount = 0
	utils.Countsm = new(sync.Mutex)
	utils.PageTitleMap = make(map[string]string)
	utils.PageTitleChan = make(chan entity.PageTitleStruct, 100) // 管道
	utils.PageLimitProNum = utils.NewTokenBucket(50)             // 消费者
	utils.PageLimitScanNum = utils.NewTokenBucket(1000)          // 生产者
	utils.PageSM = new(sync.Mutex)
	utils.PageTitleList = list.New()
	utils.UserSubMap = make(map[string][]entity.UserSubStruct)
	utils.SubUserMap = make(map[entity.UserSubStruct][]string)
	fmt.Println("pro chan", utils.PageLimitProNum.Len(), "scan chan", utils.PageLimitScanNum.Len())
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now(), "map len", len(utils.PageTitleMap), "list len", utils.PageTitleList.Len(), "runcount", utils.Runcount, "count", utils.Count, "repacecount", utils.Repacecount)
		}
	}()
}

func TestDownloadHtml(t *testing.T) {
	startPC()
	url := "https://blog.csdn.net/jcjc918/article/list"
	html := DownloadHtml(url)
	//fmt.Println(html)
	fmt.Println(time.Now())
	for i := 0; i < 1000; i++ {
		TrimHtml(html)
	}
	fmt.Println(time.Now())
}

//func TestParseHtml(t *testing.T) {
//	startPC()
//	url := "http://www.xingximing.cn"
//	keyword := "java"
//	token := "xingximing"
//	userid := 123
//	starttoken := utils.PageTitleStruct{"网站首页", url}
//	usersub := utils.UserSubStruct{Keyword: keyword}
//	passtoken := utils.PageSiteTokeStruct{Site: url, Token: token}
//	utils.SubUserMap[usersub] = append(make([]int, 0), userid)
//	utils.UserSubMap[userid] = append(make([]utils.UserSubStruct, 0), usersub)
//	utils.PageTitleMap[starttoken.URL] = starttoken.Title
//	utils.PageTitleList.PushBack(starttoken)
//
//	for i := 0; i < utils.PROJECTNUM; i++ {
//		utils.Countsm.Lock()
//		utils.Runcount++
//		utils.Countsm.Unlock()
//		utils.PageLimitScanNum.Get()
//		go func() {
//			defer func() {
//				utils.Countsm.Lock()
//				utils.Runcount--
//				utils.Countsm.Unlock()
//				utils.PageLimitScanNum.Put()
//			}()
//			scancount := 0
//
//			//wg := new(sync.WaitGroup)
//			ele := utils.PageTitleList.Front()
//			utils.PageTitleList.Remove(ele)
//			//data := (ele.Value.(utils.PageTitleStruct))
//			for data := range utils.PageTitleChan {
//				ch := make(chan struct{}, 1)
//				html := DownloadHtml(data.URL)
//				if html == "" {
//					continue
//				}
//				//ret := ParseHtml(passtoken, data, html, keyword,)
//				//wg.Add(1)
//
//				go func(ret []utils.PageTitleStruct) {
//					utils.Countsm.Lock()
//					utils.Count++
//					utils.Countsm.Unlock()
//					defer func() {
//						utils.Countsm.Lock()
//						utils.Count--
//						utils.Countsm.Unlock()
//						ch <- struct{}{}
//						//fmt.Println(1)
//						//wg.Done()
//					}()
//					//fmt.Println(1)
//
//					for _, r := range ret {
//						r.URL = strings.Split(r.URL, "#")[0]
//						r.URL = strings.Split(r.URL, ";")[0]
//						utils.PageSM.Lock()
//						if _, ok := (utils.PageTitleMap)[r.URL]; ok {
//							utils.Repacecount++
//							utils.PageSM.Unlock()
//							continue
//						}
//						utils.PageSM.Unlock()
//						bj := false
//						for {
//							if bj {
//								break
//							}
//							select {
//							case utils.PageTitleChan <- r:
//								utils.PageSM.Lock()
//								(utils.PageTitleMap)[r.URL] = r.Title
//								utils.PageSM.Unlock()
//								bj = true
//								break
//							default:
//								//fmt.Println(scancount)
//								time.Sleep(10 * time.Millisecond)
//								break
//							}
//						}
//					}
//				}(ret)
//				fmt.Println(<-ch)
//				scancount++
//				if scancount > utils.PACOUNT {
//					fmt.Println("break one")
//					break
//				}
//				//fmt.Println("not break one")
//			}
//		}()
//
//	}
//	//defer utils.PageLimitProNum.Close()
//	defer utils.PageLimitScanNum.Close()
//}
//
//func TestParseHtml2(t *testing.T) {
//	startPC()
//	url := "https://blog.csdn.net/chenbaoke/article/details/42780895"
//	keyword := "go"
//	token := "blog.csdn.net"
//	userid := 123
//	starttoken := utils.PageTitleStruct{"网站首页", url}
//	usersub := utils.UserSubStruct{Keyword: keyword}
//	passtoken := utils.PageSiteTokeStruct{Site: url, Token: token}
//	utils.SubUserMap[usersub] = append(make([]int, 0), userid)
//	utils.UserSubMap[userid] = append(make([]utils.UserSubStruct, 0), usersub)
//	utils.PageTitleMap[starttoken.URL] = starttoken.Title
//	utils.PageTitleList.PushBack(starttoken)
//
//	ch := make(chan struct{}, utils.PROJECTNUM)
//	for i := 0; i < utils.PROJECTNUM; i++ {
//
//		utils.Countsm.Lock()
//		utils.Runcount++
//		utils.Countsm.Unlock()
//
//		go func() {
//
//			defer func() {
//				utils.Countsm.Lock()
//				utils.Runcount--
//				utils.Countsm.Unlock()
//			}()
//			scancount := 0
//
//			//wg := new(sync.WaitGroup)
//			for {
//				if utils.PageTitleList.Len() == 0 {
//					time.Sleep(10 * time.Microsecond)
//					continue
//				}
//				utils.Countsm.Lock()
//				ele := utils.PageTitleList.Front()
//				if ele == nil {
//					continue
//				}
//				data := (ele.Value.(utils.PageTitleStruct))
//				utils.PageTitleList.Remove(ele)
//				utils.Countsm.Unlock()
//
//				html := DownloadHtml(data.URL)
//				if html == "" {
//					continue
//				}
//				//ret := ParseHtml(passtoken, data, html, keyword)
//				utils.Countsm.Lock()
//				utils.Count++
//				utils.Countsm.Unlock()
//
//				for _, r := range ret {
//					r.URL = strings.Split(r.URL, "#")[0]
//					r.URL = strings.Split(r.URL, ";")[0]
//					utils.PageSM.Lock()
//					if _, ok := (utils.PageTitleMap)[r.URL]; ok {
//						utils.Countsm.Lock()
//						utils.Repacecount++
//						utils.Countsm.Unlock()
//						utils.PageSM.Unlock()
//						continue
//					}
//					utils.PageTitleList.PushBack(r)
//					(utils.PageTitleMap)[r.URL] = r.Title
//					utils.PageSM.Unlock()
//				}
//				//}()
//				scancount++
//				fmt.Println(scancount)
//				if scancount > utils.PACOUNT {
//					fmt.Println("break one")
//					break
//				}
//				utils.Countsm.Lock()
//				utils.Count--
//				utils.Countsm.Unlock()
//			}
//		}()
//
//	}
//	<-ch
//	defer utils.PageLimitScanNum.Close()
//}
//func TestDemo(t *testing.T) {
//	//startPC()
//	//url := "https://blog.csdn.net/jcjc918/article/list/4"
//	//keyword := ""
//	//token := "blog.csdn.net/jcjc918"
//	////userid := 123
//	//starttoken := utils.PageTitleStruct{"网站首页", url}
//	////usersub := utils.UserSubStruct{Keyword: keyword}
//	//passtoken := utils.PageSiteTokeStruct{Site: url, Token: token}
//	//html := DownloadHtml(starttoken.URL)
//	//ret := ParseHtml(passtoken, starttoken, html, keyword)
//	//fmt.Println("len ret", len(ret))
//	//for _, r := range ret {
//	//
//	//	if r.URL == "https://blog.csdn.net/jcjc918/article/details/9897703" {
//	//		fmt.Println("find it")
//	//	} else {
//	//		fmt.Println(r.URL)
//	//	}
//	//}
//}

func TestStartPC(t *testing.T) {
	startPC()
	url := "https://sustyuxiao.github.io/"  // 搜索起始url
	keyword := ""                           // 全文搜索关键字
	token := "sustyuxiao.github.io"         // 控制网站搜索域
	userid := 123                           // 用户id
	site := "https://sustyuxiao.github.io/" // 搜索站点
	titleKeyword := make([]string, 0)       // 标题关键字
	//titleKeyword = append(titleKeyword, "c++")
	//titleKeyword = append(titleKeyword, "面经")
	keyword = strings.ToLower(keyword)
	StartPC(url, keyword, site, token, userid, titleKeyword)
	fmt.Println(time.Now(), "map len", len(utils.PageTitleMap), "list len", utils.PageTitleList.Len(), "runcount", utils.Runcount, "count", utils.Count, "repacecount", utils.Repacecount)
}
