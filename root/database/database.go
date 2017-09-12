package database

import "github.com/jinzhu/gorm"

type DatabaseConnector interface {
	Connection() *gorm.DB
}
