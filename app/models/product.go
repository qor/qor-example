package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/media_library"
	"github.com/qor/slug"
)

type Product struct {
	gorm.Model
	Name            string
	NameWithSlug    slug.Slug
	Code            string
	CategoryID      uint
	Category        Category
	MadeCountry     string
	Price           float32
	Description     string `sql:"size:2000"`
	Images          []ProductImage
	ColorVariations []ColorVariation
}

type ProductImage struct {
	gorm.Model
	Image media_library.FileSystem
}

type ColorVariation struct {
	gorm.Model
	ProductID      uint
	ColorID        uint
	Color          Color
	SizeVariations []SizeVariation
}

type SizeVariation struct {
	gorm.Model
	ColorVariationID  uint
	SizeID            uint
	Size              Size
	AvailableQuantity uint
}
