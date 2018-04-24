// file create by daihao, time is 2018/4/13 17:19
package utils

import (
	"fmt"
	"net/http"
	"time"
)

type HtmlCookie struct {
	Cookie *http.Cookie
}

func NewHtmlCookie() *HtmlCookie {
	ret := &HtmlCookie{Cookie: new(http.Cookie)}
	return ret
}

func (this *HtmlCookie) SetCookie(w http.ResponseWriter, name, value string, expires time.Duration) {
	//go func() {
	this.Cookie.Name = name
	this.Cookie.Value = value
	this.Cookie.Expires = time.Now().Add(expires)
	http.SetCookie(w, this.Cookie)
	//}()
}

func (this *HtmlCookie) ReadCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if fmt.Sprintf("%s", err) == "http: named cookie not present" { // cookie不存在
		return "", nil
	} else if err != nil {
		return "", err
	}
	return c.Value, nil
}

func (this *HtmlCookie) IsExistCookie(r *http.Request, name string) (bool, error) {
	_, err := r.Cookie(name)
	if fmt.Sprintf("%s", err) == "http: named cookie not present" {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
