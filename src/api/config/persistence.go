package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	//config
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "/InmoTricarico"

	//tables
	Inmueble    = "inmueble"
	Inquilino   = "inquilino"
	Pagos       = "pagos"
	Alquiler    = "alquiler"
	Propietario = "propietario"
)

var (
	DB sql.DB
)

func InitDB() {
	DB, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		panic(err)
	}
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
