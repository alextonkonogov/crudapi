package main

import (
	"github.com/alextonkonogov/crudapi/app/controller"
	"github.com/alextonkonogov/crudapi/app/server"
	"github.com/julienschmidt/httprouter"

	"log"
	"net/http"
)

func main() {
	//инициализируем подключение к базе данных
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	//создаем и запускаем в работу роутер для обслуживания запросов
	r := httprouter.New()
	routes(r)

	//прикрепляемся хосту и порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	err = http.ListenAndServe("localhost:5151", r)
	if err != nil {
		log.Fatal(err)
	}

}

func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.StartPage)
	r.GET("/cars", controller.GetCars)
	r.POST("/car/add", controller.AddCar)
	r.POST("/car/update/:carId", controller.UpdateCar)
	r.DELETE("/car/delete/:carId", controller.DeleteCar)
}
