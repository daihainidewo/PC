// file create by daihao, time is 2018/4/8 10:26
package proj

import (
	"golang/entity"
	"golang/utils"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
	"golang/logger"
	"golang/service"
	"encoding/json"
)

type PC struct {
}

func NewPCInit() *PC {
	return new(PC)
}

// 处理网页无关信息
func (this *PC) trimHtml(src string) string {
	//将HTML标签全转换成小写
	var re *regexp.Regexp

	//去除注释
	re, _ = regexp.Compile("<!--[\\S\\s]+?-->")
	src = re.ReplaceAllString(src, "")

	//将网页小写
	src = strings.ToLower(src)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// 将转移字符转义回来
	re, _ = regexp.Compile("&lt;")
	src = re.ReplaceAllString(src, "<")
	re, _ = regexp.Compile("&gt;")
	src = re.ReplaceAllString(src, ">")
	re, _ = regexp.Compile("&quot;")
	src = re.ReplaceAllString(src, `"`)

	re, _ = regexp.Compile("<([\\s]*?)/a([\\s]*?)>")
	src = re.ReplaceAllString(src, "<a>")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("<([\\s]*?)[^a][\\S\\s]*?>")
	src = re.ReplaceAllString(src, "\n")

	re, _ = regexp.Compile("<a>")
	src = re.ReplaceAllString(src, "</a>")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	//logger.Println(time.Now())

	return strings.TrimSpace(src)
}

// 网页下载器
func (this *PC) downloadHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logger.Println(err)
		return ""
	}
	defer resp.Body.Close()
	content, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		logger.Println(err1)
		return ""
	}
	html := string(content)
	return html
}

// 网页标题是否存在标题关键字组
func (this *PC) titleContainKeyWord(title string, titlekeyword []string) bool {
	for _, word := range titlekeyword {
		if !strings.Contains(title, word) {
			return false
		}
	}
	return true
}

// 解析网页
func (this *PC) parseHtml(passtoken entity.PageSiteTokeStruct, msg entity.PageTitleStruct, html, keyword string, titleKeyword []string) []entity.PageTitleStruct {
	temp := new(entity.PageTitleStruct)
	protocol := strings.Split(msg.URL, "/")[0] // 协议：http或者https
	ret := make([]entity.PageTitleStruct, 0)

	re := regexp.MustCompile(`(?is:<title>(.*?)</title>)`)
	res := re.FindAllStringSubmatch(html, 1)
	//logger.Println(msg.URL)
	if len(res) == 0 {
		return nil
	}
	temp.Title = strings.TrimSpace(res[0][1])

	html = this.trimHtml(html) // 去除网页标签，将字母全部小写

	// 标题含关键字直接订阅
	if this.titleContainKeyWord(strings.ToLower(temp.Title), titleKeyword) || strings.Count(html, keyword) >= utils.SUBSCRIBENUM {
		// 这个是用户订阅的网页，过滤掉某些url所包含的关键字
		if strings.Contains(msg.URL, "category") || strings.Contains(msg.URL, "month") || strings.Contains(msg.URL, "page") {
			// 这里可以做成自定义的
		} else {
			//logger.Println("这个是用户订阅的网页，title:", temp.Title, "url:", msg.URL)
			tempret := entity.PageTitleStruct{Title: temp.Title, URL: msg.URL}
			utils.UserSubUrl = append(utils.UserSubUrl, tempret)
		}
	}

	re = regexp.MustCompile(`(?s:<a(.*?)href="(.*?)"(.*?)>(.*?)</a>)`)
	res = re.FindAllStringSubmatch(html, -1)

	for _, link := range res {
		if len(link[2]) < 2 {
			continue
		}
		if link[2][0] == '#' {
			continue
		}

		if len(link[2]) > 4 && link[2][0:4] == "http" {
			temp.URL = link[2] // http://www.baidu.com
		} else if link[2][0] == '/' {
			if link[2][1] == '/' {
				temp.URL = protocol + link[2] // //www.baidu.com
			} else {
				temp.URL = passtoken.Site + link[2] // /a.html
			}
		} else { // 字母开头的url
			if link[2][0:2] == "./" || strings.Contains(link[2], ".htm") {
				wz := 0
				for idx, char := range msg.URL {
					if char == '/' {
						wz = idx
					}
				}
				temp.URL = msg.URL[0:wz+1] + link[2]
			} else if link[2][0] == '?' {
				temp.URL = msg.URL + link[2]
			} else {
				if strings.Contains(link[2], ":") {
					continue
				}
				logger.Println("非法的URL：", link[2], "原网页url：", msg.URL)
				continue
			}
		}
		if !strings.Contains(temp.URL, passtoken.Token) {
			continue
		}
		temp.URL = strings.Split(temp.URL, "#")[0]
		temp.URL = strings.Split(temp.URL, ";")[0]
		temp.URL = strings.Split(temp.URL, "?")[0]
		ret = append(ret, entity.PageTitleStruct{"", temp.URL})
	}
	//logger.Println("end",len(ret))
	return ret
}

// 启动爬虫
func (this *PC) startPC(url, keyword, site, token, userid string, titleKeyword []string) {
	passtoken := entity.PageSiteTokeStruct{Site: site, Token: token}
	ch := make(chan struct{}, utils.PROJECTNUM)
	countsm := new(sync.Mutex)
	for i := 0; i < utils.PROJECTNUM; i++ {
		go func() {
			defer func() {
				ch <- struct{}{}
			}()
			scancount := 1 // 协程退出标志
			starttime := time.Now().Unix()
			for {
				if time.Now().Unix()-starttime > utils.PATIME {
					//logger.Println("break one")
					break
				}
				countsm.Lock()
				ele := utils.PageTitleList.Front()
				if ele == nil || ele.Value == nil {
					//logger.Println("ele is nil", utils.PageTitleList.Len())
					time.Sleep(utils.NONEDATASLEEPTIME)
					countsm.Unlock()
					continue
				}
				//fmt.Println((ele.Value).(string))
				data := ele.Value.(entity.PageTitleStruct)
				utils.PageTitleList.Remove(ele)
				countsm.Unlock()
				html := this.downloadHtml(data.URL)
				if html == "" {
					continue
				}
				ret := this.parseHtml(passtoken, data, html, keyword, titleKeyword)
				for _, r := range ret {

					utils.PageSM.Lock()
					if _, ok := (utils.PageTitleMap)[r.URL]; ok {
						utils.Repacecount++
						utils.PageSM.Unlock()
						continue
					}
					utils.PageTitleList.PushBack(r)
					//logger.DebugPrintln("set one", r.URL)
					(utils.PageTitleMap)[r.URL] = r.Title
					utils.PageSM.Unlock()
				}
				logger.Println(utils.PageTitleList.Len())
				scancount++
				if scancount > utils.PACOUNT {
					//logger.Println("break one")
					break
				}
			}
		}()
	}
	for i := 0; i < utils.PROJECTNUM; i++ {
		<-ch
	}
}

func (this *PC) CtrlPC() {
	ch := make(chan int, 1)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			logger.LogPrintln("the next pc")
			utils.UserSubUrl = make([]entity.PageTitleStruct, 0)
			// 准备下一个爬虫
			userid, pcbs, err := service.ProjService.StartNextPC()
			if err != nil {
				if !strings.Contains(err.Error(), "redis: nil") {
					logger.ErrPrintln(err)
				}
				continue
			}
			logger.LogPrintln("PC ing ...")
			this.startPC(pcbs.URL, pcbs.Keyword, pcbs.Site, pcbs.Token, userid, pcbs.TitleKeyWord)
			if userid == "" {
				logger.LogPrintln("userid is nil")
				continue
			}
			_, err = service.ProjService.SetUserSubMsgNoRead(userid, utils.UserSubUrl)
			if err != nil {
				logger.ErrPrintln(err)
				continue
			}
			//logger.LogPrintln(utils.PageTitleList.Len())
			// 将爬虫存放进爬取队列
			pcbs.PageTitleMap = utils.PageTitleMap
			pcbs.PageTitleList2Slice = make([]string, utils.PageTitleList.Len())
			iter := utils.PageTitleList.Front()
			for i := 0; i < utils.PageTitleList.Len(); i++ {
				t, _ := json.Marshal(iter.Value.(entity.PageTitleStruct))
				pcbs.PageTitleList2Slice[i] = string(t)
				iter = iter.Next()
			}
			//logger.Println()
			err = service.ProjService.SetPCBody(userid, pcbs)
			if err != nil {
				logger.ErrPrintln(err)
				continue
			}
			//logger.Println("map len", len(pcbs.PageTitleMap))
			//logger.Println("sleep...")
			//time.Sleep(1 * time.Minute)
		}
	}()
	<-ch
}

func (this *PC) Close() {

}
