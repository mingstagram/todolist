package db

import (
	"backend/config"
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// 연결 테스트
	// if err := db.Ping(); err != nil {
	// 	log.Fatalf("Could not ping the database: %v", err)
	// }

	log.Println("Database connected successfully!")
	return db
}