package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/media_library"
)

type Product struct {
	gorm.Model
	Name            string
	Code            string
	Price           float32
	MadeCountry     string
	Description     string `sql:"size:2000"`
	Images          []ProductImage
	ColorVariations []ColorVariation
	Category        Category
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
