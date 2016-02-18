package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
)

type Car struct {
	gorm.Model
	Name           string `gorm:"column:name" sql:"type:varchar(30);unique_index" json:"name"`
	CarNumber      string
	OrganizationID uint
	Organization   Organization
	Drivers        []User `gorm:"many2many:driver_user;"`
	IsActive       bool   `gorm:"column:is_active"json:"active"`
	Comment        string
	Picture        media_library.FileSystem
}

// func (car Car) DisplayName() string {
// 	return car.Name
// }

func (car Car) Stringify() string {
	return fmt.Sprintf("%s (%s)", car.Name, car.CarNumber)
}
