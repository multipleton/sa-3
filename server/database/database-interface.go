package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SSLMode  string
}

func NewDatabaseConnection(config DatabaseConfiguration) (*sql.DB, error) {
	template := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	params := fmt.Sprintf(template, config.Host, config.Port, config.User, config.Password, config.Dbname, config.SSLMode)
	return sql.Open("postgres", params)
}
