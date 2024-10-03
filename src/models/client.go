package models

type Client struct {
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"nome,omitempty" validate:"require"`
	Phone string `json:"telefone,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
}
