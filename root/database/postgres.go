package database

import (
	"github.com/ethereal-go/ethereal"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabasePostgres struct {
	User         string
	Host         string
	Login        string
	Password     string
	DatabaseName string
	SslMode      string
}

func (m *DatabasePostgres) Connection() *gorm.DB {
	db, err := gorm.Open("postgres", "host="+
		m.Host+" user="+m.User+
		" dbname="+m.DatabaseName+"sslmode="+m.SslMode+" password="+m.Password)

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func (m *DatabasePostgres) Parse() DatabaseConnector {
	m.User = ethereal.GetCnf("DATABASE.USER").(string)
	m.Host = ethereal.GetCnf("DATABASE.HOST").(string)
	m.Password = ethereal.GetCnf("DATABASE.PASSWORD").(string)
	m.DatabaseName = ethereal.GetCnf("DATABASE.NAME").(string)
	m.SslMode = ethereal.GetCnf("DATABASE.SSLMODE").(string)
	return m
}
