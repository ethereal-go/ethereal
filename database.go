package ethereal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_LOGIN")+
		":"+os.Getenv("MYSQL_PASSWORD")+
		"@/"+os.Getenv("MYSQL_DATABASE")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func ConstructorDb() *gorm.DB {
	if app.Db == nil {
		envLoading()
		app.Db = Database()
	}
	return app.Db

}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token   string  `json:token`
	Role     Role   `json:"role"`
	RoleID   int    `gorm:"index"`
}

type Role struct {
	gorm.Model
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"password"`
}
