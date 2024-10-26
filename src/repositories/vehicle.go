package repositories

import (
	"database/sql"
	"fmt"
	"lavajato/src/models"
)

type Vehicle struct {
	db *sql.DB
}

func NewVehicle(db *sql.DB) *Vehicle {
	return &Vehicle{db}
}

func (repository Vehicle) CreateVehicle(vehicle models.Vehicle) (uint64, error) {
	var name string
	err := repository.db.QueryRow("SELECT cli.nome FROM veiculos AS vei INNER JOIN clientes AS cli ON cli.id = vei.cliente_id WHERE cli.id = ?", vehicle.ClientId).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("cliente não encontrado")
		}
		return 0, fmt.Errorf("erro ao buscar cliente: %v", err)
	}

	var count int
	err = repository.db.QueryRow("SELECT COUNT(*) FROM veiculos WHERE placa = ?", vehicle.Plate).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("erro ao contar veículos: %v", err)
	}

	if count > 0 {
		return 0, fmt.Errorf("a placa já existe no sistema e pertence a %s", name)
	}

	stament, err := repository.db.Prepare("INSERT INTO veiculos(cliente_id, marca, modelo, placa, ano) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("erro ao preparar a instrução: %v", err)
	}
	defer stament.Close()

	createVehicle, err := stament.Exec(vehicle.ClientId, vehicle.Brand, vehicle.Model, vehicle.Plate, vehicle.Year)
	if err != nil {
		return 0, fmt.Errorf("erro ao inserir veículo: %v", err)
	}

	lastId, err := createVehicle.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("erro ao obter o ID do veículo criado: %v", err)
	}

	return uint64(lastId), nil
}

func (repository Vehicle) SearchVehicle(plate string) ([]models.VehicleResponse, error) {
	if plate == "" {
		return nil, fmt.Errorf("a placa não pode estar vazia")
	}

	plateFormatted := fmt.Sprintf("%%%s%%", plate)

	rows, err := repository.db.Query("SELECT vei.id, cli.nome, vei.marca, vei.modelo, vei.placa, vei.ano FROM veiculos AS vei INNER JOIN clientes AS cli ON cli.id = vei.cliente_id WHERE placa LIKE ?", plateFormatted)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var vehicles []models.VehicleResponse

	for rows.Next() {
		var vehicle models.VehicleResponse

		if err := rows.Scan(&vehicle.ID, &vehicle.NameClient, &vehicle.Brand, &vehicle.Model, &vehicle.Plate, &vehicle.Year); err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	if len(vehicles) == 0 {
		return nil, fmt.Errorf("não foi possível encontrar a placa desejada: %v", plate)
	}

	return vehicles, nil
}

func (repository Vehicle) SearchVehicleID(ID uint64) ([]models.VehicleResponse, error) {

}
