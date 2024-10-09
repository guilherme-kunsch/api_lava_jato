package models

import "strings"

type Client struct {
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"nome,omitempty" validate:"required"`
	Phone string `json:"telefone,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
}

func (client *Client) Prepare() error {
	client.format()
	return nil
}

func (client *Client) format() {
	client.Name = strings.TrimSpace(client.Name)
	client.Phone = strings.TrimSpace(client.Phone)
	client.Email = strings.TrimSpace(client.Email)
}
