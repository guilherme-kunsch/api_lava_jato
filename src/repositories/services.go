package repositories

import (
	"database/sql"
	"fmt"
	"lavajato/src/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db}
}

func (repository Service) CreateService(service models.Service) (uint64, error) {
	stament, err := repository.db.Prepare("INSERT INTO servicos(descricao, preco) VALUES (?, ?)")
	if err != nil {
		return 0, nil
	}

	defer stament.Close()

	insert, err := stament.Exec(service.Descripiton, service.Amount)
	if err != nil {
		return 0, nil
	}

	lastId, err := insert.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastId), nil
}

func (repository Service) SearchServiceID(ID uint64) ([]models.Service, error) {
	rows, err := repository.db.Query("SELECT * FROM servicos WHERE id = ?", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var services []models.Service

	for rows.Next() {
		var service models.Service

		if err := rows.Scan(&service.ID, &service.Descripiton, &service.Amount); err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

func (repository Service) UpdateService(ID uint64, service models.Service) error {
	statement, err := repository.db.Prepare("UPDATE servicos SET descricao = ?, preco = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(&service.Descripiton, &service.Amount, ID); err != nil {
		return err
	}

	return nil
}

func (repository Service) DeleteService(ID uint64) error {
	var count int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM servicos WHERE id = ?", ID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("não é possível excluir um serviço, pois ele possui %d ordem de serviço associadas", count)
	}

	statement, err := repository.db.Prepare("DELETE FROM servico WHERE id = ?")
	if err != nil {
		return nil
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil

}
