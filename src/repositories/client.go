package repositories

import (
	"database/sql"
	"fmt"
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

		if err = rows.Scan(&c.Name, &c.Phone, &c.Email); err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	return clients, nil
}

func (repository Client) SearchClient(client string) ([]models.Client, error) {
	client = fmt.Sprintf("%%%s%%", client)

	rows, err := repository.db.Query("select * from clientes where nome LIKE ?", client)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clients []models.Client

	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Phone, &client.Email); err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (repository Client) UpdateClient(ID uint64, client models.Client) error {
	statement, err := repository.db.Prepare("update clientes set nome = ?, telefone = ?, email = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(&client.Name, &client.Phone, &client.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository Client) DeleteClient(ID uint64) error {
	var count int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM ordens_de_servico WHERE cliente_id = ?", ID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("não é possível excluir o cliente, pois ele possui %d ordens de serviço associadas", count)
	}
	statement, err := repository.db.Prepare("delete from clientes where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
