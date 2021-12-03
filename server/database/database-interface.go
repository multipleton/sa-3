package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DatabaseInterface struct {
	connection *sql.DB
}

func (d *DatabaseInterface) Init(host string, port string, user string, password string, dbname string, sslmode string) error {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode
	var err error
	d.connection, err = sql.Open("postgres", connStr)
	return err
}

func (d *DatabaseInterface) GetConnection() *sql.DB {
	return d.connection
}

func (d *DatabaseInterface) CloseConnection() {
	if d.connection != nil {
		d.connection.Close()
	}
}
