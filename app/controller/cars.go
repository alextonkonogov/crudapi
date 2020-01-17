package controller

import (
	"encoding/json"
	"github.com/alextonkonogov/crudapi/app/model/car"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"strings"

	"net/http"
)

func GetCars(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cars, err := car.GetAllCars()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode(cars)
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

	car, err := car.NewCar(firmId, markId, litresVolume)
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

	car, err := car.GetCarById(carId)
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

	car, err := car.GetCarById(carId)
	if err != nil {
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
