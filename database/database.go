package database

import (
	"go-todo/config"
	"go-todo/pkg/database/mysql"
	"gorm.io/gorm"
)

func newDb() (db *gorm.DB) {
	var err error

	if config.Cfg().DBDriver == "mysql" {
		db, err = mysql.NewDb(mysql.Config{
			Host:     config.Cfg().DBHost,
			Port:     config.Cfg().DBPort,
			User:     config.Cfg().DBUser,
			Password: config.Cfg().DBPassword,
			DBName:   config.Cfg().DBName,
		})
	}

	if db == nil {
		panic("DB driver not supported")
	}

	if err != nil {
		panic(err)
	}

	return db
}

var db = newDb()

func DB() *gorm.DB {
	return db
}
