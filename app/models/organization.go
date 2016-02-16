package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/media_library"
)

type Organization struct {
	gorm.Model
	Name           string `gorm:"column:name" sql:"type:varchar(30);unique_index" json:"name"`
	IsActive       bool   `gorm:"column:is_active"json:"active"`
	Director       string
	Email          string `sql:"type:varchar(75)" json:"email"`
	Phone          []Phone
	Logo           media_library.FileSystem
	Edrpou         uint
	Ipn            uint
	SvidPlanNaloga uint
	Score          uint64
	Bank           string
	Mfo            uint
	Addresses      []Address
	Comment        string
	location.Location
}

func (organization Organization) DisplayName() string {
	return organization.Name
}
