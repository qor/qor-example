package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/sorting"
)

type Color struct {
	gorm.Model
	sorting.Sorting
	Name string
	Code string
}
