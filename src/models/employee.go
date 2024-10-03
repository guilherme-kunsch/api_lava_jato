package models

type Employee struct {
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"nome,omitempty" validate:"required"`
	Phone string `json:"telefone,omitempty" validate:"required"`
	Area  string `json:"cargo,omitempty" validate:"required"`
}
