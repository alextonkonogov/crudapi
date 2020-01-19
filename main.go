package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/alextonkonogov/crudapi/app/controller"
	"github.com/alextonkonogov/crudapi/app/model/user"
	"github.com/alextonkonogov/crudapi/app/server"
)

func main() {
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	server.InitCache()

	r := httprouter.New()
	routes(r)

	fmt.Println("service is working!")
	err = http.ListenAndServe(":5151", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.StartPage)
	r.POST("/login", controller.Login)
	r.GET("/logout", controller.Logout)

	r.GET("/cars", controller.GetCars)
	r.POST("/car/add", controller.AddCar)
	r.POST("/car/update/:carId", controller.UpdateCar)
	r.DELETE("/car/delete/:carId", controller.DeleteCar)

	r.GET("/firms", controller.GetFirms)
	r.POST("/firm/add", controller.AddFirm)
	r.POST("/firm/update/:firmId", controller.UpdateFirm)
	r.DELETE("/firm/delete/:firmId", controller.DeleteFirm)

	r.GET("/users", authorized(controller.GetUsers))
	r.POST("/user/add", authorized(admin(controller.AddUser, "users")))
	r.POST("/user/update/:userId", authorized(admin(controller.UpdateUser, "users")))
	r.DELETE("/user/delete/:userId", authorized(admin(controller.DeleteUser, "users")))
}

func authorized(next httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := server.ReadCookie("token", r)
		if err != nil {
			http.Error(rw, "Вы не авторизованы", 400)
			return
		}

		cache, err := server.GetCache()
		_, exists := cache.Get(token)

		if !exists {
			http.Error(rw, "Вы не авторизованы", 400)
			return
		}

		next(rw, r, ps)
	}
}

func admin(next httprouter.Handle, table string) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userId, _ := server.ReadCookie("Id", r)
		admin, _ := user.IsAdmin(userId, table)

		if !admin {
			http.Error(rw, "У Вас нет прав для выполнения данной операции", 400)
			return
		}

		next(rw, r, ps)
	}
}
