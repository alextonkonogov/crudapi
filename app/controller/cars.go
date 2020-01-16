package controller

import (
	"encoding/json"
	"fmt"
	"github.com/alextonkonogov/crudapi/app/model"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"strings"

	"html/template"
	"net/http"
	"path/filepath"
)

func GetCars(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	cars, err := model.GetAllCars()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//указываем пути к файлам с шаблонами
	main := filepath.Join("public", "html", "cars.html")
	common := filepath.Join("public", "html", "common.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
	err = tmpl.ExecuteTemplate(rw, "cars", cars)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func AddCar(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	firmId := r.FormValue("firmId")
	markId := r.FormValue("markId")
	litresVolume := r.FormValue("litresVolume")

	if firmId == "" || markId == "" || litresVolume == "" {
		http.Error(rw, "Все значения должны быть заполнены", 400)
		return
	}

	car, err := model.NewCar(firmId, markId, litresVolume)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = car.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Автомобиль успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func DeleteCar(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	carId := p.ByName("carId")

	car, err := model.GetCarById(carId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = car.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Автомобиль был успешно удален")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateCar(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	carId := p.ByName("carId")
	firmIdStr := r.FormValue("firmId")
	markIdStr := r.FormValue("markId")
	litresVolumeStr := r.FormValue("litresVolume")

	firmId, err := strconv.Atoi(firmIdStr)
	if err != nil {
		return
	}
	markId, err := strconv.Atoi(markIdStr)
	if err != nil {
		return
	}

	litresVolume, err := strconv.ParseFloat(strings.TrimSpace(litresVolumeStr), 64)
	if err != nil {
		return
	}

	car, err := model.GetCarById(carId)
	if err != nil {
		fmt.Println("here")
		http.Error(rw, err.Error(), 400)
		return
	}

	car.FirmId = firmId
	car.MarkId = markId
	car.LitresVolume = litresVolume

	err = car.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Автомобиль был успешно изменен")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
