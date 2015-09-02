package models

import "github.com/jinzhu/gorm"

type Color struct {
	gorm.Model
	Name string
	Code string
}
