package ethereal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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

type EtherealModel struct{
	ID        uint `json:"id";gorm:"primary_key"`
	gorm.Model
}

type User struct {
	EtherealModel
	Email    string `json:"email";gorm:"type:unique_index"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Role     Role   `json:"role"`
	RoleID   int    `gorm:"index"`
}

type Role struct {
	EtherealModel
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"password"`
}
