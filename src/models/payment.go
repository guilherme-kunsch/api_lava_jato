package models

import (
	"time"
)

type Payment struct {
	ID          uint64    `json:"id,omitempty"`
	EmployeeId  uint64    `json:"funcionario_id,omitempty" validate:"required"`
	DatePayment time.Time `json:"data_pagamento,omitempty" validate:"required" time_format:"2006-01-02"`
	Value       float64   `json:"valor,omitempty" validate:"required"`
}

type PaymentResponse struct {
	ID           uint64    `json:"id"`
	EmployeeName string    `json:"nome_funcionario"`
	PaymentDate  time.Time `json:"data_pagamento"`
	Role         string    `json:"cargo"`
	Amount       float64   `json:"valor"`
}
