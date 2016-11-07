package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
)

type SEOSetting struct {
	gorm.Model
	l10n.Locale
	SiteName string
}
