package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/qor/l10n"
)

type FeeSetting struct {
	ShippingFee     uint
	GiftWrappingFee uint
	CODFee          uint `gorm:"column:cod_fee"`
	TaxRate         int
}

type Setting struct {
	gorm.Model
	FeeSetting
	CompanyAddressID uint
	CompanyAddress   Location
	l10n.Locale
}

type Location struct {
	gorm.Model
	location.Location
}
