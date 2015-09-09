package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/qor/sorting"
)

type Store struct {
	gorm.Model
	Name  string
	Phone string
	Email string
	location.Location
	sorting.Sorting
}
