package repositories

// 데이터베이스 작업을 수행

import (
	"backend/internal/models"
	"database/sql"
)

type TestboardRepository struct {
	DB *sql.DB
}

func NewTestboardRepository(db *sql.DB) *TestboardRepository {
	return &TestboardRepository{DB: db}
}

func (r *TestboardRepository) GetAll() ([]models.Testboard, error) {
	rows, err := r.DB.Query("SELECT id, name, description FROM testboard")
	if err != nil {
		return nil, err
	}
	defer rows.Close();

	var testboards []models.Testboard
	for rows.Next() {
		var t models.Testboard
		if err := rows.Scan(&t.ID, &t.Name, &t.Description); err != nil {
			return nil, err
		}
		testboards = append(testboards, t)
	}

	return testboards, nil
}
