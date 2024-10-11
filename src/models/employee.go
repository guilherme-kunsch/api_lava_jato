package models

import "strings"

type Employee struct {
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"nome,omitempty" validate:"required"`
	Phone string `json:"telefone,omitempty" validate:"required"`
	Area  uint64 `json:"cargo_id,omitempty" validate:"required"`
}

func (employee *Employee) Prepare() error {
	employee.format()
	return nil
}

func (employee *Employee) format() {
	employee.Name = strings.TrimSpace(employee.Name)
	employee.Phone = strings.TrimSpace(employee.Phone)
}
