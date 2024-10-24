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

	err := repository.db.QueryRow("SELECT EXISTS(SELECT 1 FROM clientes WHERE id = ?)", serviceOrders.ClientId).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("erro ao verificar cliente: %w", err)
	}
	if !exists {
		return 0, errors.New("cliente não encontrado")
	}

	err = repository.db.QueryRow("SELECT EXISTS(SELECT 1 FROM veiculos WHERE id = ?)", serviceOrders.VehicleId).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("erro ao verificar veículo: %w", err)
	}
	if !exists {
		return 0, errors.New("veículo não encontrado")
	}

	err = repository.db.QueryRow("SELECT EXISTS(SELECT 1 FROM funcionarios WHERE id = ?)", serviceOrders.EmployeeId).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("erro ao verificar funcionário: %w", err)
	}
	if !exists {
		return 0, errors.New("funcionário não encontrado")
	}

	statement, err := repository.db.Prepare("INSERT INTO ordens_de_servico (cliente_id, veiculo_id, funcionario_id, data_servico, total) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("erro ao preparar a query de inserção: %w", err)
	}
	defer statement.Close()

	insert, err := statement.Exec(serviceOrders.ClientId, serviceOrders.VehicleId, serviceOrders.EmployeeId, serviceOrders.ServiceDate, serviceOrders.Total)
	if err != nil {
		return 0, fmt.Errorf("erro ao executar a query de inserção: %w", err)
	}

	lastId, err := insert.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("erro ao pegar o último ID inserido: %w", err)
	}

	return uint64(lastId), nil
}

func (repository ServiceOrders) SearchServiceOrders(nameClient string) ([]models.ServiceOrdersResponse, error) {
	nameClient = fmt.Sprintf("%%%s%%", nameClient)

	rows, err := repository.db.Query("SELECT ordem.id as 'id', cli.nome as 'nome', serv.descricao as 'descricao', ordem.total as 'total', ordem.data_servico as 'data_do_servico' FROM ordens_de_servico as ordem INNER JOIN clientes as cli ON cli.id = ordem.cliente_id LEFT JOIN ordem_servico_servicos as oss ON oss.ordem_servico_id = ordem.id LEFT JOIN servicos as serv ON serv.id = oss.servico_id WHERE LOWER(cli.nome) LIKE ?", nameClient)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var servicesOrders []models.ServiceOrdersResponse

	for rows.Next() {
		var serviceOrder models.ServiceOrdersResponse

		if err := rows.Scan(&serviceOrder.ID, &serviceOrder.NameClient, &serviceOrder.Description, &serviceOrder.Amount, &serviceOrder.Date); err != nil {
			return nil, err
		}

		servicesOrders = append(servicesOrders, serviceOrder)
	}

	return servicesOrders, nil
}

func (repository ServiceOrders) SearchServiceOrdersID(ID uint64) ([]models.ServiceResponse, error) {
	rows, err := repository.db.Query("SELECT  ordem.id as 'id', cli.nome as 'nome_cliente', serv.descricao as 'descricao', vei.placa as 'placa', func.nome as 'nome_funcionario', ordem.total as 'total', ordem.data_servico as 'data_do_servico' FROM ordens_de_servico as ordem INNER JOIN clientes as cli ON cli.id = ordem.cliente_id LEFT JOIN ordem_servico_servicos as oss ON oss.ordem_servico_id = ordem.id LEFT JOIN servicos as serv ON serv.id = oss.servico_id LEFT JOIN veiculos as vei ON vei.id = ordem.veiculo_id LEFT JOIN funcionarios as func ON func.id = ordem.funcionario_id WHERE ordem.id = ?", ID)

	if err != nil {
		return nil, err
	}

	var services []models.ServiceResponse

	for rows.Next() {
		var service models.ServiceResponse
		if err := rows.Scan(&service.ID, &service.Name, &service.Description, &service.Plate, &service.NameEmployee, &service.Amount, &service.DateService); err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, err
}

func (repository ServiceOrders) UpdateServiceOrders(ID uint64, serviceOrders models.ServiceOrder) error {
	statement, err := repository.db.Prepare("UPDATE ordens_de_servico SET cliente_id = ?, veiculo_id = ?, funcionario_id = ?, data_servico = ?, total = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(&serviceOrders.ClientId, &serviceOrders.VehicleId, &serviceOrders.EmployeeId, &serviceOrders.ServiceDate, &serviceOrders.Total, ID); err != nil {
		return err
	}

	return nil
}

// func (repository ServiceOrders) DeleteServiceOrders(ID uint64) error {

// }
