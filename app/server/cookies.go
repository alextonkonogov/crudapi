package server

import (
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"net/http"
	"net/url"
	"time"
)

func SetCookiesFromStruct(strct interface{}, rw http.ResponseWriter, livingTime time.Duration) {
	mp := structs.Map(strct)
	for k, v := range mp {
		if k == "Login" || k == "Password" {
			continue
		}
		vStr := fmt.Sprintf("%v", v)
		expiration := time.Now().Add(livingTime)
		cookie := http.Cookie{Name: k, Value: url.QueryEscape(vStr), Expires: expiration}
		http.SetCookie(rw, &cookie)
	}
}

func ReadCookie(name string, r *http.Request) (value string, err error) {
	if name == "" {
		return value, errors.New("You are trying to read empty cookie")
	}
	cookie, err := r.Cookie(name)
	if err != nil {
		return value, err
	}
	str := cookie.Value
	value, _ = url.QueryUnescape(str)
	return value, err
}

func DeleteCookies(rw http.ResponseWriter, r *http.Request) {
	for _, v := range r.Cookies() {
		c := http.Cookie{
			Name:   v.Name,
			MaxAge: -1}
		http.SetCookie(rw, &c)
	}
}
