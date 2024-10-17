package repositories

import (
	"database/sql"
	"fmt"
	"lavajato/src/models"
)

type Payment struct {
	db *sql.DB
}

func NewPayment(db *sql.DB) *Payment {
	return &Payment{db}
}

func (repository Payment) CreatePayment(payment models.Payment) (uint64, error) {
	statement, err := repository.db.Prepare("insert into pagamentos (funcionario_id, data_pagamento, valor) values (?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	insertId, err := statement.Exec(&payment.EmployeeId, &payment.DatePayment, &payment.Value)
	if err != nil {
		return 0, nil
	}

	lastId, err := insertId.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastId), err
}

func (repository Payment) SearchPaymentID(ID uint64) ([]models.Payment, error) {
	rows, err := repository.db.Query("select * from pagamentos where id = ?", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		if err = rows.Scan(&payment.ID, &payment.EmployeeId, &payment.DatePayment, &payment.Value); err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, err
}

func (repository Payment) SearchPayment(employee string) ([]models.PaymentResponse, error) {
	employee = fmt.Sprintf("%%%s%%", employee)

	rows, err := repository.db.Query("SELECT p.id, f.nome, p.data_pagamento, c.cargo, p.valor FROM pagamentos AS p LEFT JOIN funcionarios AS f ON p.funcionario_id = f.id LEFT JOIN cargos AS c ON f.cargo_id = c.id WHERE LOWER(f.nome) LIKE ?", employee)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.PaymentResponse

	for rows.Next() {
		var payment models.PaymentResponse
		if err := rows.Scan(&payment.ID, &payment.EmployeeName, &payment.PaymentDate, &payment.Role, &payment.Amount); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (repository Payment) UpdatePayment(ID uint64, payment models.Payment) error {
	statement, err := repository.db.Prepare("update pagamentos set funcionario_id = ?, data_pagamento = ?, valor = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(payment.EmployeeId, payment.DatePayment, payment.Value, ID); err != nil {
		return err
	}

	return nil
}

func (repository Payment) DeletePayment(ID uint64) error {
	statement, err := repository.db.Prepare("delete from pagamentos where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return err
}
