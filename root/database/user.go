package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
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

