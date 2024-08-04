package service

import (
	"entity"
)

type Repo interface {
	Post(*entity.Adv) error
	Get(int) (*entity.Adv, error)
	GetPage(string) (*[]entity.Adv, error)
}

type Service interface {
	Post(*entity.Adv) error
	Get(int) (*entity.Adv, error)
	GetPage(string) (*[]entity.Adv, error)
}

type myService struct {
	repo Repo
}

func NewService(repo Repo) Service {
	myService := myService{repo: repo}
	return &myService
}

func (s *myService) Post(adv *entity.Adv) error {
	return s.repo.Post(adv)
}
func (s *myService) Get(id int) {
	return s.repo.Get(id)
}
func (s *myService) GetPage(orderBy string) (*[]entity.Adv, error) {
	return s.repo.GetPage(orderBy)
}
