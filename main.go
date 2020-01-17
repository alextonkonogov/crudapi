package main

import (
	"fmt"
	"github.com/alextonkonogov/crudapi/app/controller"
	"github.com/alextonkonogov/crudapi/app/server"
	"github.com/julienschmidt/httprouter"

	"log"
	"net/http"
)

func main() {
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := httprouter.New()
	routes(r)

	fmt.Println("service is working!")

	err = http.ListenAndServe("localhost:5151", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.StartPage)
	r.GET("/cars", controller.GetCars)
	r.POST("/car/add", controller.AddCar)
	r.POST("/car/update/:carId", controller.UpdateCar)
	r.DELETE("/car/delete/:carId", controller.DeleteCar)

	r.GET("/firms", controller.GetFirms)
	r.POST("/firm/add", controller.AddFirm)
	r.POST("/firm/update/:firmId", controller.UpdateFirm)
	r.DELETE("/firm/delete/:firmId", controller.DeleteFirm)

	r.GET("/marks", controller.GetMarks)
	r.POST("/mark/add", controller.AddMark)
	r.POST("/mark/update/:markId", controller.UpdateMark)
	r.DELETE("/mark/delete/:markId", controller.DeleteMark)
}
