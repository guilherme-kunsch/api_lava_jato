package repositories

import (
	"database/sql"
	"lavajato/src/models"
)

type Client struct {
	db *sql.DB
}

func NewClient(db *sql.DB) *Client {
	return &Client{db}
}

func (repository Client) CreateClient(client models.Client) (uint64, error) {
	statement, err := repository.db.Prepare("insert into clientes (nome, telefone, email) values (?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	insert, err := statement.Exec(&client.Name, &client.Phone, &client.Email)
	if err != nil {
		return 0, nil
	}

	lastID, err := insert.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

func (repository Client) SearchClientID(ID uint64) ([]models.Client, error) {
	rows, err := repository.db.Query("select nome, telefone, email from clientes where id = ?", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clients []models.Client

	for rows.Next() {
		var c models.Client

		if err = rows.Scan(&c.Name, &c.Phone, &c.Email, &c.ID); err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	return clients, nil
}
