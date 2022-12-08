package repository

import (
	"barCpGo/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BarCommPostgres struct {
	db *sqlx.DB
}

func (b BarCommPostgres) GetBars() ([]models.Bar, error) {
	var bars []models.Bar
	query := fmt.Sprintf("SELECT * FROM BARS")
	err := b.db.Select(&bars, query)
	if err != nil {
		fmt.Println(err)
	}
	return bars, nil
}

func (b BarCommPostgres) GetBarName(id_place int) (string, error) {
	var bar models.Bar
	query := fmt.Sprintf("SELECT * FROM BARS WHERE BARS.id = $1")
	err := b.db.Select(&bar, query, id_place)
	if err != nil {
		fmt.Println(err)
	}
	return bar.Name, nil
}

func NewBarCommPostgres(db *sqlx.DB) *BarCommPostgres {
	return &BarCommPostgres{db: db}
}
