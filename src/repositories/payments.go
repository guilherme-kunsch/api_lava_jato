package repositories

import (
	"database/sql"
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
