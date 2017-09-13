package database

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type DatabaseConnector interface {
	Connection() *gorm.DB
	Parse() DatabaseConnector
}

func FactoryDatabase(typeDB string) (DatabaseConnector, error) {
	switch typeDB {
	case "mysql":
		return DatabaseMysql{}, nil
	case "postgres":
		return DatabasePostgres{}, nil
	case "sqlite3":
		return DatabasePostgres{}, nil
	case "sqlServer":
		return SQLServer{}, nil
	default:
		return nil, errors.New("You have not selected a database type.")
	}
}
