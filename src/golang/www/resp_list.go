// file create by daihao, time is 2018/4/12 16:16
package main

import (
	"fmt"
	"golang/service"
	"golang/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
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
	username := r.Form.Get("username")
	password := r.Form.Get("passwd")
	// 验证用户正确性
	userid, err := service.WWWService.CheckUser(username, password)
	if err != nil {
		fmt.Println(err)
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}
	if userid == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "密码错误")
		w.Write(res)
		return
	}
	fmt.Println(username, password)

	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], userid)
	w.Write(res)
}

func userZhuce(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /user/zhuce", time.Now())
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	// 存入mysql中
	fmt.Println(username, passwd)
	userid, err := service.WWWService.SetUserMsg(username, passwd)
	if err != nil {
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}
	if userid == "" {
		res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "用户名重复")
		w.Write(res)
		return
	}
	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], userid)
	w.Write(res)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /index", time.Now())
	//
	//cookie, err1 := utils.Htmlcookie.ReadCookie(r, "userid")
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//if cookie == "" { // 没有cookie，用户需要登录
	//	indexNoLogin(w, r)
	//	return
	//}

	r.ParseForm()
	//user := r.PostForm.Get("user")
	//fmt.Println(user)
	fp, err := os.Open("html\\index.html")
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
	if userid == "" || suburl == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "参数错误")
		w.Write(res)
		return
	}

	titlekey := strings.Split(titlekw, ",")
	titlekeyword := make([]string, 0)
	// 去除空值
	for _, v := range titlekey {
		if v == "" {
			continue
		}
		titlekeyword = append(titlekeyword, v)
	}
	err := service.WWWService.SetUserSubMsg(userid, suburl, keyword, site, token, titlekeyword)
	if err != nil {
		fmt.Println(err)
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}
	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "设置成功")
	w.Write(res)
}

func userGetSub(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /userGetSub", time.Now())
	r.ParseForm()
	userid := r.Form.Get("userid")
	if userid == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数")
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserSubMsg(userid)
	if err != nil {
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}

	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret)
	w.Write(res)
}

func userReaded(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /userReaded", time.Now())
	r.ParseForm()
	userid := r.Form.Get("userid")
	if userid == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数")
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserReaded(userid)
	if err != nil {
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}

	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret)
	w.Write(res)
}

func userNoread(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /userNoread", time.Now())
	r.ParseForm()
	userid := r.Form.Get("userid")
	if userid == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数")
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserNoread(userid)
	if err != nil {
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}

	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret)
	w.Write(res)
}
