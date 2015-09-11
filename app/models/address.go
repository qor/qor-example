package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	UserID      uint
	ContactName string
	Phone       string
	City        string
	Address1    string
	Address2    string
}
