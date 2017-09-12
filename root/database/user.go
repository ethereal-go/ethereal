package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

type User struct {
	ID        uint   `json:"id";gorm:"primary_key"`
	Email     string `json:"email";gorm:"type:unique_index"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Role      Role   `json:"role"`
	RoleID    int    `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
