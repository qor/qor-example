package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/media_library"
	"github.com/qor/qor/publish"
)

type Book struct {
	gorm.Model
	publish.Status
	l10n.Locale

	Title       string
	Synopsis    string
	ReleaseDate time.Time
	Authors     []*Author `gorm:"many2many:book_authors"`
	Price       float64
	CoverImage  CoverImage
	Types       []BookType
}

type CoverImage struct {
	media_library.FileSystem
}

func (CoverImage) GetSizes() map[string]media_library.Size {
	return map[string]media_library.Size{
		"list":       {Width: 100, Height: 100},
		"list@2x":    {Width: 200, Height: 200},
		"display":    {Width: 250, Height: 250},
		"display@2x": {Width: 500, Height: 500},
	}
}

type BookType struct {
	gorm.Model
	BookID uint
	Name   string
}
