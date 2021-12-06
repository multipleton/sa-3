package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDatabaseConnection(host string, port string, user string, password string, dbname string, sslmode string) (*sql.DB, error) {
	template := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	params := fmt.Sprintf(template, host, port, user, password, dbname, sslmode)
	return sql.Open("postgres", params)
}
