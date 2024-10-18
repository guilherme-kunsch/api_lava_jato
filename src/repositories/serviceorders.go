package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"lavajato/src/models"
)

type ServiceOrders struct {
	db *sql.DB
}

func NewServiceOrders(db *sql.DB) *ServiceOrders {
	return &ServiceOrders{db}
}

func (repository ServiceOrders) CreateServiceOrders(serviceOrders models.ServiceOrder) (uint64, error) {
	var exists bool

	// Verificando se o cliente existe
	err := repository.db.QueryRow("SELECT EXISTS(SELECT 1 FROM clientes WHERE id = ?)", serviceOrders.ClientId).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("erro ao verificar cliente: %w", err)
	}
	if !exists {
		return 0, errors.New("cliente não encontrado")
	}

	// Verificando se o veículo existe
	err = repository.db.QueryRow("SELECT EXISTS(SELECT 1 FROM veiculos WHERE id = ?)", serviceOrders.VehicleId).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("erro ao verificar veículo: %w", err)
	}
	if !exists {
		return 0, errors.New("veículo não encontrado")
	}

	// Verificando se o funcionário existe
	err = repository.db.QueryRow("SELECT EXISTS(SELECT 1 FROM funcionarios WHERE id = ?)", serviceOrders.EmployeeId).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("erro ao verificar funcionário: %w", err)
	}
	if !exists {
		return 0, errors.New("funcionário não encontrado")
	}

	// Preparando a inserção
	statement, err := repository.db.Prepare("INSERT INTO ordens_de_servico (cliente_id, veiculo_id, funcionario_id, data_servico, total) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("erro ao preparar a query de inserção: %w", err)
	}
	defer statement.Close()

	// Executando a inserção
	insert, err := statement.Exec(serviceOrders.ClientId, serviceOrders.VehicleId, serviceOrders.EmployeeId, serviceOrders.ServiceDate, serviceOrders.Total)
	if err != nil {
		return 0, fmt.Errorf("erro ao executar a query de inserção: %w", err)
	}

	// Pegando o último ID inserido
	lastId, err := insert.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("erro ao pegar o último ID inserido: %w", err)
	}

	return uint64(lastId), nil
}

// func (repository ServiceOrders) SearchServiceOrders(serviceOrders string) ([]models.ServiceOrder, error) {

// }

// func (repository ServiceOrders) SearchServiceOrdersID(ID uint64) ([]models.ServiceOrder, error) {

// }

// func (repository ServiceOrders) UpdateServiceOrders(ID uint64, serviceOrders models.ServiceOrder) error {

// }

// func (repository ServiceOrders) DeleteServiceOrders(ID uint64) error {

// }
