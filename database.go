package ethereal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	// configuration parameters
	var (
		login    string = config("DATABASE.LOGIN").(string)
		password string = config("DATABASE.PASSWORD").(string)
		database string = config("DATABASE.NAME").(string)
	)

	db, err := gorm.Open("mysql", login+
		":"+password+
		"@/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type User struct {
	gorm.Model
	Email    string `json:"email";gorm:"type:unique_index"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Role     Role   `json:"role"`
	RoleID   int    `gorm:"index"`
}

type Role struct {
	gorm.Model
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"password"`
}
