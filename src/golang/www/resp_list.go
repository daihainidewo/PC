// file create by daihao, time is 2018/4/12 16:16
package main

import (
	"golang/service"
	"golang/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"golang/logger"
)

func indexNoLogin(w http.ResponseWriter, r *http.Request) {
	logger.Println("get index no login", time.Now())
	ret := `<a href="/user/login">用户请登录</a>`
	w.Write([]byte(ret))
	return
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userLogin", w)
	logger.Println("get /user/login", time.Now())
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("passwd")
	// 验证用户正确性
	userid, err := service.WWWService.CheckUser(username, password)
	if err != nil {
		logger.Println(err)
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}
	if userid == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "密码错误")
		w.Write(res)
		return
	}
	logger.Println(username, password)

	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], userid)
	w.Write(res)
}

func userZhuce(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userZhuce", w)
	logger.Println("get /user/zhuce", time.Now())
	r.ParseForm()
	username := r.PostForm.Get("username")
	passwd := r.PostForm.Get("password")
	// 存入mysql中
	logger.Println(username, passwd)
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
	defer errorReport("index", w)
	logger.Println("get /index", time.Now())
	//
	//cookie, err1 := utils.Htmlcookie.ReadCookie(r, "userid")
	//if err1 != nil {
	//	logger.Println(err1)
	//}
	//if cookie == "" { // 没有cookie，用户需要登录
	//	indexNoLogin(w, r)
	//	return
	//}

	r.ParseForm()
	//user := r.PostForm.Get("user")
	//logger.Println(user)
	fp, err := os.Open("html\\index.html")
	if err != nil {
		logger.Println(err)
	}
	html, err1 := ioutil.ReadAll(fp)
	if err1 != nil {
		logger.Println(err1)
	}
	w.Write(html)
}

func userSub(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userSub", w)
	logger.Println("get /userSub", time.Now())
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
		logger.Println(err)
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}
	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "设置成功")
	w.Write(res)
}

func userGetSub(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userGetSub", w)
	logger.Println("get /userGetSub", time.Now())
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
	defer errorReport("userReaded", w)
	logger.Println("get /userReaded", time.Now())
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
	defer errorReport("userNoread", w)
	logger.Println("get /userNoread", time.Now())
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

func userReadMsg(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userReadMsg", w)
	logger.Println("/user/readmsg", time.Now())
	r.ParseForm()
	userid := r.Form.Get("userid")
	if userid == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数")
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserReadMsg(userid)
	if err != nil {
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}
	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret)
	w.Write(res)
}
