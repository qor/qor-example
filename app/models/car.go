package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
	"github.com/qor/validations"
)

type Car struct {
	gorm.Model
	Name           string `gorm:"column:name" sql:"type:varchar(30);unique_index" json:"name"`
	CarNumber      string
	OrganizationID uint
	Organization   Organization
	Drivers        []User `gorm:"many2many:driver_user;"`
	IsActive       bool   `sql:"default:false" gorm:"column:is_active"json:"active"`
	Comment        string
	Picture        media_library.FileSystem
}

// func (car Car) DisplayName() string {
// 	return car.Name
// }

func (car Car) Stringify() string {
	return fmt.Sprintf("%s (%s)", car.Name, car.CarNumber)
}

func (car Car) Validate(db *gorm.DB) {
	if strings.TrimSpace(car.Name) == "" {
		db.AddError(validations.NewError(car, "Name", "Name can not be empty"))
	}

	if strings.TrimSpace(car.CarNumber) == "" {
		db.AddError(validations.NewError(car, "CarNumber", "Car Number can not be empty"))
	}
}
