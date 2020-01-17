package controller

import (
	"encoding/json"
	"github.com/alextonkonogov/crudapi/app/model/mark"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

func GetMarks(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	marks, err := mark.GetAllMarks()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode(marks)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func AddMark(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	title := r.FormValue("title")
	firmId := r.FormValue("firmId")

	if title == "" || firmId == "" {
		http.Error(rw, "Все поля должны быть заполнены", 400)
		return
	}

	mark, err := mark.NewMark(title, firmId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = mark.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Марка успешно добавлена!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateMark(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	markId := p.ByName("markId")
	title := strings.TrimSpace(p.ByName("title"))
	firmIdStr := strings.TrimSpace(p.ByName("firmId"))
	firmId, err := strconv.Atoi(firmIdStr)
	if err != nil {
		return
	}

	mark, err := mark.GetMarkById(markId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	mark.Title = title
	mark.FirmId = firmId

	err = mark.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Марка была успешно изменена")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func DeleteMark(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	markId := p.ByName("markId")

	mark, err := mark.GetMarkById(markId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = mark.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode("Марка была успешно удалена")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
