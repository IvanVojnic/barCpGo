package service

import (
	"barCpGo/models"
	"barCpGo/pkg/repository"
)

type BarCommService struct {
	repo repository.BarComm
}

func (b BarCommService) GetBars() ([]models.Bar, error) {
	bars, err := b.repo.GetBars()
	if err != nil {
		return bars, err
	}
	return bars, nil
}

func (b BarCommService) GetBarName(id_place int) (string, error) {
	barName, err := b.repo.GetBarName(id_place)
	if err != nil {
		return barName, err
	}
	return barName, nil
}

func NewBarComm(repo repository.BarComm) *BarCommService {
	return &BarCommService{repo: repo}
}
