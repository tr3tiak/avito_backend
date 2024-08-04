package main

import (
	"internal/controller"
	"internal/repository"
	"internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	repo := repository.NewRepo()
	service := service.NewService(repo)
	controller := controller.NewController(service)

	r := mux.NewRouter()

	r.HandleFunc("/post", controller.HandlerPost)
	http.ListenAndServe("localhost:8080", r)
}
