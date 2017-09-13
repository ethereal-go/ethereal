package ethereal

import (
	"github.com/ethereal-go/ethereal/root/database"
	"github.com/jinzhu/gorm"
	"github.com/ethereal-go/ethereal/root/config"
)

func Database() *gorm.DB {
	db, err := database.FactoryDatabase(config.GetCnf("DATABASE.TYPE").(string))
	if err != nil {
		panic(err)
	}
	return db.Parse().Connection()
}
