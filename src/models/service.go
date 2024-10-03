package models

type OrderServiceService struct {
	ID             uint64 `json:"id,omitempty"`
	OrderServiceId uint64 `json:"ordem_servico_id,omitempty" validate:"required"`
	ServiceId      uint64 `json:"servico_id,omitempty" validate:"required"`
}
