package database

import (
	"github.com/ethereal-go/ethereal"
	"github.com/jinzhu/gorm"
)

type DatabaseMysql struct {
	Login        string
	Password     string
	DatabaseName string
}

func (m *DatabaseMysql) Connection() *gorm.DB {
	// configuration parameters
	db, err := gorm.Open("mysql", m.Login+
		":"+m.Password+
		"@/"+m.DatabaseName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func (m *DatabaseMysql) Parse() DatabaseConnector {
	// configuration parameters
	m.Login = ethereal.GetCnf("DATABASE.LOGIN").(string)
	m.Password = ethereal.GetCnf("DATABASE.PASSWORD").(string)
	m.DatabaseName = ethereal.GetCnf("DATABASE.NAME").(string)
	return m
}
