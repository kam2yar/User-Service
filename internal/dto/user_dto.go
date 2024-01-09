package dto

import (
	"gorm.io/gorm"
	"time"
)

type UserDto struct {
	id        uint
	name      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
	deletedAt gorm.DeletedAt
}

func (u *UserDto) SetId(i uint) {
	u.id = i
}

func (u *UserDto) GetId() uint {
	return u.id
}

func (u *UserDto) SetName(i string) {
	u.name = i
}

func (u *UserDto) GetName() string {
	return u.name
}

func (u *UserDto) SetEmail(i string) {
	u.email = i
}

func (u *UserDto) GetEmail() string {
	return u.email
}

func (u *UserDto) SetPassword(i string) {
	u.password = i
}

func (u *UserDto) GetPassword() string {
	return u.password
}

func (u *UserDto) SetCreatedAt(i time.Time) {
	u.createdAt = i
}

func (u *UserDto) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *UserDto) SetUpdatedAt(i time.Time) {
	u.updatedAt = i
}

func (u *UserDto) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *UserDto) SetDeletedAt(i gorm.DeletedAt) {
	u.deletedAt = i
}

func (u *UserDto) GetDeletedAt() gorm.DeletedAt {
	return u.deletedAt
}
