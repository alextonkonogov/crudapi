package controller

import (
	"encoding/json"
	"github.com/alextonkonogov/crudapi/app/model/firm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func GetFirms(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	firms, err := firm.GetAllFirms()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode(firms)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func AddFirm(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	title := r.FormValue("title")

	if title == "" {
		http.Error(rw, "Укажите название фирмы", 400)
		return
	}

	firm := firm.NewFirm(title)

	err := firm.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Фирма успешно добавлена!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateFirm(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	firmId := p.ByName("firmId")
	title := strings.TrimSpace(p.ByName("title"))

	firm, err := firm.GetFirmById(firmId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	if title != "" {
		firm.Title = title
	}

	err = firm.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Фирма была успешно изменен")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func DeleteFirm(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	firmId := p.ByName("firmId")

	firm, err := firm.GetFirmById(firmId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = firm.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Фирма была успешно удалена")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
