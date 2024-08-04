package controller

import (
	"encoding/json"
	"net/http"

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
	err := c.s.Post(&adv)
	if err != nil {
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/plain")

	// Устанавливаем код состояния 200 OK
	w.WriteHeader(http.StatusOK)
}
