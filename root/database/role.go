package database

import (
	"time"
)


type Role struct {
	ID          uint   `json:"id";gorm:"primary_key"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}