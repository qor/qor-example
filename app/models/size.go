package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
	"github.com/qor/qor/sorting"
)

type Size struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`
}
