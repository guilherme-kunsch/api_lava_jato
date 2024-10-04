package repositories

import (
	"database/sql"
	"lavajato/src/models"
)

type Breakdowns struct {
	db *sql.DB
}

func NewBreakdowns(db *sql.DB) *Breakdowns {
	return &Breakdowns{db}
}

func (repository Breakdowns) Create(breakdowns models.Breakdowns) (uint64, error) {
	statement, err := repository.db.Prepare("insert into avarias (ordem_servico_id, descricao, funcionario_id, desconto_aplicado) values (?, ?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	insert, err := statement.Exec(&breakdowns.OrdersService, &breakdowns.Description, &breakdowns.EmployeeID, &breakdowns.DiscountApplied)
	if err != nil {
		return 0, nil
	}

	lastID, err := insert.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

func (repository Breakdowns) SearchBreakdown(ID uint64) ([]models.Breakdowns, error) {
	rows, err := repository.db.Query("SELECT ordem_servico_id, descricao, funcionario_id, desconto_aplicado FROM avarias WHERE ordem_servico_id = ?", ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var breakdowns []models.Breakdowns

	for rows.Next() {
		var b models.Breakdowns
		if err := rows.Scan(&b.OrdersService, &b.Description, &b.EmployeeID, &b.DiscountApplied); err != nil {
			return nil, err
		}

		breakdowns = append(breakdowns, b)
	}

	return breakdowns, nil
}
