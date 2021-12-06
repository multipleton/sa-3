package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDatabaseConnection(host string, port string, user string, password string, dbname string, sslmode string) (*sql.DB, error) {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode
	return sql.Open("postgres", connStr)
}
