package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
)

type Collection struct {
	gorm.Model
	Name string
	l10n.LocaleCreatable
	publish.Status
}
