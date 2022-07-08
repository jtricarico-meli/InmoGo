package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	return gorm.Open(mysql.Open("root:root@/InmoTricarico?parseTime=true"), &gorm.Config{})
}
