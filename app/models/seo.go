package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/seo"
)

type SEOSetting struct {
	gorm.Model
	l10n.Locale
	SiteName    string
	DefaultPage seo.Setting
	HomePage    seo.Setting
	ProductPage seo.Setting `seo:"Name,Code"`
}
