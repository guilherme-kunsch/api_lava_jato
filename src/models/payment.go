package models

import "time"

type Payment struct {
	ID          uint64    `json:"id,omitempty"`
	EmployeeId  uint64    `json:"id_funcionario,omitempty" validate:"required"`
	DatePayment time.Time `json:"data_pagamento,omitempty" validate:"required"`
	Value       float64   `json:"valor,omitempty" validate:"required"`
}
