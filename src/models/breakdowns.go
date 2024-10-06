package models

import (
	"strings"

	"github.com/shopspring/decimal"
)

type Breakdowns struct {
	ID              uint64          `json:"id,omitempty"`
	OrdersService   uint64          `json:"ordem_servico_id,omitempty" validate:"required"`
	Description     string          `json:"descricao,omitempty" validate:"required"`
	EmployeeID      uint64          `json:"funcionario_id,omitempty"`
	DiscountApplied bool            `json:"desconto_aplicado,omitempty"`
	ValueDiscount   decimal.Decimal `json:"valor_descontado,omitempty"`
}

func (breakdown *Breakdowns) Prepare() error {
	breakdown.format()
	return nil
}

func (breakdown *Breakdowns) format() {
	breakdown.Description = strings.TrimSpace(breakdown.Description)
}
