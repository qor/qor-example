package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"not null;unique_index"` // TODO/problem: can't reuse emails

	Hashed []byte

	// only used for input validation and renewal
	Current  string `gorm:"-"`
	Password string `gorm:"-"`
	Repeat   string `gorm:"-"`

	Role   string
	Name   string
	Gender string
}

func (user User) DisplayName() string {
	return user.Name
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "zh-CN"}
}
