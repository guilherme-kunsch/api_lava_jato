package models

type Breakdowns struct {
	ID              uint64  `json:"id,omitempty"`
	OrdersService   uint64  `json:"numero_ordem,omitempty" validate:"required"`
	Description     string  `json:"descricao,omitempty" validate:"required"`
	FuncionarioID   *int    `json:"funcionario_id,omitempty"`
	DiscountApplied float64 `json:"desconto_aplicado,omitempty"`
}
