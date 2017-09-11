package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/ethereal-go/ethereal/root/database"
)

type UserService struct {
	DB *gorm.DB
}

// User returns a user for a given id.
func (s *UserService) Users(id int) (users []*database.User, err error) {
	s.DB.Find(&users)
	return
}

