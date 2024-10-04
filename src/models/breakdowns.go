package models

import (
	"strings"

	"github.com/shopspring/decimal"
)

type Breakdowns struct {
	ID              uint64          `json:"id,omitempty"`
	OrdersService   uint64          `json:"ordem_servico_id,omitempty" validate:"required"`
	Description     string          `json:"descricao,omitempty" validate:"required"`
	EmployeeID      *int            `json:"funcionario_id,omitempty"`
	DiscountApplied decimal.Decimal `json:"desconto_aplicado,omitempty"`
}

func (breakdown *Breakdowns) Prepare() error {
	breakdown.format()
	return nil
}

func (breakdown *Breakdowns) format() {
	breakdown.Description = strings.TrimSpace(breakdown.Description)
}
