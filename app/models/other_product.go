package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
)

type OtherProduct struct {
	gorm.Model

	Code string

	ColorVariats []ColorVariat
}

type ColorVariat struct {
	gorm.Model

	OtherProductID uint
	OtherProduct   OtherProduct

	ColorID uint
	Color   Color

	Images []OtherProductImage
}

type OtherProductImage struct {
	gorm.Model

	ColorVariatID uint
	ColorVariat   ColorVariat

	Kind  string
	Image media_library.FileSystem
}
