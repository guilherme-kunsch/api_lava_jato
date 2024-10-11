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
	statement, err := repository.db.Prepare("insert into funcionarios (nome, telefone, cargo_id) values(?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	insert, err := statement.Exec(&employee.Name, &employee.Phone, &employee.Area)
	if err != nil {
		return 0, nil
	}

	lastID, err := insert.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

func (repository Employee) SearchEmployeeId(ID uint64) ([]models.Employee, error) {
	rows, err := repository.db.Query("select * from funcionarios where id = ?", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee

		if err = rows.Scan(&employee.ID, &employee.Name, &employee.Phone, &employee.Area); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, err
}

func (repository Employee) SearchEmployee(ID uint64) ([]models.Employee, error) {

}
