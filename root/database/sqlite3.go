package database

import (
	"github.com/ethereal-go/ethereal"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Sqlite3 struct {
	Path string
}

func (m *Sqlite3) Connection() *gorm.DB {
	db, err := gorm.Open("sqlite3", m.Path)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func (m *Sqlite3) Parse() DatabaseConnector {
	m.Path = ethereal.GetCnf("DATABASE.PATH").(string)
}
