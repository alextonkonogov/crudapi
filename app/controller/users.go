package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/alextonkonogov/crudapi/app/model/user"
	"github.com/alextonkonogov/crudapi/app/server"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := user.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func Login(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	if login == "" || password == "" {
		err := errors.New("Необходимо указать логин и пароль!")
		http.Error(rw, err.Error(), 400)
		return
	}

	hash := md5.Sum([]byte(password))
	hashedPass := hex.EncodeToString(hash[:])

	user, err := user.Login(login, hashedPass)
	if err != nil {
		http.Error(rw, "Вы ввели неверный логин или пароль!", 400)
		return
	}

	//логин и пароль совпадают, поэтому генерируем токен, пишем его в кеш и в куки
	time64 := time.Now().Unix()
	timeInt := string(time64)
	token := login + password + timeInt

	hashToken := md5.Sum([]byte(token))
	hashedToken := hex.EncodeToString(hashToken[:])

	caches, err := server.GetCache()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	livingTime := 60 * time.Minute

	expiration := time.Now().Add(livingTime)
	caches.SetTTL(time.Duration(livingTime))
	caches.Set(hashedToken, user)

	cookie := http.Cookie{Name: "token", Value: url.QueryEscape(hashedToken), Expires: expiration}
	http.SetCookie(rw, &cookie)

	server.SetCookiesFromStruct(user, rw, livingTime)
}

func Logout(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	server.DeleteCookies(rw, r)
}

func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := strings.TrimSpace(r.FormValue("name"))
	surname := strings.TrimSpace(r.FormValue("surname"))
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))

	if name == "" || surname == "" || login == "" || password == "" {
		http.Error(rw, "Все поля должны быть заполнены", 400)
		return
	}

	user := user.NewUser(login, password, name, surname)

	err := user.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")

	user, err := user.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = user.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Пользователь был успешно удален")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")
	name := strings.TrimSpace(r.FormValue("name"))
	surname := strings.TrimSpace(r.FormValue("surname"))
	login := strings.TrimSpace(r.FormValue("login"))

	user, err := user.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	user.Name = name
	user.Surname = surname
	user.Login = login

	err = user.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Пользователь был успешно изменен")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
