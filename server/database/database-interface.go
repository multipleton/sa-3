package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConfiguration struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func NewDatabaseConnection(config DatabaseConfiguration) (*sql.DB, error) {
	template := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	params := fmt.Sprintf(template, config.host, config.port, config.user, config.password, config.dbname, config.sslmode)
	return sql.Open("postgres", params)
}
