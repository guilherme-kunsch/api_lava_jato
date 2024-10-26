package models

import "strings"

type Vehicle struct {
	ID       uint64 `json:"id,omitempty"`
	ClientId uint64 `json:"cliente_id,omitempty" validate:"required"`
	Brand    string `json:"marca,omitempty" validate:"required"`
	Model    string `json:"modelo,omitempty" validate:"required"`
	Plate    string `json:"placa,omitempty" validate:"required"`
	Year     int    `json:"ano,omitempty" validate:"required"`
}

type VehicleResponse struct {
	ID         uint64 `json:"id,omitempty"`
	NameClient string `json:"nome,omitempty" validate:"required"`
	Brand      string `json:"marca,omitempty" validate:"required"`
	Model      string `json:"modelo,omitempty" validate:"required"`
	Plate      string `json:"placa,omitempty" validate:"required"`
	Year       int    `json:"ano,omitempty" validate:"required"`
}

func (vehicle *Vehicle) Prepare() error {
	vehicle.format()
	return nil
}

func (vehicle *Vehicle) format() {
	vehicle.Brand = strings.TrimSpace(vehicle.Brand)
	vehicle.Model = strings.TrimSpace(vehicle.Model)
	vehicle.Plate = strings.TrimSpace(vehicle.Plate)
}
