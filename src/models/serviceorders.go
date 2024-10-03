package models

import "time"

type ServiceOrder struct {
	ID          uint64    `json:"id,omitempty"`
	ClientId    uint64    `json:"cliente_id,omitempty" validate:"required"`
	VehicleId   uint64    `json:"veiculo_id,omitempty" validate:"required"`
	EmployeeId  uint64    `json:"funcionario_id,omitempty" validate:"required"`
	ServiceDate time.Time `json:"data_servico,omitempty" validate:"required"`
	Total       float64   `json:"total,omitempty" validate:"required"`
}
