package repository

import (
	"barCpGo/models"
	"github.com/jmoiron/sqlx"
)

type BarComm interface {
	GetBars() ([]models.Bar, error)
}

type Repository struct {
	BarComm
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		BarComm: NewBarCommPostgres(db),
	}
}
