package models

import "strings"

type Service struct {
	ID          uint64  `json:"id,omitempty"`
	Descripiton string  `json:"descricao,omitempty" validate:"required"`
	Amount      float64 `json:"preco,omitempty" validate:"required"`
}

func (service *Service) Prepare() error {
	service.format()
	return nil
}

func (service *Service) format() {
	service.Descripiton = strings.TrimSpace(service.Descripiton)
}
