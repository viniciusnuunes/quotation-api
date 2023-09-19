package database

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbName   = "postgres"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &sql.DB{}, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to create database connection pool %v", err)
	}

	return sqlDB, nil
}
