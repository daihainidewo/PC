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
	"golang/entity"
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
	username := r.Form.Get("username")
	password := r.Form.Get("passwd")
	callback := r.Form.Get("callback")
	// 验证用户正确性
	userid, err := service.WWWService.CheckUser(username, password)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}
	logger.LogPrintln(username, password, callback)

	if userid == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "密码错误", callback)
		w.Write(res)
		return
	}
	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], userid, callback)
	w.Write(res)
}
func userCheckName(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userCheckName", w)
	logger.LogPrintln("get /user/checkname")
	r.ParseForm()
	username := r.Form.Get("username")
	callback := r.Form.Get("callback")
	// 存入mysql中
	userid, err := service.WWWService.CheckUserName(username)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}
	if userid == "" {
		res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "用户名重复", callback)
		w.Write(res)
		return
	}
	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "用户名可用", callback)
	w.Write(res)
	return
}

func userZhuce(w http.ResponseWriter, r *http.Request) {
	defer errorReport("userZhuce", w)
	logger.LogPrintln("get /user/zhuce")
	r.ParseForm()
	username := r.Form.Get("username")
	// 使用md5加密算法
	passwd := utils.MD5(r.Form.Get("password"))
	callback := r.Form.Get("callback")
	// 存入mysql中
	userid, err := service.WWWService.SetUserMsg(username, passwd)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}
	if userid == "" {
		res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "用户名重复", callback)
		w.Write(res)
		return
	}
	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], userid, callback)
	w.Write(res)
}

func index(w http.ResponseWriter, r *http.Request) {
	defer errorReport("index", w)
	logger.LogPrintln("get /index")
	r.ParseForm()
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
	suburl := r.Form.Get("url")
	keyword := r.Form.Get("keyword")
	token := r.Form.Get("token")
	titlekw := r.Form.Get("titlekeyword")
	callback := r.Form.Get("callback")
	// 获取主站域名
	urlarr := strings.Split(suburl, "/")
	site := urlarr[0] + "//" + urlarr[2]
	if userid == "" || suburl == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "参数错误", callback)
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
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}
	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "设置成功", callback)
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
	callback := r.Form.Get("callback")
	if userid == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数", callback)
		w.Write(res)
		return
	}
	ret, err := service.WWWService.GetUserReadMsg(userid)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}
	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], ret, callback)
	w.Write(res)
}

func userDelSub(w http.ResponseWriter, r *http.Request) {
	logger.LogPrintln("/user/test")
	r.ParseForm()
	userid := r.Form.Get("userid")
	suburl := r.Form.Get("url")
	keyword := r.Form.Get("keyword")
	token := r.Form.Get("token")
	titlekw := r.Form.Get("titlekeyword")
	callback := r.Form.Get("callback")
	if userid == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数", callback)
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
	err := service.WWWService.DelUserSub(userid, suburl, keyword, token, titlekeyword)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}
	res := utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "", callback)
	w.Write(res)
}

func userTest(w http.ResponseWriter, r *http.Request) {
	logger.LogPrintln("/user/test")
	r.ParseForm()
	callback := r.Form.Get("callback")
	logger.LogPrintln(callback)
	w.Write(utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "测试成功", callback))
}
func userRead(w http.ResponseWriter, r *http.Request) {
	logger.LogPrintln("/user/read")
	r.ParseForm()
	callback := r.Form.Get("callback")
	userid := r.Form.Get("userid")
	title := r.Form.Get("title")
	url := r.Form.Get("url")
	if userid == "" || url == "" {
		res := utils.RespFormat(utils.INVALID_PARAMS, utils.RespMsg[utils.INVALID_PARAMS], "非法参数", callback)
		w.Write(res)
		return
	}
	val := new(entity.PageTitleStruct)
	val.URL = url
	val.Title = title
	err := service.WWWService.ReadUserSubMsg(userid, val)
	if err != nil {
		logger.ErrPrintln(err)
		res := utils.RespFormat(utils.SYSTEM_ERROR, utils.RespMsg[utils.SYSTEM_ERROR], "系统错误", callback)
		w.Write(res)
		return
	}
	w.Write(utils.RespFormat(utils.SUCCESS, utils.RespMsg[utils.SUCCESS], "阅读成功", callback))
}
