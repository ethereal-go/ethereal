package database

import (
	"github.com/ethereal-go/ethereal"
	"github.com/jinzhu/gorm"
)

//type DatabaseConnector interface {
//	Connection() *gorm.DB
//}

type DatabaseMysql struct {
	Login        string
	Password     string
	DatabaseName string
}

func (m *DatabaseMysql) Connection() *gorm.DB {
	// configuration parameters
	var (
		login    string = ethereal.GetCnf("DATABASE.LOGIN").(string)
		password string = ethereal.GetCnf("DATABASE.PASSWORD").(string)
		database string = ethereal.GetCnf("DATABASE.NAME").(string)
	)

	db, err := gorm.Open("mysql", login+
		":"+password+
		"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
