package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"

	"github.com/alextonkonogov/crudapi/app/model/car"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cars, err := car.GetAllCars()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	main := filepath.Join("public", "html", "cars.html")
	common := filepath.Join("public", "html", "common.html")

	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(rw, "cars", cars)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
