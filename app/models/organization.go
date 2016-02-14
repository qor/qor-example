package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/media_library"
)

type Organization struct {
	gorm.Model
	Name      string `gorm:"column:name" sql:"type:varchar(30);unique_index" json:"username"`
	IsActive  bool   `gorm:"column:is_active"json:"active"`
	Director  string
	Email     string `sql:"type:varchar(75)" json:"email"`
	Phone     []Phone
	Logo      media_library.FileSystem
	Addresses []Address
	Comment   string
	location.Location
}

func (organization Organization) DisplayName() string {
	return organization.Name
}
