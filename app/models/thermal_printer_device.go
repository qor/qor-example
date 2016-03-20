package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
)

type ThermalPrinterDevice struct {
	gorm.Model
	Name         string `gorm:"column:name" sql:"type:varchar(30);unique_index" json:"name"`
	SerialNumber string
	Token        string
	CarID        uint
	Car          Car
	DeviceIp     string
	CurrentIp    string
	Status       string
	IsActive     bool `sql:"default:false" gorm:"column:is_active"json:"active"`
	Comment      string
}

func (tp ThermalPrinterDevice) DisplayName() string {
	return tp.Name
}

func (tp ThermalPrinterDevice) Stringify() string {
	return fmt.Sprintf("%s (%s)", tp.Name, tp.SerialNumber)
}

func (tp ThermalPrinterDevice) Validate(db *gorm.DB) {
	if strings.TrimSpace(tp.Name) == "" {
		db.AddError(validations.NewError(tp, "Name", "Name can not be empty"))
	}
}
