package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string
	Password  string
	Name      string
	FirstName string
	LastNname string
	Gender    string
	Role      string
	Addresses []Address
}

func (user User) DisplayName() string {
	return user.Name
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "zh-CN", "ru-RU"}
}
