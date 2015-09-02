package models

import "github.com/jinzhu/gorm"

type Store struct {
	gorm.Model
	Name      string
	AddressID uint
	Address   Address
}
