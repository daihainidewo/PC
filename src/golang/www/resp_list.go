// file create by daihao, time is 2018/4/12 16:16
package main

import (
	"golang/service"
	"golang/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"golang/logger"
)

func indexNoLogin(w http.ResponseWriter, r *http.Request) {
	logger.LogPrintln("get index no login")
	ret := `<a href="/user/login">用户请登录</a>`
	w.Write([]byte(ret))
	return
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userLogin", w)
	logger.LogPrintln("get /user/login")
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("passwd")
	// 验证用户正确性
	userid, err := service.WWWService.CheckUser(username, password)
	if err != nil {
		logger.ErrPrintln(err)
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
	logger.LogPrintln("get /user/zhuce")
	r.ParseForm()
	username := r.PostForm.Get("username")
	passwd := r.PostForm.Get("password")
	// 存入mysql中
	logger.Println(username, passwd)
	userid, err := service.WWWService.SetUserMsg(username, passwd)
	if err != nil {
		logger.ErrPrintln(err)
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
	logger.LogPrintln("get /index")
	//
	//cookie, err1 := utils.Htmlcookie.ReadCookie(r, "userid")
	//if err1 != nil {
	//	logger.LogPrintln()(err1)
	//}
	//if cookie == "" { // 没有cookie，用户需要登录
	//	indexNoLogin(w, r)
	//	return
	//}

	r.ParseForm()
	//user := r.PostForm.Get("user")
	//logger.LogPrintln()(user)
	fp, err := os.Open("html\\index.html")
	if err != nil {
		logger.ErrPrintln(err)
	}
	html, err1 := ioutil.ReadAll(fp)
	if err1 != nil {
		logger.ErrPrintln(err1)
	}
	w.Write(html)
}

func userSub(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userSub", w)
	logger.LogPrintln("get /userSub")
	r.ParseForm()
	userid := r.Form.Get("userid")
	suburl := r.Form.Get("suburl")
	keyword := r.Form.Get("keyword")
	token := r.Form.Get("token")
	titlekw := r.Form.Get("titlekeyword")
	callback := r.Form.Get("callback")

	// 获取主站域名
	urlarr := strings.Split(suburl, "/")
	site := urlarr[0] + "//" + urlarr[2]
	if userid == "" || suburl == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "参数错误",callback)
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
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误",callback)
		w.Write(res)
		return
	}
	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "设置成功",callback)
	w.Write(res)
}

func userGetSub(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userGetSub", w)
	logger.LogPrintln("get /userGetSub")
	r.ParseForm()
	userid := r.Form.Get("userid")
	callback := r.Form.Get("callback")
	if userid == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数", callback)
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserSubMsg(userid)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}

	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret, callback)
	w.Write(res)
}

func userReaded(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userReaded", w)
	logger.LogPrintln("get /user/readed")
	r.ParseForm()
	userid := r.Form.Get("userid")
	callback := r.Form.Get("callback")
	if userid == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数", callback)
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserReaded(userid)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}

	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret, callback)
	w.Write(res)
}

func userNoread(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userNoread", w)
	logger.LogPrintln("/user/noread")
	r.ParseForm()
	userid := r.Form.Get("userid")
	callback := r.Form.Get("callback")
	if userid == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数", callback)
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserNoread(userid)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}

	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret, callback)
	w.Write(res)
}

func userReadMsg(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userReadMsg", w)
	logger.LogPrintln("/user/readmsg")
	r.ParseForm()
	userid := r.Form.Get("userid")
	if userid == "" {
		res := utils.RespJson(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数")
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserReadMsg(userid)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespJson(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误")
		w.Write(res)
		return
	}
	res := utils.RespJson(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret)
	w.Write(res)
}
