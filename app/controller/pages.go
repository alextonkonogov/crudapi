package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//path := filepath.Join("public", "html", "startPage.html")
	//
	//tmpl, err := template.ParseFiles(path)
	//if err != nil {
	//	http.Error(rw, err.Error(), 400)
	//	return
	//}
	//
	//err = tmpl.Execute(rw, nil)
	//if err != nil {
	//	http.Error(rw, err.Error(), 400)
	//	return
	//}

	fmt.Fprint(rw, "Hello from app!")
}

//func GetCars(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	//получаем список всех пользователей
//	cars, err := car.GetAllCars()
//	if err != nil {
//		http.Error(rw, err.Error(), 400)
//		return
//	}
//
//	//указываем пути к файлам с шаблонами
//	main := filepath.Join("public", "html", "cars.html")
//	common := filepath.Join("public", "html", "common.html")
//
//	//создаем html-шаблон
//	tmpl, err := template.ParseFiles(main, common)
//	if err != nil {
//		http.Error(rw, err.Error(), 400)
//		return
//	}
//
//	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
//	err = tmpl.ExecuteTemplate(rw, "cars", cars)
//	if err != nil {
//		http.Error(rw, err.Error(), 400)
//		return
//	}
//}
