package services

import (
	"database/sql"
	"log"
)

type TestboardService struct {
	DB *sql.DB
}

func NewTestboardService(db *sql.DB) *TestboardService {
	return &TestboardService{DB: db}
}

func (s *TestboardService) GetAllTestboards() ([]map[string]interface{}, error) {
	rows, err := s.DB.Query("SELECT id, name, description FROM testboard")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testboards []map[string]interface{}
	for rows.Next() {
		var id int
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		testboards = append(testboards, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
		})
	}
	return testboards, nil
}
