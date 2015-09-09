package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/sorting"
)

type Category struct {
	gorm.Model
	Name string
	sorting.Sorting
}
