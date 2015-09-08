package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
)

type Store struct {
	gorm.Model
	Name string
	location.Location
}
