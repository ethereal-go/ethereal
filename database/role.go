package database

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"password"`
}

