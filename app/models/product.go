package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/media_library"
	"github.com/qor/qor/sorting"
	"github.com/qor/slug"
)

type Product struct {
	gorm.Model
	sorting.SortingDESC
	Name            string
	NameWithSlug    slug.Slug
	Code            string
	CategoryID      uint
	Category        Category
	MadeCountry     string
	Price           float32
	Description     string `sql:"size:2000"`
	ColorVariations []ColorVariation
}

type ColorVariation struct {
	gorm.Model
	ProductID      uint
	ColorID        uint
	Color          Color
	Images         []ColorVariationImage
	SizeVariations []SizeVariation
}

type ColorVariationImage struct {
	gorm.Model
	ColorVariationID uint
	Image            ColorVariationImageStorage `sql:"type:varchar(4096)"`
}

type ColorVariationImageStorage struct{ media_library.FileSystem }

func (ColorVariationImageStorage) GetSizes() map[string]media_library.Size {
	return map[string]media_library.Size{
		"small":  {Width: 320, Height: 320},
		"middle": {Width: 640, Height: 640},
		"big":    {Width: 1280, Height: 1280},
	}
}

type SizeVariation struct {
	gorm.Model
	ColorVariationID  uint
	SizeID            uint
	Size              Size
	AvailableQuantity uint
}
