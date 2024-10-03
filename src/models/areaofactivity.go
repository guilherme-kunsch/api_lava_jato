package models

import (
	"strings"

	"github.com/shopspring/decimal"
)

type AreaOfActivity struct {
	ID      uint64          `json:"id,omitempty"`
	Cargo   string          `json:"cargo,omitempty" validate:"required"`
	Salario decimal.Decimal `json:"salario,omitempty" validate:"required"`
}

func (area *AreaOfActivity) Prepare() error {
	area.format()
	return nil
}

func (area *AreaOfActivity) format() {
	area.Cargo = strings.TrimSpace(area.Cargo)
}
