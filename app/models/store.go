package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/sorting"
)

type Store struct {
	gorm.Model
	Name  string
	Phone string
	Email string
	User  []User `gorm:"many2many:store_user;"`
	location.Location
	sorting.Sorting
}
