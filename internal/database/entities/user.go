package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255;index:idx_email,unique"`
	Password string `gorm:"size:255"`
}
