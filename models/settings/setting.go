package settings

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/location"
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
