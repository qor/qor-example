package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
)

type Collection struct {
	gorm.Model
	Name string
	l10n.LocaleCreatable
}
