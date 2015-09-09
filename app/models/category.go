package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
	"github.com/qor/qor/sorting"
)

type Category struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.Sorting
	Name string
}
