package repository

import (
	"barCpGo/models"
	"github.com/jmoiron/sqlx"
)

type BarCommPostgres struct {
	db *sqlx.DB
}

func (b BarCommPostgres) GetBars() ([]models.Bar, error) {
	panic("implement me")
}

func NewBarCommPostgres(db *sqlx.DB) *BarCommPostgres {
	return &BarCommPostgres{db: db}
}
