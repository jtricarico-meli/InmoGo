package config

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("InmoGo.db"), &gorm.Config{})
}
