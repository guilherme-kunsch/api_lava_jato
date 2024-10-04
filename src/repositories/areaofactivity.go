package repositories

import (
	"database/sql"
	"fmt"
	"lavajato/src/models"
)

type AreaOfActivity struct {
	db *sql.DB
}

func NewAreaOfActivity(db *sql.DB) *AreaOfActivity {
	return &AreaOfActivity{db}
}

func (repository AreaOfActivity) Create(area models.AreaOfActivity) (uint64, error) {
	statement, err := repository.db.Prepare("insert into cargos(cargo, salario) values (?, ?)")
	if err != nil {
		return 0, nil
	}

	insert, err := statement.Exec(area.Cargo, area.Salario)
	if err != nil {
		return 0, nil
	}

	lastID, err := insert.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

// Search Area
func (repository AreaOfActivity) Search(area string) ([]models.AreaOfActivity, error) {
	area = fmt.Sprintf("%%%s%%", area)
	rows, err := repository.db.Query("select cargo, salario from cargos where cargo LIKE ?", area)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var areas []models.AreaOfActivity

	for rows.Next() {
		var area models.AreaOfActivity
		if err := rows.Scan(&area.Cargo, &area.Salario); err != nil {
			return nil, err
		}

		areas = append(areas, area)
	}

	return areas, nil
}

//Search ID

func (respository AreaOfActivity) SearchId(ID uint64) (models.AreaOfActivity, error) {
	rows, err := respository.db.Query("select * from cargos where id = ?", ID)
	if err != nil {
		return models.AreaOfActivity{}, nil
	}

	defer rows.Close()

	var area models.AreaOfActivity

	if rows.Next() {
		if err = rows.Scan(&area.ID, &area.Cargo, &area.Salario); err != nil {
			return models.AreaOfActivity{}, err
		}
	}

	return area, nil
}
