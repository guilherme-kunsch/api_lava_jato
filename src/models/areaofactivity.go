package models

type AreaOfActivity struct {
	ID      uint64  `json:"id,omitempty"`
	Cargo   string  `json:"cargo,omitempty" validate:"required"`
	Salario float64 `json:"salario,omitempty" validate:"required"`
}
