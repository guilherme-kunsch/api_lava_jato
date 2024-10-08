package repositories

import (
	"database/sql"
	"fmt"
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

// Search breakdown ID
func (repository Breakdowns) SearchBreakdownId(ID uint64) ([]models.Breakdowns, error) {
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

// Search breakdowns
func (repository Breakdowns) SearchBreakdown(breakdown string) ([]models.Breakdowns, error) {
	breakdown = fmt.Sprintf("%%%s%%", breakdown)
	rows, err := repository.db.Query("select id, ordem_servico_id, descricao, funcionario_id, desconto_aplicado from avarias where descricao LIKE ?", breakdown)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var breakdowns []models.Breakdowns

	for rows.Next() {
		var breakdown models.Breakdowns
		if err := rows.Scan(&breakdown.ID, &breakdown.OrdersService, &breakdown.Description, &breakdown.EmployeeID, &breakdown.DiscountApplied); err != nil {
			return nil, err
		}

		breakdowns = append(breakdowns, breakdown)
	}

	return breakdowns, nil
}

func (repository Breakdowns) UpdateBreakdown(ID uint64, breakdown models.Breakdowns) error {
	statement, err := repository.db.Prepare("update avarias set ordem_servico_id = ?, descricao = ?, funcionario_id = ?, desconto_aplicado = ?, valor_descontado = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(breakdown.OrdersService, breakdown.Description, breakdown.EmployeeID, breakdown.DiscountApplied, breakdown.ValueDiscount, ID); err != nil {
		return err
	}

	return nil

}

func (repository Breakdowns) DeleteBreakdowns(ID uint64) error {
	statement, err := repository.db.Prepare("delete from avarias where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil

}
