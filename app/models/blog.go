package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model

	Title string

	Cats []Cat `gorm:"many2many:blog_cats;ForeignKey:id;AssociationForeignKey:id"`
}

type Cat struct {
	gorm.Model

	Name string

	Blogs []Blog `gorm:"many2many:blog_cats;ForeignKey:id;AssociationForeignKey:id"`
}
