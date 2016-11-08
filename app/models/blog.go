package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/accessibility/hyperlink"
)

type Article struct {
	gorm.Model
	Author   User
	AuthorID uint
	Title    string
	Content  string `gorm:"type:text"`
	FromURL  hyperlink.HyperLink
}
