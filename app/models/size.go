package models

import "github.com/jinzhu/gorm"

type Size struct {
	gorm.Model
	Code string
	Name string
}
