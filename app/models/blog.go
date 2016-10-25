package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Author   User
	AuthorID uint
	Title    string
	Content  string `gorm:"type:text"`
}
