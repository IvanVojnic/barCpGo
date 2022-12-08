package service

import (
	"barCpGo/models"
	"barCpGo/pkg/repository"
)

type BarComm interface {
	GetBars() ([]models.Bar, error)
	GetBarName(id_place int) (string, error)
}

type Service struct {
	BarComm
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		BarComm: NewBarComm(repos.BarComm),
	}
}
