// file create by daihao, time is 2018/4/12 16:16
package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"golang/utils"
	"time"
	"golang/service"
	"strings"
)

func indexNoLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get index no login", time.Now())
	ret := `<a href="/user/login">用户请登录</a>`
	w.Write([]byte(ret))
	return
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userLogin", w)
	fmt.Println("get /user/login", time.Now())
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("passwd")
	// 检查用户键值对是否正确
	fmt.Println(username, password)

	// 获取用户id
	userid := "123"

	// 设置cookie
	utils.Htmlcookie.SetCookie(w, "userid", userid, utils.COOKIEEXPIRE)

	// 返回主页
	index(w, r)
}

func userZhuce(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /user/zhuce", time.Now())
	r.ParseForm()
	//fmt.Println(r.URL)
	//fmt.Println(r.Form)
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	userid := fmt.Sprintf("%d", time.Now().Unix())
	// 存入mysql中
	fmt.Println(username, passwd)
	// 设置cookie
	utils.Htmlcookie.SetCookie(w, "userid", userid, utils.COOKIEEXPIRE)
	// 返回主页
	index(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /index", time.Now())

	cookie, err1 := utils.Htmlcookie.ReadCookie(r, "userid")
	if err1 != nil {
		fmt.Println(err1)
	}
	if cookie == "" { // 没有cookie，用户需要登录
		indexNoLogin(w, r)
		return
	}

	r.ParseForm()
	//user := r.PostForm.Get("user")
	//fmt.Println(user)
	fp, err := os.Open(utils.GetUrl("index.html"))
	if err != nil {
		fmt.Println(err)
	}
	html, err1 := ioutil.ReadAll(fp)
	if err1 != nil {
		fmt.Println(err1)
	}
	w.Write(html)
}

func userSub(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /userSub", time.Now())
	r.ParseForm()
	userid := r.Form.Get("userid")
	suburl := r.Form.Get("suburl")
	keyword := r.Form.Get("keyword")
	token := r.Form.Get("token")
	site := r.Form.Get("site")
	titlekw := r.Form.Get("titlekeyword")
	if userid == "" || suburl == "" || keyword == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "")
		w.Write(res)
		return
	}
	titlekeyword := strings.Split(titlekw, ",")
	err := service.WWWService.SetUserSubMsg(userid, suburl, keyword, site, token, titlekeyword)
	if err != nil {
		if err != nil {
			res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "")
			w.Write(res)
			return
		}
	}
	res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "")
	w.Write(res)
}
