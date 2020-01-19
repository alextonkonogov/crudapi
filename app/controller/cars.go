package controller

import (
	"encoding/json"
	"strconv"
	"strings"

	"net/http"

	"github.com/alextonkonogov/crudapi/app/model/car"
	"github.com/julienschmidt/httprouter"
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
	mark := r.FormValue("mark")
	litresVolume := r.FormValue("litresVolume")

	if firmId == "" || mark == "" || litresVolume == "" {
		http.Error(rw, "Все значения должны быть заполнены", 400)
		return
	}

	car, err := car.NewCar(firmId, mark, litresVolume)
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
	mark := r.FormValue("mark")
	litresVolumeStr := r.FormValue("litresVolume")

	firmId, err := strconv.Atoi(firmIdStr)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	litresVolume, err := strconv.ParseFloat(strings.TrimSpace(litresVolumeStr), 64)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	car, err := car.GetCarById(carId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	car.FirmId = firmId
	car.Mark = mark
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
