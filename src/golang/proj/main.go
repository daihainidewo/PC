// file create by daihao
package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	"regexp"
	"strings"
	"io/ioutil"
)

const DATE_FORMAT = "200601"

func LastMonth() string {
	day := time.Now()
	//w := day.Weekday()
	//sunday = 0
	firstDay := day.AddDate(0, -1, 0)
	y, m, _ := firstDay.Date()

	return strconv.Itoa(y*100 + int(m))
}

type AA struct {
	A int `json:"a"`
	B int `json:"b"`
}

type BB struct {
	C int `json:"c"`
	D int `json:"d"`
	AA
}

func trimHtml(src string) string {
	//fmt.Println(time.Now())
	//将HTML标签全转换成小写
	//re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	//src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//将网页小写
	src = strings.ToLower(src)
	//去除STYLE
	re, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

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
func downloadHtml(url string) string {
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
func main() {
	fmt.Println(time.Now().Unix())
	url := "https://blog.csdn.net/jcjc918/article/list/4"
	html := downloadHtml(url)
	//html := `<a class="" href="/jcjc918?t=1" for="original">只看原创</a>`
	html = trimHtml(html)
	re := regexp.MustCompile(`(?s:<a(.*?)href="(.*?)"(.*?)>(.*?)</a>)`)
	res := re.FindAllStringSubmatch(html, -1)
	fmt.Println(len(res), res[2])
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Value:   "some value",
		Expires: time.Now().Add(1 * time.Minute),
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	ret, _ := json.Marshal(time.Now())
	fmt.Fprintln(w, string(ret))
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	ret, _ := json.Marshal(c)
	fmt.Fprintln(w, "YOUR COOKIE:")

	fmt.Fprintln(w, string(ret))
}
