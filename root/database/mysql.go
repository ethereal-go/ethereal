package database

import (
	"github.com/ethereal-go/ethereal/root/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		panic("failed to connect database : " + err.Error())
	}
	return db
}

func (m *DatabaseMysql) Parse() DatabaseConnector {
	// configuration parameters
	m.Login = config.GetCnf("DATABASE.MYSQL.LOGIN").(string)
	m.Password = config.GetCnf("DATABASE.MYSQL.PASSWORD").(string)
	m.DatabaseName = config.GetCnf("DATABASE.MYSQL.NAME").(string)
	return m
}
