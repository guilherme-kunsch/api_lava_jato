package repositories

import (
	"database/sql"
	"lavajato/src/models"
)

type AreaOfActivity struct {
	db *sql.DB
}

func NewAreaOfActivity(db *sql.DB) *AreaOfActivity {
	return &AreaOfActivity{db}
}

func (repository AreaOfActivity) Create(area models.AreaOfActivity) (uint64, error) {
	statement, err := repository.db.Prepare("insert into cargos(cargo, salario) values (?, ?)")
	if err != nil {
		return 0, nil
	}

	insert, err := statement.Exec(area.Cargo, area.Salario)
	if err != nil {
		return 0, nil
	}

	lastID, err := insert.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}
