package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/media_library"
	"github.com/qor/qor/publish"
	"github.com/qor/qor/sorting"
	"github.com/qor/qor/validations"
	"github.com/qor/slug"
)

type Product struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.SortingDESC

	Name            string
	NameWithSlug    slug.Slug        `l10n:"sync"`
	Code            string           `l10n:"sync"`
	CategoryID      uint             `l10n:"sync"`
	Category        Category         `l10n:"sync"`
	MadeCountry     string           `l10n:"sync"`
	Price           float32          `l10n:"sync"`
	Description     string           `sql:"size:2000"`
	ColorVariations []ColorVariation `l10n:"sync"`
}

func (product Product) Validate(db *gorm.DB) {
	if strings.Trim(product.Name, " ") == "" {
		db.AddError(validations.NewError(product, "Name", "Name can not be empty"))
	}

	if strings.Trim(product.Code, " ") == "" {
		db.AddError(validations.NewError(product, "Code", "Code can not be empty"))
	}
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
