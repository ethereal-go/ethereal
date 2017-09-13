package ethereal

import (
	"github.com/ethereal-go/ethereal/root/database"
	"github.com/jinzhu/gorm"
)

// TODO to install different variations of database

func Database() *gorm.DB {
	db, err := database.FactoryDatabase(GetCnf("DATABASE.TYPE").(string))
	if err != nil {
		panic(err)
	}
	return db.Parse().Connection()
}
