package models

type Vehicle struct {
	ID       uint64 `json:"id,omitempty"`
	ClientId uint64 `json:"cliente_id,omitempty" validate:"required"`
	Brand    string `json:"marca,omitempty" validate:"required"`
	Model    string `json:"modelo,omitempty" validate:"required"`
	Plate    string `json:"placa,omitempty" validate:"required"`
	Year     int    `json:"ano,omitempty" validate:"required"`
}
