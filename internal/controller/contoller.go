package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tr3tiak/avito_backend/internal/entity"
	"github.com/tr3tiak/avito_backend/internal/service"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{s: s}
}

func (c *Controller) HandlerPost(w http.ResponseWriter, r *http.Request) {
	adv := entity.Adv{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&adv)
	fmt.Println("service started", adv.Name, adv.Description)
	err := c.s.Post(&adv)
	if err != nil {
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/plain")

	// Устанавливаем код состояния 200 OK
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) HandlerGet(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(data["id"].(string))
	if err != nil {
		fmt.Println("Ошибка при преобразовании строки в число:", err)
		return
	}
	adv, err := c.s.Get(id)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(adv)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(response)
	w.WriteHeader(http.StatusOK)

}

func (c *Controller) HandlerGetPage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	orderBy := data["orderBy"].(string)
	adv, err := c.s.GetPage(orderBy)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(adv)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(response)
	w.WriteHeader(http.StatusOK)

}
