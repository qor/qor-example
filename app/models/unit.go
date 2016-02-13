package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/qor/publish"
	"github.com/qor/sorting"
	"github.com/qor/validations"
)

type Unit struct {
	gorm.Model
	l10n.Locale
	publish.Status
	sorting.Sorting
	Name string
}

func (unit Unit) Validate(db *gorm.DB) {
	if strings.TrimSpace(unit.Name) == "" {
		db.AddError(validations.NewError(unit, "Name", "Name can not be empty"))
	}
}
