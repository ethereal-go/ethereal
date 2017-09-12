package ethereal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// TODO to install different variations of database
func Database() *gorm.DB {
	// configuration parameters
	var (
		login    string = GetCnf("DATABASE.LOGIN").(string)
		password string = GetCnf("DATABASE.PASSWORD").(string)
		database string = GetCnf("DATABASE.NAME").(string)
	)

	db, err := gorm.Open("mysql", login+
		":"+password+
		"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}