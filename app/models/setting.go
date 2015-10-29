package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/qor/l10n"
	"github.com/qor/seo"
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
	location.Location `location:"name:Company Address"`
	l10n.Locale
}

type Seo struct {
	gorm.Model
	SiteName    string
	SiteHost    string
	HomePage    seo.Setting `seo:"category"`
	ProductPage seo.Setting `seo:"name,price"`
}
