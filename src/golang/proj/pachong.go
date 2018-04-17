// file create by daihao, time is 2018/4/8 10:26
package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
	"golang/utils"
	"strings"
	"time"
	"sync"
	"golang/entity"
	"golang/dao"
)

// 网页下载器
func DownloadHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	content, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)
		return ""
	}
	html := string(content)
	return html
}

// 网页标题是否存在标题关键字组
func TitleContainKeyWord(title string, titlekeyword []string) bool {
	for _, word := range titlekeyword {
		if !strings.Contains(title, word) {
			return false
		}
	}
	return true
}

// 解析网页
func ParseHtml(passtoken entity.PageSiteTokeStruct, msg entity.PageTitleStruct, html, keyword string, titleKeyword []string) []entity.PageTitleStruct {
	temp := new(entity.PageTitleStruct)
	protocol := strings.Split(msg.URL, "/")[0] // 协议：http或者https
	ret := make([]entity.PageTitleStruct, 0)

	re := regexp.MustCompile(`(?is:<title>(.*?)</title>)`)
	res := re.FindAllStringSubmatch(html, 1)
	//fmt.Println(msg.URL)
	if len(res) == 0 {
		return nil
	}
	temp.Title = strings.TrimSpace(res[0][1])

	html = TrimHtml(html) // 去除网页标签，将字母全部小写

	// 标题含关键字直接订阅
	if TitleContainKeyWord(strings.ToLower(temp.Title), titleKeyword) || strings.Count(html, keyword) >= utils.SUBSCRIBENUM {
		// 这个是用户订阅的网页
		if strings.Contains(msg.URL, "category") || strings.Contains(msg.URL, "month") || strings.Contains(msg.URL, "page") {

		} else {
			fmt.Println("这个是用户订阅的网页，title:", temp.Title, "url:", msg.URL)
		}
	}

	utils.Countsm.Lock()
	utils.Count++
	utils.Countsm.Unlock()

	re = regexp.MustCompile(`(?s:<a(.*?)href="(.*?)"(.*?)>(.*?)</a>)`)
	res = re.FindAllStringSubmatch(html, -1)
	for _, link := range res {
		if len(link[2]) < 2 {
			continue
		}
		if link[2][0] == '#' {
			continue
		}
		if len(link[2]) > 11 && link[2][0:11] == "javascript:" {
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
				if !strings.Contains(link[2], ":") {
					//if len(link[2]) > 7 && link[2][0:7] != "mailto:" && link[2][0:4] != "ftp:" {
					fmt.Println("非法的URL：", link[2], "原网页url：", msg.URL)
				}
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
	//fmt.Println("end",len(ret))
	return ret
}

// 处理网页无关信息
func TrimHtml(src string) string {
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
	//fmt.Println(time.Now())

	return strings.TrimSpace(src)
}

// 启动爬虫
func StartPC(url, keyword, site, token string, userid int, titleKeyword []string) {
	starttoken := entity.PageTitleStruct{"网站首页", url}
	usersub := entity.UserSubStruct{Keyword: keyword}
	passtoken := entity.PageSiteTokeStruct{Site: site, Token: token}
	utils.SubUserMap[usersub] = append(make([]int, 0), userid)
	utils.UserSubMap[userid] = append(make([]entity.UserSubStruct, 0), usersub)
	utils.PageTitleMap[starttoken.URL] = starttoken.Title
	utils.PageTitleList.PushBack(starttoken)
	ch := make(chan struct{}, utils.PROJECTNUM)
	countsm := new(sync.Mutex)
	for i := 0; i < utils.PROJECTNUM; i++ {
		go func() {
			defer func() {
				ch <- struct{}{}
			}()
			scancount := 1 // 协程退出标志
			for {
				countsm.Lock()
				ele := utils.PageTitleList.Front()
				if ele == nil || ele.Value == nil {
					//fmt.Println("ele is nil", utils.PageTitleList.Len())
					time.Sleep(utils.NONEDATASLEEPTIME)
					countsm.Unlock()
					continue
				}
				data := (ele.Value.(entity.PageTitleStruct))
				utils.PageTitleList.Remove(ele)
				countsm.Unlock()

				html := DownloadHtml(data.URL)
				if html == "" {
					continue
				}
				ret := ParseHtml(passtoken, data, html, keyword, titleKeyword)

				for _, r := range ret {

					utils.PageSM.Lock()
					if _, ok := (utils.PageTitleMap)[r.URL]; ok {
						utils.Repacecount++
						utils.PageSM.Unlock()
						continue
					}
					//countsm.Lock()
					utils.PageTitleList.PushBack(r)
					//countsm.Unlock()
					(utils.PageTitleMap)[r.URL] = r.Title
					utils.PageSM.Unlock()
				}
				scancount++
				if scancount > utils.PACOUNT {
					//fmt.Println("break one")
					break
				}
			}
		}()
	}
	for i := 0; i < utils.PROJECTNUM; i++ {
		<-ch
	}
	// 当目标网页小于当前预估网页时，程序会阻塞
	//time.Sleep(1 * time.Second)
}

// 切换爬虫
func CutovePC() {
	// 从redis中获取现在需要爬取的相关信息
	key := utils.GetWaitPCQueueKey()
	val := dao.RedisCacheDao.GetPCBody(key)
	if val == "" {
		CutovePC()
	}
	// 调用爬取程序

	// 存放当前信息去redis排队

	// 存放相关信息去mysql
}
