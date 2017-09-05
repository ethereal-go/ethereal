package ethereal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"github.com/spf13/viper"
	"strings"
)

func Database() *gorm.DB {
	var (
		login string
		password string
		database string

	)
	if login = os.Getenv("DATABASE.LOGIN"); login == "" {
		login = viper.GetString(strings.ToLower("DATABASE.LOGIN"))
	}
	if password = os.Getenv("DATABASE.PASSWORD"); password == "" {
		password = viper.GetString(strings.ToLower("DATABASE.PASSWORD"))
	}
	if database = os.Getenv("DATABASE.NAME"); database == "" {
		database = viper.GetString(strings.ToLower("DATABASE.NAME"))
	}

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
