package repositories

import (
	"database/sql"
	"fmt"
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

	insert, err := statement.Exec(employee.Name, employee.Phone, employee.Area)
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

func (repository Employee) SearchEmployee(employee string) ([]models.Employee, error) {
	employee = fmt.Sprintf("%%%s%%", employee)

	results, err := repository.db.Query("select * from funcionarios where nome LIKE ?", employee)
	if err != nil {
		return nil, err
	}

	var employees []models.Employee

	for results.Next() {
		var employee models.Employee
		if err := results.Scan(&employee.ID, &employee.Name, &employee.Phone, &employee.Area); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (repository Employee) UpdateEmployee(ID uint64, employee models.Employee) error {
	statement, err := repository.db.Prepare("update funcionarios set nome = ?, telefone = ?, cargo_id = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(&employee.Name, &employee.Phone, &employee.Area, ID); err != nil {
		return err
	}

	return nil
}

func (repository Employee) DeleteEmployee(ID uint64) error {
	var count int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM ordens_de_servico WHERE funcionario_id = ?", ID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("não é possível excluir o funcionário, pois ele possui %d ordens de serviço associadas", count)
	}
	statement, err := repository.db.Prepare("delete from funcionarios where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
