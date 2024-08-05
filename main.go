package main

import (
	"github.com/tr3tiak/avito_backend/internal/service"

	"github.com/tr3tiak/avito_backend/internal/controller"
	"github.com/tr3tiak/avito_backend/internal/repository"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	repo := repository.NewRepo()
	service := service.NewService(repo)
	controller := controller.NewController(service)

	r := mux.NewRouter()

	r.HandleFunc("/post", controller.HandlerPost)
	r.HandleFunc("/get", controller.HandlerGet)
	http.ListenAndServe("localhost:8080", r)
}
