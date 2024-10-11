package repositories

import (
	"database/sql"
	"lavajato/src/models"
)

type Employee struct {
	db *sql.DB
}

func NewEmployee(db *sql.DB) *Employee {
	return &Employee{db}
}

func (repository Employee) CreateEmployee(employee models.Employee) (uint64, error) {
	body, err := repository.db.Prepare("insert into funcionarios")
	if err != nil {
		return 0, nil
	}
}
